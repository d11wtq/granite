package ast

// Integer literals
type IntegerNode struct {
	Value int64
}

func (node *IntegerNode) Accept(visitor ASTVisitor) {
	visitor.VisitIntegerNode(node)
}
