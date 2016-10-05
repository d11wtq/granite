package ast

// Throw expressions
type ThrowNode struct {
	Expression ASTNode
}

// Create a new throw for expression.
func NewThrow(expression ASTNode) *ThrowNode {
	return &ThrowNode{expression}
}

func (node *ThrowNode) Accept(visitor ASTVisitor) {
	visitor.VisitThrow(node)
}
