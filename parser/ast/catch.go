package ast

// Catch expressions
type CatchExpressionNode struct {
	Expression ASTNode
	Cases      []*PairNode
}

// Create a catch expr with expeession and cases.
func NewCatchExpression(expr ASTNode, cases ...*PairNode) *CatchExpressionNode {
	return &CatchExpressionNode{expr, cases}
}

func (node *CatchExpressionNode) Accept(visitor ASTVisitor) {
	visitor.VisitCatchExpression(node)
}
