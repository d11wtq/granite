package ast

// Single Nil
type NilNode struct {
}

func (node *NilNode) Accept(visitor ASTVisitor) {
	visitor.VisitNil(node)
}

var Nil = &NilNode{}
