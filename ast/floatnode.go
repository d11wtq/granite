package ast

// Floating point literals
type FloatNode struct {
	Value float64
}

func (node *FloatNode) Accept(visitor ASTVisitor) {
	visitor.VisitFloatNode(node)
}
