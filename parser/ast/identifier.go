package ast

// Variable names
type IdentifierNode struct {
	Name string
}

// Return a new identifier with name
func NewIdentifier(name string) *IdentifierNode {
	return &IdentifierNode{name}
}

func (node *IdentifierNode) Accept(visitor ASTVisitor) {
	visitor.VisitIdentifier(node)
}

// The special "blank" value
var Underscore = &IdentifierNode{"_"}
