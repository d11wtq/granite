package ast

// Pattern matching assignments
type AssignNode struct {
	Left  ASTNode
	Right ASTNode
}

// Create a new assignment node.
func NewAssignNode(left, right ASTNode) *AssignNode {
	return &AssignNode{left, right}
}

func (node *AssignNode) Accept(visitor ASTVisitor) {
	visitor.VisitAssignNode(node)
}
