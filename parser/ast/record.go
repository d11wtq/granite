package ast

// Record literals
type RecordNode struct {
	Type     ASTNode
	Elements []*PairNode
}

// Create a record with type and elements.
func NewRecord(type_ ASTNode, elements ...*PairNode) *RecordNode {
	return &RecordNode{type_, elements}
}

func (node *RecordNode) Accept(visitor ASTVisitor) {
	visitor.VisitRecord(node)
}
