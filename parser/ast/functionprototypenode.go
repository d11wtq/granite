package ast

// Functions
type FunctionPrototypeNode struct {
	Signature *VectorNode
	Body      *Collection
}

// Return a new FunctionPrototypeNode with signature and body.
func NewFunctionPrototypeNode(signature *VectorNode, body *Collection) *FunctionPrototypeNode {
	return &FunctionPrototypeNode{signature, body}
}

func (node *FunctionPrototypeNode) Accept(visitor ASTVisitor) {
	visitor.VisitFunctionPrototypeNode(node)
}
