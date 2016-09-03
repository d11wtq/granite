package ast

// Records
type RecordNode struct {
	Fields []*RecordFieldNode
}

// Return a new RecordNode with fields.
func NewRecordNode(fields ...*RecordFieldNode) *RecordNode {
	return &RecordNode{fields}
}

func (node *RecordNode) Append(field *RecordFieldNode) {
	node.Fields = append(node.Fields, field)
}

func (node *RecordNode) Accept(visitor ASTVisitor) {
	visitor.VisitRecordNode(node)
}
