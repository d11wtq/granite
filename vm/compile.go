package vm

import (
	"../parser/ast"
	"fmt"
)

// Compile the given AST into bytecode.
func Compile(prog ast.ASTNode) ([]byte, error) {
	compiler := NewCompiler()
	prog.Accept(compiler)
	return compiler.GetCode()
}

// Bytecode compiler for Bijou code
type Compiler struct {
	// Bytecode assembler
	Assembler *Assembler
	// Location of local variables in registers
	LocalMap map[string]uint32
	// Register where result should be kept
	RegIdx uint32
	// Pool of available registers
	RegPool *RegisterPool
	// Compile error
	Error error
}

// Create a new bytecode compiler.
func NewCompiler() *Compiler {
	cmp := &Compiler{
		Assembler: NewAssembler(),
		LocalMap:  make(map[string]uint32),
		RegPool:   NewRegisterPool(256),
	}

	cmp.RegIdx = cmp.RegPool.Reserve()

	return cmp
}

// Produce bytecode for the processed AST
func (c *Compiler) GetCode() ([]byte, error) {
	if c.Error != nil {
		return nil, c.Error
	}

	return c.Assembler.GetCode(), nil
}

func (c *Compiler) loadConstant(v Value) {
	c.Assembler.AddInstruction(
		encodeAxBx(OP_LOADK, c.RegIdx, c.Assembler.AddConstant(v)),
	)
}

func (c *Compiler) returnError(err error) {
	c.Error = err
	c.Assembler.AddInstruction(encodeAx(OP_RETURN, 0)) // safeguard-only
}

func (c *Compiler) VisitNil(node *ast.NilNode) {
	c.loadConstant(Nil)
}

func (c *Compiler) VisitInteger(node *ast.IntegerNode) {
	c.loadConstant(Integer(node.Value))
}

func (c *Compiler) VisitBoolean(node *ast.BooleanNode) {
	c.loadConstant(Boolean(node.Value))
}

func (c *Compiler) VisitFloat(node *ast.FloatNode) {
}

func (c *Compiler) VisitString(node *ast.StringNode) {
	c.loadConstant(String(node.String))
}

func (c *Compiler) VisitSymbol(node *ast.SymbolNode) {
}

func (c *Compiler) VisitIdentifier(node *ast.IdentifierNode) {
	if regIdx, ok := c.LocalMap[node.Name]; ok {
		c.RegIdx = regIdx
		return
	}

	c.returnError(NameError(node.Name))
}

func (c *Compiler) VisitExpressionList(node *ast.ExpressionList) {
	regIdx := c.RegIdx
	tailIdx := len(node.Elements) - 1

	for i, e := range node.Elements {
		e.Accept(c)
		if c.Error != nil {
			return
		}

		if c.RegIdx != regIdx && i == tailIdx {
			c.Assembler.AddInstruction(encodeAxBx(OP_MOVE, regIdx, c.RegIdx))
		}
		c.RegIdx = regIdx
	}
}

func (c *Compiler) visitAssignment(node *ast.BinaryExpressionNode) {
	id := node.Left.(*ast.IdentifierNode)
	regX, ok := c.LocalMap[id.Name]
	if ok == false {
		c.LocalMap[id.Name] = c.RegPool.Reserve()
		node.Right.Accept(c)
		c.Assembler.AddInstruction(encodeAxBx(OP_MOVE, c.LocalMap[id.Name], c.RegIdx))
	} else {
		regA := c.RegIdx
		regB := c.RegPool.Reserve()
		c.RegIdx = regB
		node.Right.Accept(c)
		c.Assembler.AddInstruction(encodeAxBxCx(OP_EQ, regA, regX, c.RegIdx))
		c.Assembler.AddInstruction(encodeAxBx(OP_JMPIF, regA, 1))
		c.Assembler.AddInstruction(encodeAxBxCx(OP_ERR, E_BADMATCH, regX, c.RegIdx))
		c.Assembler.AddInstruction(encodeAxBx(OP_MOVE, regA, c.RegIdx))
		c.RegIdx = regA
		c.RegPool.Release(regB)
	}
}

