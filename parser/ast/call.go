package ast

// Function calls
type CallNode struct {
	Target    ASTNode
	Arguments ASTNode
}

// Create a new function application with target and args.
func NewCall(target ASTNode, args ASTNode) *CallNode {
	return &CallNode{target, args}
}

func (node *CallNode) Accept(visitor ASTVisitor) {
	visitor.VisitCall(node)
}
