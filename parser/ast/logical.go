package ast

// Logical 'and' expression.
type LogicalAndNode struct {
	Left  ASTNode
	Right ASTNode
}

func NewLogicalAndNode(left, right ASTNode) *LogicalAndNode {
	return &LogicalAndNode{left, right}
}

func (node *LogicalAndNode) Accept(visitor ASTVisitor) {
	visitor.VisitLogicalAndNode(node)
}

// Logical 'or' expression.
type LogicalOrNode struct {
	Left  ASTNode
	Right ASTNode
}

func NewLogicalOrNode(left, right ASTNode) *LogicalOrNode {
	return &LogicalOrNode{left, right}
}

func (node *LogicalOrNode) Accept(visitor ASTVisitor) {
	visitor.VisitLogicalOrNode(node)
}
