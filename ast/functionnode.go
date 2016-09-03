package ast

// Functions
type FunctionNode struct {
	Signature *VectorNode
	Body      *Collection
}

// Return a new FunctionNode with signature and body.
func NewFunctionNode(signature *VectorNode, body *Collection) *FunctionNode {
	return &FunctionNode{signature, body}
}

func (node *FunctionNode) Accept(visitor ASTVisitor) {
	visitor.VisitFunctionNode(node)
}
