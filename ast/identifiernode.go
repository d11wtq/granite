package ast

// Bare identifiers
type IdentifierNode struct {
	Name string
}

func (node *IdentifierNode) Accept(visitor ASTVisitor) {
	visitor.VisitIdentifierNode(node)
}
