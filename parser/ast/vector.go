package ast

// Vector literals
type VectorNode struct {
	Elements []ASTNode
}

// Create a new Vector with elements.
func NewVector(elements ...ASTNode) *VectorNode {
	return &VectorNode{elements}
}

// Append the given element to this vector.
func (node *VectorNode) Append(element ASTNode) {
	node.Elements = append(node.Elements, element)
}

func (node *VectorNode) Accept(visitor ASTVisitor) {
	visitor.VisitVector(node)
}
