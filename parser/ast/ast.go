package ast

// Represents a single node in the parsed AST
type ASTNode interface {
	Accept(visitor ASTVisitor)
}

// Visitor that can traverse an AST
type ASTVisitor interface {
	VisitNil(node *NilNode)
	VisitInteger(node *IntegerNode)
	VisitBoolean(node *BooleanNode)
	VisitFloat(node *FloatNode)
	VisitString(node *StringNode)
	VisitSymbol(node *SymbolNode)
	VisitIdentifier(node *IdentifierNode)
	VisitExpressionList(node *ExpressionList)
	VisitBinaryExpression(node *BinaryExpressionNode)
	VisitUnaryExpression(node *UnaryExpressionNode)
	VisitVector(node *VectorNode)
	VisitMap(node *MapNode)
	VisitRecord(node *RecordNode)
	VisitPair(node *PairNode)
	VisitCall(node *CallNode)
	VisitKeyAccess(node *KeyAccessNode)
	VisitIfThenElse(node *IfThenElseNode)
	VisitCaseExpression(node *CaseExpressionNode)
	VisitCatchExpression(node *CatchExpressionNode)
	VisitThrow(node *ThrowNode)
}
