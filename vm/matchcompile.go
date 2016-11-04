package vm

import (
	"../parser/ast"
)

// Bytecode compiler for Bijou patterns
type MatchCompiler struct {
	// The target compiler
	Compiler *Compiler
	// The location of the value being matched
	RegIdx uint32
}

// Create a new compiler for handling patterns
func NewMatchCompiler(c *Compiler) *MatchCompiler {
	return &MatchCompiler{c, c.RegIdx}
}

func (c *MatchCompiler) assignLocalVariable(node *ast.IdentifierNode) {
	c.Compiler.LocalMap[node.Name] = c.Compiler.RegPool.Reserve()
	c.Compiler.Assembler.Move(Reg(c.Compiler.LocalMap[node.Name]), Reg(c.RegIdx))
}

func (c *MatchCompiler) assertEq(node ast.ASTNode) {
	var (
		regA = c.Compiler.RegIdx
		regB = c.Compiler.RegPool.Reserve()
		regC = c.Compiler.RegPool.Reserve()
		isEq = c.Compiler.Assembler.GenLabel()
	)

	c.Compiler.RegIdx = regB
	node.Accept(c.Compiler)

	c.Compiler.Assembler.Eq(Reg(regC), Reg(c.Compiler.RegIdx), Reg(c.RegIdx))
	c.Compiler.Assembler.JmpIf(Reg(regC), Jmp(isEq))
	c.errorBadMatch()
	c.Compiler.Assembler.SetLabel(isEq)

	c.Compiler.RegPool.Release(regB)
	c.Compiler.RegPool.Release(regC)
	c.Compiler.RegIdx = regA
}

func (c *MatchCompiler) errorBadMatch() {
	c.Compiler.Assembler.Err(Reg(E_BADMATCH), Reg(c.RegIdx))
}

func (c *MatchCompiler) VisitNil(node *ast.NilNode) {
	c.assertEq(node)
}

func (c *MatchCompiler) VisitInteger(node *ast.IntegerNode) {
	c.assertEq(node)
}

func (c *MatchCompiler) VisitBoolean(node *ast.BooleanNode) {
	c.assertEq(node)
}

func (c *MatchCompiler) VisitFloat(node *ast.FloatNode) {
	c.assertEq(node)
}

func (c *MatchCompiler) VisitString(node *ast.StringNode) {
	c.assertEq(node)
}

func (c *MatchCompiler) VisitSymbol(node *ast.SymbolNode) {
	c.assertEq(node)
}

func (c *MatchCompiler) VisitIdentifier(node *ast.IdentifierNode) {
	if _, isExisting := c.Compiler.LocalMap[node.Name]; isExisting {
		c.assertEq(node)
		return
	}

	c.assignLocalVariable(node)
}

func (c *MatchCompiler) VisitExpressionList(node *ast.ExpressionList) {
	c.assertEq(node)
}

func (c *MatchCompiler) VisitBinaryExpression(node *ast.BinaryExpressionNode) {
	if node.Op != ast.OP_ASS {
		// error
	}
}

func (c *MatchCompiler) VisitUnaryExpression(node *ast.UnaryExpressionNode) {
}

func (c *MatchCompiler) VisitVector(node *ast.VectorNode) {
	var (
		regX   = c.Compiler.RegIdx
		regA   = c.Compiler.RegPool.Reserve()
		regB   = c.Compiler.RegPool.Reserve()
		typeOk = c.Compiler.Assembler.GenLabel()
		lenOk  = c.Compiler.Assembler.GenLabel()
	)

	// Make sure RHS is a Vector too
	c.Compiler.Assembler.IsA(Reg(regA), Reg(c.RegIdx), Reg(V_VEC))
	c.Compiler.Assembler.JmpIf(Reg(regA), Jmp(typeOk))
	c.errorBadMatch()
	c.Compiler.Assembler.SetLabel(typeOk)

	// Make sure RHS is the same length
	c.Compiler.Assembler.Len(Reg(regA), Reg(c.RegIdx))
	c.Compiler.RegIdx = regB
	c.Compiler.loadConstant(Integer(len(node.Elements)))
	c.Compiler.Assembler.Eq(Reg(regA), Reg(regA), Reg(regB))
	c.Compiler.Assembler.JmpIf(Reg(regA), Jmp(lenOk))
	c.errorBadMatch()
	c.Compiler.Assembler.SetLabel(lenOk)

	c.Compiler.RegPool.Release(regB)

	// Recurse across all the elements within the vector
	for i, e := range node.Elements {
		c.Compiler.RegIdx = regA
		c.Compiler.loadConstant(Integer(i))
		c.Compiler.Assembler.Get(
			Reg(regA),
			Reg(c.RegIdx),
			Reg(c.Compiler.RegIdx),
		)
		e.Accept(NewMatchCompiler(c.Compiler))
	}

	c.Compiler.RegPool.Release(regA)
	c.Compiler.RegIdx = regX
}

func (c *MatchCompiler) VisitMap(node *ast.MapNode) {
}

func (c *MatchCompiler) VisitRecord(node *ast.RecordNode) {
}

func (c *MatchCompiler) VisitPair(node *ast.PairNode) {
}

func (c *MatchCompiler) VisitFunctionApplication(node *ast.FunctionApplicationNode) {
}

func (c *MatchCompiler) VisitKeyAccess(node *ast.KeyAccessNode) {
}

func (c *MatchCompiler) VisitIfThenElse(node *ast.IfThenElseNode) {
}

func (c *MatchCompiler) VisitCaseExpression(node *ast.CaseExpressionNode) {
}

func (c *MatchCompiler) VisitCatchExpression(node *ast.CatchExpressionNode) {
}

func (c *MatchCompiler) VisitThrow(node *ast.ThrowNode) {
}
