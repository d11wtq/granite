package ast

// Symbols :symbol
type SymbolNode struct {
	Name string
}

func NewSymbolNode(name string) *SymbolNode {
	return &SymbolNode{name}
}

func (node *SymbolNode) Accept(visitor ASTVisitor) {
	visitor.VisitSymbolNode(node)
}
