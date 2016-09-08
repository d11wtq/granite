package ast

// Module imports
type ImportNode struct {
	Path ASTNode
}

// Create a new ImportNode for path.
func NewImportNode(path ASTNode) *ImportNode {
	return &ImportNode{path}
}

func (node *ImportNode) Accept(visitor ASTVisitor) {
	visitor.VisitImportNode(node)
}
