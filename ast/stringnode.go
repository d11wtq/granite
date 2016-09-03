package ast

// String literals
type StringNode struct {
	String string
}

func (node *StringNode) Accept(visitor ASTVisitor) {
	visitor.VisitStringNode(node)
}
