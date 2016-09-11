package ast

// Bare identifiers
type IdentifierNode struct {
	Name string
}

// Return a new identifier with name
func NewIdentifierNode(name string) *IdentifierNode {
	return &IdentifierNode{name}
}

func (node *IdentifierNode) Accept(visitor ASTVisitor) {
	visitor.VisitIdentifierNode(node)
}
