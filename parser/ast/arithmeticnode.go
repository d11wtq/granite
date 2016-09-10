package ast

// Arithmetic calculation
type ArithmeticNode struct {
	Op    int
	Left  ASTNode
	Right ASTNode
}

func NewArithmeticNode(op int, left, right ASTNode) *ArithmeticNode {
	return &ArithmeticNode{op, left, right}
}

func (node *ArithmeticNode) Accept(visitor ASTVisitor) {
	visitor.VisitArithmeticNode(node)
}
