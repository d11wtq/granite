package compiler

import (
	"../parser/ast"
	"../vm"
	"bytes"
	"encoding/binary"
	"fmt"
)

// Convert boolean v to a uint8.
func Btoi(v bool) uint8 {
	if v {
		return 1
	}

	return 0
}

// Bytecode compiler for Bijou code
type Compiler struct {
	Constants    []vm.Value
	ConstantMap  map[vm.Value]uint32
	Instructions []uint32
	LocalMap     map[string]uint8
	RegIdx       uint32
}

// Create a new bytecode compiler.
func NewCompiler() *Compiler {
	return &Compiler{
		Constants:    make([]vm.Value, 0),
		ConstantMap:  make(map[vm.Value]uint32),
		Instructions: make([]uint32, 0),
		LocalMap:     make(map[string]uint8),
		RegIdx:       0,
	}
}

// Produce bytecode for the processed AST
func (c *Compiler) GetCode() []byte {
	b := new(bytes.Buffer)
	c.encodeConstants(b)
	c.encodeInstructions(b)
	return b.Bytes()
}

// Encode a three operand instruction.
func encodeAxBxCx(op, ax, bx, cx uint32) uint32 {
	// 00000000 00000000 00000000 00111111 // op
	// 11111100 00000000 00000000 00000000 // op << 26

	// 00000000 00000000 00000001 11111111 // ax
	// 00000011 11111110 00000000 00000000 // ax << 17

	// 00000000 00000000 00000001 11111111 // bx
	// 00000000 00000001 11111111 00000000 // bx << 8

	// 00000000 00000000 00000000 11111111 // cx
	return (op << 26) | (ax << 17) | (bx << 8) | cx
}

// Encode a two operand instruction.
func encodeAxBx(op, ax, bx uint32) uint32 {
	// 00000000 00000000 00000000 00111111 // op
	// 11111100 00000000 00000000 00000000 // op << 26

	// 00000000 00000000 00111111 11111111 // ax
	// 00000011 11111111 11100000 00000000 // ax << 13

	// 00000000 00000000 00011111 11111111 // bx
	return (op << 26) | (ax << 13) | bx
}

// Encode a one operand instruction.
func encodeAx(op, ax uint32) uint32 {
	// 00000000 00000000 00000000 00111111 // op
	// 11111100 00000000 00000000 00000000 // op << 26

	// 00000011 11111111 11111111 11111111 // ax
	return (op << 26) | ax
}

func (c *Compiler) encodeConstants(b *bytes.Buffer) {
	binary.Write(b, vm.ByteOrder, uint32(len(c.Constants)))
	for _, v := range c.Constants {
		binary.Write(b, vm.ByteOrder, v.Type())
		switch t := v.(type) {
		case vm.Integer:
			binary.Write(b, vm.ByteOrder, v)
		case vm.Boolean:
			binary.Write(b, vm.ByteOrder, Btoi(bool(t)))
		case vm.String:
			s := []byte(string(t))
			binary.Write(b, vm.ByteOrder, uint32(len(s)))
			b.Write(s)
		}
	}
}

func (c *Compiler) encodeInstructions(b *bytes.Buffer) {
	binary.Write(b, vm.ByteOrder, uint64(len(c.Instructions)+2))

	for _, inst := range c.Instructions {
		binary.Write(b, vm.ByteOrder, inst)
	}

	binary.Write(b, vm.ByteOrder, encodeAx(vm.OP_PRINT, 0))
	binary.Write(b, vm.ByteOrder, encodeAx(vm.OP_RETURN, 0))
}

func (c *Compiler) addConstant(v vm.Value) uint32 {
	if _, ok := c.ConstantMap[v]; ok == false {
		c.ConstantMap[v] = uint32(len(c.Constants))
		c.Constants = append(c.Constants, v)
	}

	return c.ConstantMap[v]
}

func (c *Compiler) addInstruction(inst uint32) {
	c.Instructions = append(c.Instructions, inst)
}

func (c *Compiler) VisitNil(node *ast.NilNode) {
}

func (c *Compiler) VisitInteger(node *ast.IntegerNode) {
	cIdx := c.addConstant(vm.Integer(node.Value))
	c.addInstruction(encodeAxBx(vm.OP_LOADK, c.RegIdx, cIdx))
	c.RegIdx++
}

func (c *Compiler) VisitBoolean(node *ast.BooleanNode) {
	cIdx := c.addConstant(vm.Boolean(node.Value))
	c.addInstruction(encodeAxBx(vm.OP_LOADK, c.RegIdx, cIdx))
	c.RegIdx++
}

func (c *Compiler) VisitFloat(node *ast.FloatNode) {
}

func (c *Compiler) VisitString(node *ast.StringNode) {
	cIdx := c.addConstant(vm.String(node.String))
	c.addInstruction(encodeAxBx(vm.OP_LOADK, c.RegIdx, cIdx))
	c.RegIdx++
}

func (c *Compiler) VisitSymbol(node *ast.SymbolNode) {
}

func (c *Compiler) VisitIdentifier(node *ast.IdentifierNode) {
}

func (c *Compiler) VisitExpressionList(node *ast.ExpressionList) {
	for _, e := range node.Elements {
		e.Accept(c)
		c.RegIdx = uint32(len(c.LocalMap))
	}
}

func (c *Compiler) VisitBinaryExpression(node *ast.BinaryExpressionNode) {
	regIdx := c.RegIdx
	node.Left.Accept(c)
	node.Right.Accept(c)
	switch node.Op {
	case ast.OP_ADD:
		c.addInstruction(encodeAxBxCx(vm.OP_ADD, regIdx, regIdx, regIdx+1))
	case ast.OP_MIN:
		c.addInstruction(encodeAxBxCx(vm.OP_SUB, regIdx, regIdx, regIdx+1))
	case ast.OP_MUL:
		c.addInstruction(encodeAxBxCx(vm.OP_MUL, regIdx, regIdx, regIdx+1))
	case ast.OP_DIV:
		c.addInstruction(encodeAxBxCx(vm.OP_DIV, regIdx, regIdx, regIdx+1))
	case ast.OP_EQL:
		c.addInstruction(encodeAxBxCx(vm.OP_EQ, regIdx, regIdx, regIdx+1))
	case ast.OP_LT:
		c.addInstruction(encodeAxBxCx(vm.OP_LT, regIdx, regIdx, regIdx+1))
	case ast.OP_GT:
		c.addInstruction(encodeAxBxCx(vm.OP_LT, regIdx, regIdx+1, regIdx))
	case ast.OP_LTE:
		c.addInstruction(encodeAxBxCx(vm.OP_LTE, regIdx, regIdx, regIdx+1))
	case ast.OP_GTE:
		c.addInstruction(encodeAxBxCx(vm.OP_LTE, regIdx, regIdx+1, regIdx))
	default:
		panic(fmt.Sprintf("Unhandled binary operator: 0x%x", node.Op))
	}
	c.RegIdx--
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
}

func (c *Compiler) VisitCaseExpression(node *ast.CaseExpressionNode) {
}

func (c *Compiler) VisitCatchExpression(node *ast.CatchExpressionNode) {
}

func (c *Compiler) VisitThrow(node *ast.ThrowNode) {
}
