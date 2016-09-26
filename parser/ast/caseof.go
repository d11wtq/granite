package ast

// Case of expressions
type CaseExpressionNode struct {
	Expression ASTNode
	Cases      []*PairNode
}

// Create a case expr with expeession and cases.
func NewCaseExpression(expr ASTNode, cases ...*PairNode) *CaseExpressionNode {
	return &CaseExpressionNode{expr, cases}
}

func (node *CaseExpressionNode) Accept(visitor ASTVisitor) {
	visitor.VisitCaseExpression(node)
}
