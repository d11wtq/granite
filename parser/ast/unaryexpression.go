package ast

// Unary expressions
type UnaryExpressionNode struct {
	Op         int
	Expression ASTNode
}

// Create a new unary expression op and expression.
func NewUnaryExpression(op int, expression ASTNode) *UnaryExpressionNode {
	return &UnaryExpressionNode{op, expression}
}

func (node *UnaryExpressionNode) Accept(visitor ASTVisitor) {
	visitor.VisitUnaryExpression(node)
}
