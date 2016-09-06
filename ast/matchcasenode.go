package ast

// Single 'case' in the match construct
type MatchCaseNode struct {
	When ASTNode
	Then ASTNode
}

// Create a new match case node.
func NewMatchCaseNode(when, then ASTNode) *MatchCaseNode {
	return &MatchCaseNode{when, then}
}

func (node *MatchCaseNode) Accept(visitor ASTVisitor) {
	// Noop
}
