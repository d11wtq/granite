package vm

import (
	"../parser/ast"
)

// Generate byte code for a pattern match routine.
// Targets a currently in-progress Compiler.
// v is the register location of the value to match against.
func compilePattern(c *Compiler, pattern ast.ASTNode, v Operand) {
	if c.Error != nil {
		return
	}

	pattern.Accept(newpatternCompiler(c, v))
	c.Result = v // propagate back down
}

// Compiler to produce byte code for a pattern match.
type patternCompiler struct {
	compiler *Compiler
	value    Operand
}

// Create a new patternCompiler targeting the given Compiler.
func newpatternCompiler(c *Compiler, v Operand) *patternCompiler {
	return &patternCompiler{c, v}
}

func (p *patternCompiler) VisitNil(node *ast.NilNode) {
	p.assertMatch(p.compiler.Visit(node).Result)
}

func (p *patternCompiler) VisitInteger(node *ast.IntegerNode) {
	p.assertMatch(p.compiler.Visit(node).Result)
}

func (p *patternCompiler) VisitBoolean(node *ast.BooleanNode) {
	p.assertMatch(p.compiler.Visit(node).Result)
}

func (p *patternCompiler) VisitFloat(node *ast.FloatNode) {
	p.assertMatch(p.compiler.Visit(node).Result)
}

func (p *patternCompiler) VisitString(node *ast.StringNode) {
	p.assertMatch(p.compiler.Visit(node).Result)
}

func (p *patternCompiler) VisitSymbol(node *ast.SymbolNode) {
	p.assertMatch(p.compiler.Visit(node).Result)
}

func (p *patternCompiler) VisitIdentifier(node *ast.IdentifierNode) {
	if v, exists := p.compiler.Symbols.Get(node.Name); exists {
		p.assertMatch(v)
		return
	}

	p.compiler.Symbols.Set(node.Name, p.value)
}

func (p *patternCompiler) VisitExpressionList(node *ast.ExpressionList) {
	p.assertMatch(p.compiler.Visit(node).Result)
}

func (p *patternCompiler) VisitBinaryExpression(node *ast.BinaryExpressionNode) {
	p.assertMatch(p.compiler.Visit(node).Result)
}

func (p *patternCompiler) VisitUnaryExpression(node *ast.UnaryExpressionNode) {
	p.assertMatch(p.compiler.Visit(node).Result)
}

func (p *patternCompiler) VisitVector(node *ast.VectorNode) {
	// Proposed bytecode for:
	//    [1, 2, a, 3] = [1, 2, 3, 4]
	//
	// New opcodes needed:
	//   1. PUSH jmpTarget
	//      - Calculate `IP' by adding `jmpTarget' to `IP' and push to stack
	//   2. POP
	//      - Pop the last `IP' from the stack and jump to it
	//
	// JMP startMatch
	// .checkElement
	//     GET r, x, i
	//     ASSERT r, s
	//     ADD i, i, 1
	//     POP
	// .startMatch
	// TYPE r, x
	// ASSERT r, V_VEC
	// LEN r, x
	// LOADK i, len(node.Elements)
	// ASSERT r, i
	// LOADK i, 0
	// LOADK s, node.Elements[0] // c.Visit(node.Elements[0]).Result
	// PUSH 2
	// JMP checkElement
	// LOADK s, node.Elements[1] // c.Visit(node.Elements[1]).Result
	// PUSH 2
	// JMP checkElement
	// GET r2, x, i
	// ADD i, i, 1
	// LOADK s, node.Elements[3] // c.Visit(node.Elements[3]).Result
	// PUSH 2
	// JMP checkElement

	p.assertMatch(p.compiler.Visit(node).Result)
}

func (p *patternCompiler) VisitMap(node *ast.MapNode) {
	p.assertMatch(p.compiler.Visit(node).Result)
}

func (p *patternCompiler) VisitRecord(node *ast.RecordNode) {
	p.assertMatch(p.compiler.Visit(node).Result)
}

func (p *patternCompiler) VisitPair(node *ast.PairNode) {
	p.assertMatch(p.compiler.Visit(node).Result)
}

func (p *patternCompiler) VisitCall(node *ast.CallNode) {
	p.assertMatch(p.compiler.Visit(node).Result)
}

func (p *patternCompiler) VisitKeyAccess(node *ast.KeyAccessNode) {
	p.assertMatch(p.compiler.Visit(node).Result)
}

func (p *patternCompiler) VisitIfThenElse(node *ast.IfThenElseNode) {
	p.assertMatch(p.compiler.Visit(node).Result)
}

func (p *patternCompiler) VisitCaseExpression(node *ast.CaseExpressionNode) {
	p.assertMatch(p.compiler.Visit(node).Result)
}

func (p *patternCompiler) VisitCatchExpression(node *ast.CatchExpressionNode) {
	p.assertMatch(p.compiler.Visit(node).Result)
}

func (p *patternCompiler) VisitThrow(node *ast.ThrowNode) {
	p.assertMatch(p.compiler.Visit(node).Result)
}

func (p *patternCompiler) assertMatch(v Operand) {
	if p.compiler.Error != nil {
		return
	}

	p.compiler.ASM.Assert(p.value, v)
}
