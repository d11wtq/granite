package ast

// Represents a single node in the parsed AST
type ASTNode interface {
	Accept(visitor ASTVisitor)
}

// Visitor that can traverse an AST
type ASTVisitor interface {
	VisitCollection(node *Collection)
	VisitIntegerNode(node *IntegerNode)
	VisitFloatNode(node *FloatNode)
	VisitStringNode(node *StringNode)
	VisitSymbolNode(node *SymbolNode)
	VisitIdentifierNode(node *IdentifierNode)
	VisitArithmeticNode(node *ArithmeticNode)
	VisitLogicalAndNode(node *LogicalAndNode)
	VisitLogicalOrNode(node *LogicalOrNode)
	VisitDefNode(node *DefNode)
	VisitAssignNode(node *AssignNode)
	VisitImportNode(node *ImportNode)
	VisitVectorNode(node *VectorNode)
	VisitMapNode(node *MapNode)
	VisitRecordPrototypeNode(node *RecordPrototypeNode)
	VisitFunctionPrototypeNode(node *FunctionPrototypeNode)
	VisitInvocationNode(node *InvocationNode)
	VisitMemberLookupNode(node *MemberLookupNode)
	VisitMatchNode(node *MatchNode)
	VisitSpreadNode(node *SpreadNode)
	VisitRecordNode(node *RecordNode)
}
