package ast

// Map/record keys and values
type KeyValueNode struct {
	Key   ASTNode
	Value ASTNode
}

func NewKeyValueNode(key, value ASTNode) *KeyValueNode {
	return &KeyValueNode{key, value}
}

func (node *KeyValueNode) IsPositional() bool {
	return node.Key == nil
}

func (node *KeyValueNode) Accept(visitor ASTVisitor) {
	// Nothing to do
}
