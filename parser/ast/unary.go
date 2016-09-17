package ast

// Unary expression
type UnaryNode struct {
	Op   int
	Expr ASTNode
}

func NewUnaryNode(op int, expr ASTNode) *UnaryNode {
	return &UnaryNode{op, expr}
}

func (node *UnaryNode) Accept(visitor ASTVisitor) {
	visitor.VisitUnaryNode(node)
}
