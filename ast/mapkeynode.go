package ast

// Map keys and values
type MapKeyNode struct {
	Key ASTNode
	Val ASTNode
}

func (node *MapKeyNode) Accept(visitor ASTVisitor) {
	// Nothing to do
}
