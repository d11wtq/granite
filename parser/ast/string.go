package ast

// String literals
type StringNode struct {
	String string
}

// Create a new string node for s
func NewString(s string) *StringNode {
	return &StringNode{s}
}

func (node *StringNode) Accept(visitor ASTVisitor) {
	visitor.VisitString(node)
}
