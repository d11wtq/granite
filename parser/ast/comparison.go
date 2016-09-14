package ast

const (
	CMP_GT = iota
	CMP_LT
)

// Comparisons
type ComparisonNode struct {
	Cmp   int
	Left  ASTNode
	Right ASTNode
}

func NewComparisonNode(cmp int, left, right ASTNode) *ComparisonNode {
	return &ComparisonNode{cmp, left, right}
}

func (node *ComparisonNode) Accept(visitor ASTVisitor) {
	visitor.VisitComparisonNode(node)
}
