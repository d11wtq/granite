package ast

import (
	"fmt"
)

// The Inspector can inspect an AST.
type Inspector struct {
	// Left padding on output lines
	pad string
	// Lazy-loaded child inspector
	child *Inspector
}

// Visit the given ASTNode.
// This also handles nil values.
func (i *Inspector) Visit(node ASTNode) {
	if node == nil {
		i.dump("<nil>")
		return
	}

	node.Accept(i)
}

func (i *Inspector) VisitCollection(node *Collection) {
	i.dump("statements (%d):", len(node.Elements))
	for _, stmt := range node.Elements {
		i.nested().Visit(stmt)
	}
}

func (i *Inspector) VisitIntegerNode(node *IntegerNode) {
	i.dump("INTEGER (%#v)", node.Value)
}

func (i *Inspector) VisitFloatNode(node *FloatNode) {
	i.dump("FLOAT (%#v)", node.Value)
}

func (i *Inspector) VisitStringNode(node *StringNode) {
	i.dump("STRING (%#v)", node.String)
}

func (i *Inspector) VisitSymbolNode(node *SymbolNode) {
	i.dump("SYMBOL (%s)", node.Name)
}

func (i *Inspector) VisitIdentifierNode(node *IdentifierNode) {
	i.dump("IDENT (%s)", node.Name)
}

func (i *Inspector) VisitArithmeticNode(node *ArithmeticNode) {
	i.dump("math (%c):", node.Op)
	i.nested().dump("left:")
	i.nested().nested().Visit(node.Left)
	i.nested().dump("right:")
	i.nested().nested().Visit(node.Right)
}

func (i *Inspector) VisitLogicalAndNode(node *LogicalAndNode) {
	i.dump("and:")
	i.nested().dump("left:")
	i.nested().nested().Visit(node.Left)
	i.nested().dump("right:")
	i.nested().nested().Visit(node.Right)
}

func (i *Inspector) VisitLogicalOrNode(node *LogicalOrNode) {
	i.dump("or:")
	i.nested().dump("left:")
	i.nested().nested().Visit(node.Left)
	i.nested().dump("right:")
	i.nested().nested().Visit(node.Right)
}

func (i *Inspector) VisitDefNode(node *DefNode) {
	i.dump("define (%s):", node.Name.Name)
	i.nested().Visit(node.Value)
}

func (i *Inspector) VisitAssignNode(node *AssignNode) {
	i.dump("assignment (=):")
	i.nested().dump("left:")
	i.nested().nested().Visit(node.Left)
	i.nested().dump("right:")
	i.nested().nested().Visit(node.Right)
}

func (i *Inspector) VisitImportNode(node *ImportNode) {
	i.dump("import:")
	i.nested().Visit(node.Path)
}

func (i *Inspector) VisitVectorNode(node *VectorNode) {
	i.dump("vector (%d)", len(node.Elements))
	for k, v := range node.Elements {
		i.nested().dump("- item (%d):", k)
		i.nested().nested().nested().Visit(v)
	}
}

func (i *Inspector) VisitMapNode(node *MapNode) {
	i.dump("map (%d):", len(node.KeyValues))
	for _, kv := range node.KeyValues {
		i.nested().dump("- key:")
		i.nested().nested().nested().Visit(kv.Key)
		i.nested().dump("  item:")
		i.nested().nested().nested().Visit(kv.Value)
	}
}

func (i *Inspector) VisitRecordPrototypeNode(node *RecordPrototypeNode) {
	i.dump("record prototype (%d):", len(node.Fields))
	for _, kv := range node.Fields {
		i.nested().dump("- field:")
		i.nested().nested().nested().Visit(kv.Key)
		i.nested().nested().dump("default:")
		i.nested().nested().nested().Visit(kv.Value)
	}
}

func (i *Inspector) VisitFunctionPrototypeNode(node *FunctionPrototypeNode) {
	i.dump("function prototype:")
	i.nested().dump("signature:")
	i.nested().nested().Visit(node.Signature)
	i.nested().dump("body:")
	i.nested().nested().Visit(node.Body)
}

func (i *Inspector) VisitInvocationNode(node *InvocationNode) {
	i.dump("invocation:")
	i.nested().dump("target:")
	i.nested().nested().Visit(node.Prototype)
	i.nested().dump("arguments:")
	i.nested().nested().Visit(node.Args)
}

func (i *Inspector) VisitMemberLookupNode(node *MemberLookupNode) {
	i.dump("lookup:")
	i.nested().dump("target:")
	i.nested().nested().Visit(node.Target)
	i.nested().dump("key:")
	i.nested().nested().Visit(node.Key)
}

func (i *Inspector) VisitMatchNode(node *MatchNode) {
	i.dump("match:")
	i.nested().dump("argument:")
	i.nested().nested().Visit(node.Expr)

	i.nested().dump("cases (%d):", len(node.Cases))
	for _, case_ := range node.Cases {
		i.nested().nested().dump("- pattern:")
		i.nested().nested().nested().nested().Visit(case_.When)
		i.nested().nested().nested().dump("then:")
		i.nested().nested().nested().nested().Visit(case_.Then)
	}
}

func (i *Inspector) VisitSpreadNode(node *SpreadNode) {
	i.dump("spread:")
	i.nested().Visit(node.Expr)
}

func (i *Inspector) VisitRecordNode(node *RecordNode) {
	i.dump("record:")
	i.nested().dump("target:")
	i.nested().nested().Visit(node.Prototype)
	i.nested().dump("fields (%d):", len(node.Fields))
	for _, kv := range node.Fields {
		i.nested().nested().dump("- key:")
		i.nested().nested().nested().nested().Visit(kv.Key)
		i.nested().nested().nested().dump("value:")
		i.nested().nested().nested().nested().Visit(kv.Value)
	}
}

func (i *Inspector) dump(s string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(i.pad+s, args...))
}

func (i *Inspector) nested() *Inspector {
	if i.child == nil {
		i.child = &Inspector{pad: i.pad + "  "}
	}

	return i.child
}

// Return a new AST inspector.
func NewInspector() *Inspector {
	return &Inspector{pad: ""}
}

// Inspect the AST by dumping it to os.Stdout.
func Inspect(node ASTNode) {
	NewInspector().Visit(node)
}
