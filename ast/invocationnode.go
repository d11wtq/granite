package ast

// Function invocations
type InvocationNode struct {
	Target ASTNode
	Args   *VectorNode
}

// Return a new InvocationNode with target and args.
func NewInvocationNode(target ASTNode, args *VectorNode) *InvocationNode {
	return &InvocationNode{target, args}
}

func (node *InvocationNode) Accept(visitor ASTVisitor) {
	visitor.VisitInvocationNode(node)
}
