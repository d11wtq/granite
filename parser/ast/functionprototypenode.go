package ast

// Functions
type FunctionPrototypeNode struct {
	Cases []*MatchCaseNode
}

// Return a new FunctionPrototypeNode wit cases.
func NewFunctionPrototypeNode(cases []*MatchCaseNode) *FunctionPrototypeNode {
	return &FunctionPrototypeNode{cases}
}

func (node *FunctionPrototypeNode) Accept(visitor ASTVisitor) {
	visitor.VisitFunctionPrototypeNode(node)
}
