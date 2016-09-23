package ast

// Function calls
type FunctionApplicationNode struct {
	Target    ASTNode
	Arguments ASTNode
}

// Create a new function application with target and args.
func NewFunctionApplication(target ASTNode, args ASTNode) *FunctionApplicationNode {
	return &FunctionApplicationNode{target, args}
}

func (node *FunctionApplicationNode) Accept(visitor ASTVisitor) {
	visitor.VisitFunctionApplication(node)
}
