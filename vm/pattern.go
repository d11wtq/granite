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
	if v, exists := p.compiler.Symbols.GetLocal(node.Name); exists {
		p.assertMatch(v)
		return
	}

	p.compiler.Symbols.SetLocal(node.Name, p.value)
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
	var (
		tmp = p.compiler.tempVar()
		idx = p.compiler.tempVar()
	)

	p.compiler.ASM.Type(tmp, p.value)
	p.compiler.ASM.Assert(
		p.compiler.loadConstant(Integer(V_VEC)).Result,
		tmp,
		p.value,
	)
	p.compiler.ASM.Len(tmp, p.value)
	p.compiler.ASM.Assert(
		p.compiler.loadConstant(Integer(len(node.Elements))).Result,
		tmp,
		p.value,
	)

	for j, elem := range node.Elements {
		tmp = p.compiler.tempVar()

		if j == 0 {
			p.compiler.ASM.LoadK(idx, &Constant{Integer(0)})
		} else {
			p.compiler.ASM.Add(
				idx,
				idx,
				p.compiler.loadConstant(Integer(1)).Result,
			)
		}

		p.compiler.ASM.Get(tmp, p.value, idx)
		compilePattern(p.compiler, elem, tmp)
	}
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

	p.compiler.ASM.Assert(p.value, v, p.value)
}
