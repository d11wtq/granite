package ast

// Record definitions
type RecordPrototypeNode struct {
	Fields []*PairNode
}

// Return a new RecordPrototypeNode with fields.
func NewRecordPrototypeNode(fields []*PairNode) *RecordPrototypeNode {
	return &RecordPrototypeNode{fields}
}

func (node *RecordPrototypeNode) Accept(visitor ASTVisitor) {
	visitor.VisitRecordPrototypeNode(node)
}
