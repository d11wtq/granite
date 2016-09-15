package ast

// Arbitrary associative data
type PairNode struct {
	Left  ASTNode
	Right ASTNode
}

func NewPairNode(left, right ASTNode) *PairNode {
	return &PairNode{left, right}
}

func (node *PairNode) IsPositional() bool {
	return node.Left == nil
}

func (node *PairNode) Accept(visitor ASTVisitor) {
	visitor.VisitPairNode(node)
}
