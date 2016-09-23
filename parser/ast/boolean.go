package ast

// Boolean true and false
type BooleanNode struct {
	Value bool
}

func (node *BooleanNode) Accept(visitor ASTVisitor) {
	visitor.VisitBoolean(node)
}

var (
	TrueNode  = &BooleanNode{true}
	FalseNode = &BooleanNode{false}
)
