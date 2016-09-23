package ast

// Arbitrary series of expressions
type ExpressionList struct {
	Elements []ASTNode
}

// Create a new expression list with elements.
func NewExpressionList(elements ...ASTNode) *ExpressionList {
	return &ExpressionList{elements}
}

// Append the given element to this list.
func (c *ExpressionList) Append(element ASTNode) {
	c.Elements = append(c.Elements, element)
}

func (c *ExpressionList) Accept(visitor ASTVisitor) {
	visitor.VisitExpressionList(c)
}
