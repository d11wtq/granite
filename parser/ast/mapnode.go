package ast

// Map literal
type MapNode struct {
	KeyValues []*KeyValueNode
}

// Create a new MapNode with the given keys.
func NewMapNode(keyvalues ...*KeyValueNode) *MapNode {
	return &MapNode{keyvalues}
}

func (node *MapNode) Append(keyvalue *KeyValueNode) {
	node.KeyValues = append(node.KeyValues, keyvalue)
}

func (node *MapNode) Accept(visitor ASTVisitor) {
	visitor.VisitMapNode(node)
}
