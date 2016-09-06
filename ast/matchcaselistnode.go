package ast

// List of cases in the match construct
type MatchCaseListNode struct {
	Cases []*MatchCaseNode
}

// Create a new match case list node.
func NewMatchCaseListNode(cases ...*MatchCaseNode) *MatchCaseListNode {
	return &MatchCaseListNode{cases}
}

func (node *MatchCaseListNode) Append(case_ *MatchCaseNode) {
	node.Cases = append(node.Cases, case_)
}

func (node *MatchCaseListNode) Accept(visitor ASTVisitor) {
	// Noop
}
