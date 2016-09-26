package ast

// If then else expressions
type IfThenElseNode struct {
	Expression ASTNode
	Then       ASTNode
	Else       ASTNode
}

// Create a new if-then-else.
func NewIfThenElseExpression(expression, then, else_ ASTNode) *IfThenElseNode {
	return &IfThenElseNode{expression, then, else_}
}

func (node *IfThenElseNode) Accept(visitor ASTVisitor) {
	visitor.VisitIfThenElse(node)
}
