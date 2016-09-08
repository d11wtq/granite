package ast

// Spreads
type SpreadNode struct {
	Expr ASTNode
}

func NewSpreadNode(expr ASTNode) *SpreadNode {
	return &SpreadNode{expr}
}

func (node *SpreadNode) Accept(visitor ASTVisitor) {
	visitor.VisitSpreadNode(node)
}
