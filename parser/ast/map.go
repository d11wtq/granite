package ast

// Map literals
type MapNode struct {
	Elements []*PairNode
}

// Create a new map with elements.
func NewMap(elements ...*PairNode) *MapNode {
	return &MapNode{elements}
}

// Append the given Pair to this map.
func (node *MapNode) Append(pair *PairNode) {
	node.Elements = append(node.Elements, pair)
}

func (node *MapNode) Accept(visitor ASTVisitor) {
	visitor.VisitMap(node)
}
