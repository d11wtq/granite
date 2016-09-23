package ast

// Key value pairs
type PairNode struct {
	Key   ASTNode
	Value ASTNode
}

// Create a new Pair.
func NewPair(key, value ASTNode) *PairNode {
	return &PairNode{key, value}
}

func (node *PairNode) Accept(visitor ASTVisitor) {
	visitor.VisitPair(node)
}
