package ast

// Floating point literals
type FloatNode struct {
	Value float64
}

// Create a new Float node
func NewFloatNode(value float64) *FloatNode {
	return &FloatNode{value}
}

func (node *FloatNode) Accept(visitor ASTVisitor) {
	visitor.VisitFloatNode(node)
}