func (c *Compiler) VisitBinaryExpression(node *ast.BinaryExpressionNode) {
	if node.Op == ast.OP_ASS {
		c.visitAssignment(node)
		return
	}

	regA := c.RegIdx
	node.Left.Accept(c)
	regB := c.RegIdx

	regC := c.RegPool.Reserve()
	c.RegIdx = regC
	node.Right.Accept(c)
	c.RegPool.Release(regC)

	switch node.Op {
	case ast.OP_ADD:
		c.Assembler.AddInstruction(encodeAxBxCx(OP_ADD, regA, regB, c.RegIdx))
	case ast.OP_MIN:
		c.Assembler.AddInstruction(encodeAxBxCx(OP_SUB, regA, regB, c.RegIdx))
	case ast.OP_MUL:
		c.Assembler.AddInstruction(encodeAxBxCx(OP_MUL, regA, regB, c.RegIdx))
	case ast.OP_DIV:
		c.Assembler.AddInstruction(encodeAxBxCx(OP_DIV, regA, regB, c.RegIdx))
	case ast.OP_EQL:
		c.Assembler.AddInstruction(encodeAxBxCx(OP_EQ, regA, regB, c.RegIdx))
	case ast.OP_LT:
		c.Assembler.AddInstruction(encodeAxBxCx(OP_LT, regA, regB, c.RegIdx))
	case ast.OP_GT:
		c.Assembler.AddInstruction(encodeAxBxCx(OP_LT, regA, c.RegIdx, regB))
	case ast.OP_LTE:
		c.Assembler.AddInstruction(encodeAxBxCx(OP_LTE, regA, regB, c.RegIdx))
	case ast.OP_GTE:
		c.Assembler.AddInstruction(encodeAxBxCx(OP_LTE, regA, c.RegIdx, regB))
	default:
		panic(fmt.Sprintf("Unhandled binary operator: 0x%x", node.Op))
	}

	c.RegIdx = regA
}

func (c *Compiler) VisitUnaryExpression(node *ast.UnaryExpressionNode) {
}

func (c *Compiler) VisitVector(node *ast.VectorNode) {
	regA := c.RegIdx
	c.loadConstant(EmptyVector)

	regB := c.RegPool.Reserve()
	for _, e := range node.Elements {
		c.RegIdx = regB
		e.Accept(c)
		c.Assembler.AddInstruction(
			encodeAxBxCx(OP_APPEND, regA, regA, c.RegIdx),
		)
	}

	c.RegPool.Release(regB)
	c.RegIdx = regA
}

func (c *Compiler) VisitMap(node *ast.MapNode) {
}

func (c *Compiler) VisitRecord(node *ast.RecordNode) {
}

func (c *Compiler) VisitPair(node *ast.PairNode) {
}

func (c *Compiler) VisitFunctionApplication(node *ast.FunctionApplicationNode) {
}

func (c *Compiler) VisitKeyAccess(node *ast.KeyAccessNode) {
}

func (c *Compiler) VisitIfThenElse(node *ast.IfThenElseNode) {
	var (
		regIdx         = c.RegIdx
		jmpIf, endElse int
	)

	node.Expression.Accept(c)

	jmpIf = c.Assembler.ReserveInstruction(OP_JMPIF)
	node.Else.Accept(c)
	endElse = c.Assembler.ReserveInstruction(OP_JMP)
	node.Then.Accept(c)

	c.Assembler.SetInstruction(
		jmpIf,
		encodeAxBx(OP_JMPIF, regIdx, uint32(endElse-jmpIf)),
	)

	c.Assembler.SetInstruction(
		endElse,
		encodeAx(OP_JMP, uint32(len(c.Assembler.Instructions)-1-endElse)),
	)
}

func (c *Compiler) VisitCaseExpression(node *ast.CaseExpressionNode) {
}

func (c *Compiler) VisitCatchExpression(node *ast.CatchExpressionNode) {
}

func (c *Compiler) VisitThrow(node *ast.ThrowNode) {
}
