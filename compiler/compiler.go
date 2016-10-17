package compiler

import (
	"../parser/ast"
	"../vm"
	"bytes"
	"encoding/binary"
)

// Bytecode compiler for Bijou code
type Compiler struct {
	Constants   []vm.Value
	ConstantMap map[vm.Value]int
}

func NewCompiler() *Compiler {
	return &Compiler{
		Constants:   make([]vm.Value, 0),
		ConstantMap: make(map[vm.Value]int),
	}
}

// Produce bytecode for the processed AST
func (c *Compiler) GetCode() []byte {
	b := new(bytes.Buffer)
	c.encodeConstants(b)
	c.encodeInstructions(b)
	return b.Bytes()
}

func (c *Compiler) encodeConstants(b *bytes.Buffer) {
	binary.Write(b, vm.ByteOrder, uint32(len(c.Constants)))
	for _, v := range c.Constants {
		binary.Write(b, vm.ByteOrder, v.Type())
		switch t := v.(type) {
		case vm.Integer:
			binary.Write(b, vm.ByteOrder, v)
		case vm.String:
			s := []byte(string(t))
			binary.Write(b, vm.ByteOrder, uint32(len(s)))
			b.Write(s)
		}
	}
}

func (c *Compiler) encodeInstructions(b *bytes.Buffer) {
	binary.Write(b, vm.ByteOrder, uint64(5)) // number of instructions

	var op, ax, bx, cx uint32

	op = vm.OP_LOADK
	ax = 0
	bx = 1
	binary.Write(
		b,
		vm.ByteOrder,
		// 00000000 00000000 00000000 00111111
		// 11111100 00000000 00000000 00000000 // op << 26

		// 00000000 00000000 00000001 11111111
		// 00000011 11111110 00000000 00000000 // ax << 17

		// 00000000 00000000 00000001 11111111
		// 00000000 00000001 11111111 00000000 // bx << 8

		// 00000000 00000000 00000000 11111111 // cx
		uint32((op<<26)|(ax<<17)|(bx<<8)|cx),
	)

	op = vm.OP_LOADK
	ax = 1
	bx = 0
	binary.Write(
		b,
		vm.ByteOrder,
		uint32((op<<26)|(ax<<17)|(bx<<8)|cx),
	)

	op = vm.OP_ADD
	ax = 0
	bx = 0
	cx = 1
	binary.Write(
		b,
		vm.ByteOrder,
		uint32((op<<26)|(ax<<17)|(bx<<8)|cx),
	)

	op = vm.OP_PRINT
	ax = 0
	binary.Write(
		b,
		vm.ByteOrder,
		uint32((op<<26)|ax),
	)

	op = vm.OP_RETURN
	binary.Write(b, vm.ByteOrder, uint32(op<<26))

	/*
		binary.Write(
			b,
			vm.ByteOrder,
			uint32((op<<26)|(ax<<13)|bx),
		)

		op = 5
		ax = 896764

		binary.Write(
			b,
			vm.ByteOrder,
			uint32((op<<26)|ax),
		)
	*/
}

func (c *Compiler) addConstant(v vm.Value) {
	if _, ok := c.ConstantMap[v]; ok == false {
		c.ConstantMap[v] = len(c.Constants)
		c.Constants = append(c.Constants, v)
	}
}

func (c *Compiler) VisitNil(node *ast.NilNode) {
}

func (c *Compiler) VisitInteger(node *ast.IntegerNode) {
	c.addConstant(vm.Integer(node.Value))
}

func (c *Compiler) VisitBoolean(node *ast.BooleanNode) {
}

func (c *Compiler) VisitFloat(node *ast.FloatNode) {
}

func (c *Compiler) VisitString(node *ast.StringNode) {
	c.addConstant(vm.String(node.String))
}

func (c *Compiler) VisitSymbol(node *ast.SymbolNode) {
}

func (c *Compiler) VisitIdentifier(node *ast.IdentifierNode) {
}

func (c *Compiler) VisitExpressionList(node *ast.ExpressionList) {
	for _, e := range node.Elements {
		e.Accept(c)
	}
}

func (c *Compiler) VisitBinaryExpression(node *ast.BinaryExpressionNode) {
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
