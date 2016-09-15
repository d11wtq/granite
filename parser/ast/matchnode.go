package ast

// Match construct
type MatchNode struct {
	Expr  ASTNode
	Cases []*PairNode
}

// Create a new match node.
func NewMatchNode(expr ASTNode, cases []*PairNode) *MatchNode {
	return &MatchNode{expr, cases}
}

func (node *MatchNode) Accept(visitor ASTVisitor) {
	visitor.VisitMatchNode(node)
}
