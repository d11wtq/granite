package vm

import (
	"../parser/ast"
	"fmt"
)

// FIXME: Ditch this and replace with AST -> CFG -> Code conversion.

// Compile the given AST into bytecode.
func Compile(prog ast.ASTNode) ([]byte, error) {
	compiler := NewCompiler()
	prog.Accept(compiler)
	return compiler.GetCode()
}

// Bytecode compiler for Bijou code
type Compiler struct {
	// Bytecode assembler
	Assembler *ASM
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
	regPool := NewRegisterPool(256)

	return &Compiler{
		Assembler: NewASM(),
		LocalMap:  make(map[string]uint32),
		RegPool:   regPool,
		RegIdx:    regPool.Reserve(),
	}
}

// Produce bytecode for the processed AST
func (c *Compiler) GetCode() ([]byte, error) {
	if c.Error != nil {
		return nil, c.Error
	}

	c.Assembler.Print(Reg(0))
	c.Assembler.Return()

	return c.Assembler.GetCode(), nil
}

func (c *Compiler) loadConstant(v Value) {
	c.Assembler.LoadK(Reg(c.RegIdx), &Constant{v})
}

func (c *Compiler) returnError(err error) {
	c.Error = err
	c.Assembler.Return()
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
			c.Assembler.Move(Reg(regIdx), Reg(c.RegIdx))
		}
		c.RegIdx = regIdx
	}
}

func (c *Compiler) visitAssignment(node *ast.BinaryExpressionNode) {
	node.Right.Accept(c)
	regIdx := c.RegIdx
	node.Left.Accept(NewMatchCompiler(c))
	c.RegIdx = regIdx
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
		c.Assembler.Add(Reg(regA), Reg(regB), Reg(c.RegIdx))
	case ast.OP_MIN:
		c.Assembler.Sub(Reg(regA), Reg(regB), Reg(c.RegIdx))
	case ast.OP_MUL:
		c.Assembler.Mul(Reg(regA), Reg(regB), Reg(c.RegIdx))
	case ast.OP_DIV:
		c.Assembler.Div(Reg(regA), Reg(regB), Reg(c.RegIdx))
	case ast.OP_EQL:
		c.Assembler.Eq(Reg(regA), Reg(regB), Reg(c.RegIdx))
	case ast.OP_LT:
		c.Assembler.Lt(Reg(regA), Reg(regB), Reg(c.RegIdx))
	case ast.OP_GT:
		c.Assembler.Lt(Reg(regA), Reg(c.RegIdx), Reg(regB))
	case ast.OP_LTE:
		c.Assembler.Lte(Reg(regA), Reg(regB), Reg(c.RegIdx))
	case ast.OP_GTE:
		c.Assembler.Lte(Reg(regA), Reg(c.RegIdx), Reg(regB))
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
		c.Assembler.Append(Reg(regA), Reg(regA), Reg(c.RegIdx))

		if c.Error != nil {
			break
		}
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
		regIdx = c.RegIdx
		thenL  = c.Assembler.GenLabel()
		doneL  = c.Assembler.GenLabel()
	)

	node.Expression.Accept(c)

	c.Assembler.JmpIf(Reg(regIdx), Jmp(thenL))
	node.Else.Accept(c)
	c.Assembler.Jmp(Jmp(doneL))
	c.Assembler.SetLabel(thenL)
	node.Then.Accept(c)
	c.Assembler.SetLabel(doneL)
}

func (c *Compiler) VisitCaseExpression(node *ast.CaseExpressionNode) {
}

func (c *Compiler) VisitCatchExpression(node *ast.CatchExpressionNode) {
}

func (c *Compiler) VisitThrow(node *ast.ThrowNode) {
}
