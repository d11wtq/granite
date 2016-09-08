package ast

// Vector literal
type VectorNode struct {
	Elements []ASTNode
}

func NewVectorNode(elements ...ASTNode) *VectorNode {
	return &VectorNode{elements}
}

func (node *VectorNode) Append(element ASTNode) {
	node.Elements = append(node.Elements, element)
}

func (node *VectorNode) Accept(visitor ASTVisitor) {
	visitor.VisitVectorNode(node)
}
