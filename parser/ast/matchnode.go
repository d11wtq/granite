package ast

// Match construct
type MatchNode struct {
	Expr  ASTNode
	Cases []*MatchCaseNode
}

// Create a new match case list node.
func NewMatchNode(expr ASTNode, cases []*MatchCaseNode) *MatchNode {
	return &MatchNode{expr, cases}
}

func (node *MatchNode) Accept(visitor ASTVisitor) {
	visitor.VisitMatchNode(node)
}
