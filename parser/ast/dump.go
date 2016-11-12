package ast

import (
	"fmt"
)

var opNames = map[int]string{
	OP_ADD:     "+",
	OP_MIN:     "-",
	OP_MUL:     "*",
	OP_DIV:     "/",
	OP_MOD:     "%",
	OP_ASS:     "=",
	OP_EQL:     "==",
	OP_GT:      ">",
	OP_LT:      "<",
	OP_GTE:     ">=",
	OP_LTE:     "<=",
	OP_AND:     "and",
	OP_OR:      "or",
	OP_NOT:     "not",
	OP_BIT_AND: "&",
	OP_BIT_OR:  "|",
	OP_BIT_XOR: "^",
	OP_BIT_NOT: "~",
	OP_BIT_SR:  ">>",
	OP_BIT_SL:  "<<",
}

// The Dumper can render an AST.
type Dumper struct {
	margin string
}

// Visit the given ASTNode.
// This also handles nil values.
func (d *Dumper) Visit(node ASTNode) {
	if node == nil {
		d.render("(nil)")
		return
	}

	node.Accept(d)
}

func (d *Dumper) VisitNil(node *NilNode) {
	d.render("NIL")
}

func (d *Dumper) VisitInteger(node *IntegerNode) {
	d.render(fmt.Sprintf("INTEGER (%d)", node.Value))
}

func (d *Dumper) VisitBoolean(node *BooleanNode) {
	d.render(fmt.Sprintf("BOOLEAN (%#v)", node.Value))
}

func (d *Dumper) VisitFloat(node *FloatNode) {
	d.render(fmt.Sprintf("FLOAT (%#v)", node.Value))
}

func (d *Dumper) VisitString(node *StringNode) {
	d.render(fmt.Sprintf("STRING (%#v)", node.String))
}

func (d *Dumper) VisitSymbol(node *SymbolNode) {
	d.render(fmt.Sprintf("SYMBOL (%s)", node.Name))
}

func (d *Dumper) VisitIdentifier(node *IdentifierNode) {
	d.render(fmt.Sprintf("IDENT (%s)", node.Name))
}

func (d *Dumper) VisitExpressionList(node *ExpressionList) {
	d.render(
		fmt.Sprintf("expressions (%d nodes)", len(node.Elements)),
		node.Elements...,
	)
}

func (d *Dumper) VisitBinaryExpression(node *BinaryExpressionNode) {
	d.render(
		fmt.Sprintf("binary expression (%s)", opNames[node.Op]),
		node.Left,
		node.Right,
	)
}

func (d *Dumper) VisitUnaryExpression(node *UnaryExpressionNode) {
	d.render(
		fmt.Sprintf("unary expression (%s)", opNames[node.Op]),
		node.Expression,
	)
}

func (d *Dumper) VisitVector(node *VectorNode) {
	d.render(
		fmt.Sprintf("vector (%d nodes)", len(node.Elements)),
		node.Elements...,
	)
}

func (d *Dumper) VisitMap(node *MapNode) {
	elements := make([]ASTNode, 0)
	for _, e := range node.Elements {
		elements = append(elements, e)
	}

	d.render(
		fmt.Sprintf("map (%d nodes)", len(node.Elements)),
		elements...,
	)
}

func (d *Dumper) VisitRecord(node *RecordNode) {
	d.render("record", node.Type, NewMap(node.Elements...))
}

func (d *Dumper) VisitPair(node *PairNode) {
	d.render("pair of", node.Key, node.Value)
}

func (d *Dumper) VisitCall(node *CallNode) {
	d.render("function call", node.Target, node.Arguments)
}

func (d *Dumper) VisitKeyAccess(node *KeyAccessNode) {
	d.render("key access", node.Target, node.Key)
}

func (d *Dumper) VisitIfThenElse(node *IfThenElseNode) {
	d.render("if-then-else", node.Expression, node.Then, node.Else)
}

func (d *Dumper) VisitCaseExpression(node *CaseExpressionNode) {
	d.render("case of", node.Expression, NewMap(node.Cases...))
}

func (d *Dumper) VisitCatchExpression(node *CatchExpressionNode) {
	d.render("catch of", node.Expression, NewMap(node.Cases...))
}

func (d *Dumper) VisitThrow(node *ThrowNode) {
	d.render("throw", node.Expression)
}

func (d *Dumper) render(name string, children ...ASTNode) {
	lastNode := len(children) - 1
	fmt.Println(name)
	for i, node := range children {
		d.renderBranch()
		d.indent(i < lastNode).Visit(node)
	}
}

func (d *Dumper) renderBranch() {
	fmt.Print(d.margin, " |", "\n")
	fmt.Print(d.margin, " •——")
}

func (d *Dumper) indent(continueBranch bool) *Dumper {
	margins := map[bool]string{
		true:  " |  ",
		false: "    ",
	}
	return &Dumper{d.margin + margins[continueBranch]}
}

// Return a new AST Dumper.
func NewDumper() *Dumper {
	return &Dumper{}
}

// Render the AST to os.Stdout.
func Dump(node ASTNode) {
	NewDumper().Visit(node)
}
