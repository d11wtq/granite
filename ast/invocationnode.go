package ast

// Function invocations
type InvocationNode struct {
	Prototype ASTNode
	Args      *VectorNode
}

// Return a new InvocationNode with prototype and args.
func NewInvocationNode(prototype ASTNode, args *VectorNode) *InvocationNode {
	return &InvocationNode{prototype, args}
}

func (node *InvocationNode) Accept(visitor ASTVisitor) {
	visitor.VisitInvocationNode(node)
}
