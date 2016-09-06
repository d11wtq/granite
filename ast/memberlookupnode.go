package ast

// Data structure member access
type MemberLookupNode struct {
	Target ASTNode
	Key    ASTNode
}

// Create a new member lookup node.
func NewMemberLookupNode(target, key ASTNode) *MemberLookupNode {
	return &MemberLookupNode{target, key}
}

func (node *MemberLookupNode) Accept(visitor ASTVisitor) {
	visitor.VisitMemberLookupNode(node)
}
