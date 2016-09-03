package ast

// Definitions (functions, records)
type DefNode struct {
	Name  *IdentifierNode
	Value ASTNode
}

// Create a new definition node.
func NewDefNode(name *IdentifierNode, value ASTNode) *DefNode {
	return &DefNode{name, value}
}

func (node *DefNode) Accept(visitor ASTVisitor) {
	visitor.VisitDefNode(node)
}
