package ast

// Record definitions
type RecordPrototypeNode struct {
	Fields []*KeyValueNode
}

// Return a new RecordPrototypeNode with fields.
func NewRecordPrototypeNode(fields []*KeyValueNode) *RecordPrototypeNode {
	return &RecordPrototypeNode{fields}
}

func (node *RecordPrototypeNode) Accept(visitor ASTVisitor) {
	visitor.VisitRecordPrototypeNode(node)
}
