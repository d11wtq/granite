package ast

import (
	"fmt"
)

// Internal use only
type sliceNode struct {
	Name     string
	Elements []ASTNode
}

// Create a new slice node for rendering in the AST
func newSliceNode(name string, elements []ASTNode) *sliceNode {
	return &sliceNode{name, elements}
}

// Dangerously asserts visitor as a *Dumper and invokes visitSliceNode()
func (node *sliceNode) Accept(visitor ASTVisitor) {
	visitor.(*Dumper).visitSliceNode(node)
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

func (d *Dumper) VisitCollection(node *Collection) {
	d.render(
		fmt.Sprintf("statements (%d nodes)", len(node.Elements)),
		node.Elements...,
	)
}

func (d *Dumper) VisitIntegerNode(node *IntegerNode) {
	d.render(fmt.Sprintf("INTEGER (%d)", node.Value))
}

func (d *Dumper) VisitBooleanNode(node *BooleanNode) {
	d.render(fmt.Sprintf("BOOLEAN (%#v)", node.Value))
}

func (d *Dumper) VisitFloatNode(node *FloatNode) {
	d.render(fmt.Sprintf("FLOAT (%f)", node.Value))
}

func (d *Dumper) VisitStringNode(node *StringNode) {
	d.render(fmt.Sprintf("STRING (%#v)", node.String))
}

func (d *Dumper) VisitSymbolNode(node *SymbolNode) {
	d.render(fmt.Sprintf("SYMBOL (%s)", node.Name))
}

func (d *Dumper) VisitIdentifierNode(node *IdentifierNode) {
	d.render(fmt.Sprintf("IDENT (%s)", node.Name))
}

func (d *Dumper) VisitPairNode(node *PairNode) {
	d.render("pair of", node.Left, node.Right)
}

func (d *Dumper) VisitArithmeticNode(node *ArithmeticNode) {
	d.render(fmt.Sprintf("math (%c)", node.Op), node.Left, node.Right)
}

func (d *Dumper) VisitComparisonNode(node *ComparisonNode) {
	operators := map[int]string{
		CMP_GT: ">",
		CMP_LT: "<",
	}
	d.render(
		fmt.Sprintf("comparison (%s)", operators[node.Cmp]),
		node.Left,
		node.Right,
	)
}

func (d *Dumper) VisitLogicalNode(node *LogicalNode) {
	operators := map[int]string{
		OP_AND: "and",
		OP_OR:  "or",
	}
	d.render(
		fmt.Sprintf("logical (%s)", operators[node.Op]),
		node.Left,
		node.Right,
	)
}

func (d *Dumper) VisitDefNode(node *DefNode) {
	d.render("define", node.Name, node.Value)
}

func (d *Dumper) VisitAssignNode(node *AssignNode) {
	d.render("assignment (=)", node.Left, node.Right)
}

func (d *Dumper) VisitImportNode(node *ImportNode) {
	d.render("import", node.Path)
}

func (d *Dumper) VisitVectorNode(node *VectorNode) {
	d.render(
		fmt.Sprintf("vector (%d nodes)", len(node.Elements)),
		node.Elements...,
	)
}

func (d *Dumper) VisitMapNode(node *MapNode) {
	pairs := make([]ASTNode, 0)
	for _, pair := range node.Pairs {
		pairs = append(pairs, pair)
	}
	d.render(fmt.Sprintf("map (%d nodes)", len(node.Pairs)), pairs...)
}

func (d *Dumper) VisitRecordPrototypeNode(node *RecordPrototypeNode) {
	fields := make([]ASTNode, 0)
	for _, field := range node.Fields {
		fields = append(fields, field)
	}
	d.render(
		fmt.Sprintf("record prototype (%d nodes)", len(node.Fields)),
		fields...,
	)
}

func (d *Dumper) VisitFunctionPrototypeNode(node *FunctionPrototypeNode) {
	cases := make([]ASTNode, 0)
	for _, case_ := range node.Cases {
		cases = append(cases, case_)
	}
	d.render(
		fmt.Sprintf("function prototype (%d nodes)", len(node.Cases)),
		cases...,
	)
}

func (d *Dumper) VisitInvocationNode(node *InvocationNode) {
	d.render("invocation", node.Prototype, node.Args)
}

func (d *Dumper) VisitMemberLookupNode(node *MemberLookupNode) {
	d.render("member lookup ([])", node.Target, node.Key)
}

func (d *Dumper) VisitMatchNode(node *MatchNode) {
	cases := make([]ASTNode, 0)
	for _, case_ := range node.Cases {
		cases = append(cases, case_)
	}
	d.render(
		"match construct",
		node.Expr,
		newSliceNode(fmt.Sprintf("cases (%d nodes)", len(node.Cases)), cases),
	)
}

func (d *Dumper) VisitSpreadNode(node *SpreadNode) {
	d.render("spread of", node.Expr)
}

func (d *Dumper) VisitRecordNode(node *RecordNode) {
	fields := make([]ASTNode, 0)
	for _, field := range node.Fields {
		fields = append(fields, field)
	}
	d.render(
		"build record",
		node.Prototype,
		newSliceNode(fmt.Sprintf("fields (%d nodes)", len(node.Fields)), fields),
	)
}

func (d *Dumper) visitSliceNode(node *sliceNode) {
	d.render(node.Name, node.Elements...)
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
