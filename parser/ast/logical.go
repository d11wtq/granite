package ast

const (
	OP_AND = iota
	OP_OR
)

// Logical 'and' or 'or' expression.
type LogicalNode struct {
	Op    int
	Left  ASTNode
	Right ASTNode
}

func NewLogicalNode(op int, left, right ASTNode) *LogicalNode {
	return &LogicalNode{op, left, right}
}

func (node *LogicalNode) Accept(visitor ASTVisitor) {
	visitor.VisitLogicalNode(node)
}
