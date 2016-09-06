package ast

// Records
type RecordPrototypeNode struct {
	Fields []*RecordFieldNode
}

// Return a new RecordPrototypeNode with fields.
func NewRecordPrototypeNode(fields ...*RecordFieldNode) *RecordPrototypeNode {
	return &RecordPrototypeNode{fields}
}

func (node *RecordPrototypeNode) Append(field *RecordFieldNode) {
	node.Fields = append(node.Fields, field)
}

func (node *RecordPrototypeNode) Accept(visitor ASTVisitor) {
	visitor.VisitRecordPrototypeNode(node)
}
