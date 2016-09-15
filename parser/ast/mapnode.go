package ast

// Map literal
type MapNode struct {
	Pairs []*PairNode
}

// Create a new MapNode with the given pairs.
func NewMapNode(pairs ...*PairNode) *MapNode {
	return &MapNode{pairs}
}

func (node *MapNode) Append(pair *PairNode) {
	node.Pairs = append(node.Pairs, pair)
}

func (node *MapNode) Accept(visitor ASTVisitor) {
	visitor.VisitMapNode(node)
}
