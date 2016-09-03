package ast

// Record fields
type RecordFieldNode struct {
	Name  string
	Value ASTNode
}

// Return a new RecordFieldNode with name and value.
func NewRecordFieldNode(name string, value ASTNode) *RecordFieldNode {
	return &RecordFieldNode{name, value}
}

func (node *RecordFieldNode) Accept(visitor ASTVisitor) {
	// Nothing to do
}
