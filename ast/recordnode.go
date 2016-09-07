package ast

// Records
type RecordNode struct {
	Prototype ASTNode
	Fields    []*KeyValueNode
}

// Return a new RecordNode with record and fields.
func NewRecordNode(record ASTNode, fields []*KeyValueNode) *RecordNode {
	return &RecordNode{record, fields}
}

func (node *RecordNode) Accept(visitor ASTVisitor) {
	visitor.VisitRecordNode(node)
}
