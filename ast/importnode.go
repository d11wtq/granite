package ast

// Module imports
type ImportNode struct {
	Argument ASTNode
}

func (node *ImportNode) Accept(visitor ASTVisitor) {
	visitor.VisitImportNode(node)
}
