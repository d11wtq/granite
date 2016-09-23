package ast

// Keyed member access.
type KeyAccessNode struct {
	Target ASTNode
	Key    ASTNode
}

// Create a new key access with target and key.
func NewKeyAccess(target, key ASTNode) *KeyAccessNode {
	return &KeyAccessNode{target, key}
}

func (node *KeyAccessNode) Accept(visitor ASTVisitor) {
	visitor.VisitKeyAccess(node)
}
