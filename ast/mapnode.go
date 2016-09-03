package ast

// Map literal
type MapNode struct {
	KeyValues []*MapKeyNode
}

// Create a new MapNode with the given keys.
func NewMapNode(keyvalues ...*MapKeyNode) *MapNode {
	return &MapNode{keyvalues}
}

func (node *MapNode) Append(keyvalue *MapKeyNode) {
	node.KeyValues = append(node.KeyValues, keyvalue)
}

func (node *MapNode) Accept(visitor ASTVisitor) {
	visitor.VisitMapNode(node)
}
