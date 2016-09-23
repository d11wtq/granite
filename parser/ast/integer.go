package ast

// Integer literals
type IntegerNode struct {
	Value int64
}

// Create a new integer node with value
func NewInteger(value int64) *IntegerNode {
	return &IntegerNode{value}
}

func (node *IntegerNode) Accept(visitor ASTVisitor) {
	visitor.VisitInteger(node)
}
