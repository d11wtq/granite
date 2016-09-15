package ast

// Records
type RecordNode struct {
	Prototype ASTNode
	Fields    []*PairNode
}

// Return a new RecordNode with prototype and fields.
func NewRecordNode(prototype ASTNode, fields []*PairNode) *RecordNode {
	return &RecordNode{prototype, fields}
}

func (node *RecordNode) Accept(visitor ASTVisitor) {
	visitor.VisitRecordNode(node)
}
