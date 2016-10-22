package vm

import (
	"../parser/ast"
	"bytes"
	"encoding/binary"
	"fmt"
)

// Compile the given AST into bytecode.
func Compile(prog ast.ASTNode) ([]byte, error) {
	compiler := NewCompiler()
	prog.Accept(compiler)
	return compiler.GetCode()
}

// Convert boolean v to a uint8.
func btoi(v bool) uint8 {
	if v {
		return 1
	}

	return 0
}

// Bytecode compiler for Bijou code
type Compiler struct {
	// Program constant pool
	Constants []Value
	// Locations of constants in the pool (duplicates removed)
	ConstantMap map[Value]uint32
	// List of executable instructions
	Instructions []uint32
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
		Constants:    make([]Value, 0),
		ConstantMap:  make(map[Value]uint32),
		Instructions: make([]uint32, 0),
		LocalMap:     make(map[string]uint32),
		RegPool:      NewRegisterPool(256),
	}

	cmp.RegIdx = cmp.RegPool.Reserve()

	return cmp
}

// Produce bytecode for the processed AST
func (c *Compiler) GetCode() ([]byte, error) {
	if c.Error != nil {
		return nil, c.Error
	}

	b := new(bytes.Buffer)
	c.encodeConstants(b)
	c.encodeInstructions(b)

	return b.Bytes(), nil
}

func (c *Compiler) encodeConstants(b *bytes.Buffer) {
	binary.Write(b, ByteOrder, uint32(len(c.Constants)))
	for _, v := range c.Constants {
		binary.Write(b, ByteOrder, v.Type())
		switch t := v.(type) {
		case *NilType:
			// noop
		case Integer:
			binary.Write(b, ByteOrder, v)
		case Boolean:
			binary.Write(b, ByteOrder, btoi(bool(t)))
		case String:
			s := []byte(string(t))
			binary.Write(b, ByteOrder, uint32(len(s)))
			b.Write(s)
		}
	}
}

func (c *Compiler) encodeInstructions(b *bytes.Buffer) {
	binary.Write(b, ByteOrder, uint64(len(c.Instructions)+2))

	for _, inst := range c.Instructions {
		binary.Write(b, ByteOrder, inst)
	}

	binary.Write(b, ByteOrder, encodeAx(OP_PRINT, 0))
	binary.Write(b, ByteOrder, encodeAx(OP_RETURN, 0))
}

func (c *Compiler) addConstant(v Value) {
	if _, ok := c.ConstantMap[v]; ok == false {
		c.ConstantMap[v] = uint32(len(c.Constants))
		c.Constants = append(c.Constants, v)
	}

	c.addInstruction(encodeAxBx(OP_LOADK, c.RegIdx, c.ConstantMap[v]))
}

func (c *Compiler) addInstruction(inst uint32) {
	c.Instructions = append(c.Instructions, inst)
}

func (c *Compiler) reserveInstruction(int) int {
	idx := len(c.Instructions)
	c.addInstruction(0)
	return idx
}

func (c *Compiler) setInstruction(idx int, inst uint32) {
	c.Instructions[idx] = inst
}

func (c *Compiler) returnError(err error) {
	c.Error = err
	c.addInstruction(encodeAx(OP_RETURN, 0)) // safeguard-only
}

func (c *Compiler) VisitNil(node *ast.NilNode) {
	c.addConstant(Nil)
}

func (c *Compiler) VisitInteger(node *ast.IntegerNode) {
	c.addConstant(Integer(node.Value))
}

func (c *Compiler) VisitBoolean(node *ast.BooleanNode) {
	c.addConstant(Boolean(node.Value))
}

func (c *Compiler) VisitFloat(node *ast.FloatNode) {
}

func (c *Compiler) VisitString(node *ast.StringNode) {
	c.addConstant(String(node.String))
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
	for _, e := range node.Elements {
		e.Accept(c)
		if c.Error != nil {
			return
		}

		if c.RegIdx != regIdx {
			c.addInstruction(encodeAxBx(OP_MOVE, regIdx, c.RegIdx))
			c.RegIdx = regIdx
		}
	}
}

func (c *Compiler) visitAssignment(node *ast.BinaryExpressionNode) {
	id := node.Left.(*ast.IdentifierNode)
	regX, ok := c.LocalMap[id.Name]
	if ok == false {
		c.LocalMap[id.Name] = c.RegPool.Reserve()
		node.Right.Accept(c)
		c.addInstruction(encodeAxBx(OP_MOVE, c.LocalMap[id.Name], c.RegIdx))
	} else {
		regA := c.RegIdx
		regB := c.RegPool.Reserve()
		c.RegIdx = regB
		node.Right.Accept(c)
		c.addInstruction(encodeAxBxCx(OP_EQ, regA, regX, c.RegIdx))
		c.addInstruction(encodeAxBx(OP_JMPIF, regA, 1))
		c.addInstruction(encodeAxBxCx(OP_ERR, E_BADMATCH, regX, c.RegIdx))
		c.addInstruction(encodeAxBx(OP_MOVE, regA, c.RegIdx))
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
		c.addInstruction(encodeAxBxCx(OP_ADD, regA, regB, c.RegIdx))
	case ast.OP_MIN:
		c.addInstruction(encodeAxBxCx(OP_SUB, regA, regB, c.RegIdx))
	case ast.OP_MUL:
		c.addInstruction(encodeAxBxCx(OP_MUL, regA, regB, c.RegIdx))
	case ast.OP_DIV:
		c.addInstruction(encodeAxBxCx(OP_DIV, regA, regB, c.RegIdx))
	case ast.OP_EQL:
		c.addInstruction(encodeAxBxCx(OP_EQ, regA, regB, c.RegIdx))
	case ast.OP_LT:
		c.addInstruction(encodeAxBxCx(OP_LT, regA, regB, c.RegIdx))
	case ast.OP_GT:
		c.addInstruction(encodeAxBxCx(OP_LT, regA, c.RegIdx, regB))
	case ast.OP_LTE:
		c.addInstruction(encodeAxBxCx(OP_LTE, regA, regB, c.RegIdx))
	case ast.OP_GTE:
		c.addInstruction(encodeAxBxCx(OP_LTE, regA, c.RegIdx, regB))
	default:
		panic(fmt.Sprintf("Unhandled binary operator: 0x%x", node.Op))
	}

	c.RegIdx = regA
}

func (c *Compiler) VisitUnaryExpression(node *ast.UnaryExpressionNode) {
}

func (c *Compiler) VisitVector(node *ast.VectorNode) {
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

	jmpIf = c.reserveInstruction(OP_JMPIF)
	node.Else.Accept(c)
	endElse = c.reserveInstruction(OP_JMP)
	node.Then.Accept(c)

	c.setInstruction(
		jmpIf,
		encodeAxBx(OP_JMPIF, regIdx, uint32(endElse-jmpIf)),
	)

	c.setInstruction(
		endElse,
		encodeAx(OP_JMP, uint32(len(c.Instructions)-1-endElse)),
	)
}

func (c *Compiler) VisitCaseExpression(node *ast.CaseExpressionNode) {
}

func (c *Compiler) VisitCatchExpression(node *ast.CatchExpressionNode) {
}

func (c *Compiler) VisitThrow(node *ast.ThrowNode) {
}
