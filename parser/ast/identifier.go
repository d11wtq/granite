package ast

// Variable names
type IdentifierNode struct {
	Name string
}

// The special "blank" value
var Underscore *IdentifierNode

func init() {
	Underscore = NewIdentifier("_")
}

// Return a new identifier with name
func NewIdentifier(name string) *IdentifierNode {
	return &IdentifierNode{name}
}

func (node *IdentifierNode) Accept(visitor ASTVisitor) {
	visitor.VisitIdentifier(node)
}

func (node *IdentifierNode) IsWilcard() bool {
	for _, r := range node.Name {
		if r != '_' {
			return false
		}
	}

	return true
}
