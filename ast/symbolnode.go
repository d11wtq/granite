package ast

// Symbols :symbol
type SymbolNode struct {
	Name string
}

func (node *SymbolNode) Accept(visitor ASTVisitor) {
	visitor.VisitSymbolNode(node)
}
