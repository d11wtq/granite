package ast

// Variable names
type IdentifierNode struct {
	Name     string
	Register uint8
}

// Return a new identifier with name
func NewIdentifier(name string) *IdentifierNode {
	return &IdentifierNode{name, 0}
}

func (node *IdentifierNode) Accept(visitor ASTVisitor) {
	visitor.VisitIdentifier(node)
}

// The special "blank" value
var Underscore = &IdentifierNode{"_", 0}
