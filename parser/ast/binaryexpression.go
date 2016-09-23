package ast

// Binary expressions
type BinaryExpressionNode struct {
	Op    int
	Left  ASTNode
	Right ASTNode
}

// Create a new binary expression list op, left and right.
func NewBinaryExpression(op int, left, right ASTNode) *BinaryExpressionNode {
	return &BinaryExpressionNode{op, left, right}
}

func (node *BinaryExpressionNode) Accept(visitor ASTVisitor) {
	visitor.VisitBinaryExpression(node)
}
