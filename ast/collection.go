package ast

// Transient data structure used during parsing.
type Collection struct {
	Elements []ASTNode
}

// Create a new collection with elements.
func NewCollection(elements ...ASTNode) *Collection {
	return &Collection{make([]ASTNode, 0)}
}

// Append the given element to this collection.
func (c *Collection) Append(element ASTNode) {
	c.Elements = append(c.Elements, element)
}

func (c *Collection) Accept(visitor ASTVisitor) {
	visitor.VisitCollection(c)
}
