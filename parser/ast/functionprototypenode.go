package ast

// Functions
type FunctionPrototypeNode struct {
	Cases []*PairNode
}

// Return a new FunctionPrototypeNode with the given cases.
func NewFunctionPrototypeNode(cases []*PairNode) *FunctionPrototypeNode {
	return &FunctionPrototypeNode{cases}
}

func (node *FunctionPrototypeNode) Accept(visitor ASTVisitor) {
	visitor.VisitFunctionPrototypeNode(node)
}
