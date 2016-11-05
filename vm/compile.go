package vm

import (
	"../parser/ast"
	"fmt"
)

// Compile the given AST into bytecode.
func Compile(prog ast.ASTNode) ([]byte, error) {
	return NewCompiler().Visit(prog).GetCode()
}

// Bytecode compiler for Bijou code
type Compiler struct {
	// Bytecode assembler
	Assembler *ASM
	// Location of the result
	Result Operand
	// Map of local vars to registers
	Locals map[string]Operand
	// Compile error
	Error error
}

// Create a new bytecode compiler.
func NewCompiler() *Compiler {
	return &Compiler{
		Assembler: NewASM(),
		Locals:    make(map[string]Operand),
	}
}

func (c *Compiler) Visit(node ast.ASTNode) *Compiler {
	node.Accept(c)
	return c
}

// Produce bytecode for the processed AST
func (c *Compiler) GetCode() ([]byte, error) {
	if c.Error != nil {
		return nil, c.Error
	}

	c.Assembler.Print(c.Result)
	c.Assembler.Return()

	return c.Assembler.GetCode(), nil
}

func (c *Compiler) tempVar() Operand {
	return Var(c.Assembler.GenLabel())
}

func (c *Compiler) loadConstant(v Value) {
	c.Result = c.tempVar()
	c.Assembler.LoadK(c.Result, &Constant{v})
}

func (c *Compiler) returnError(err error) {
	c.Error = err
	c.Assembler.Return()
}

func (c *Compiler) visitAssign(node *ast.BinaryExpressionNode) {
	c.setLocal(
		node.Left.(*ast.IdentifierNode).Name,
		c.Visit(node.Right).Result,
	)
}

func (c *Compiler) setLocal(name string, reg Operand) {
	c.Locals[name] = reg
}

func (c *Compiler) getLocal(name string) Operand {
	return c.Locals[name]
}

func (c *Compiler) VisitNil(node *ast.NilNode) {
	c.loadConstant(Nil)
}

func (c *Compiler) VisitInteger(node *ast.IntegerNode) {
	c.loadConstant(Integer(node.Value))
}

func (c *Compiler) VisitBoolean(node *ast.BooleanNode) {
	c.loadConstant(Boolean(node.Value))
}

func (c *Compiler) VisitFloat(node *ast.FloatNode) {
}

func (c *Compiler) VisitString(node *ast.StringNode) {
	c.loadConstant(String(node.String))
}

func (c *Compiler) VisitSymbol(node *ast.SymbolNode) {
}

func (c *Compiler) VisitIdentifier(node *ast.IdentifierNode) {
	c.Result = c.getLocal(node.Name)
}

func (c *Compiler) VisitExpressionList(node *ast.ExpressionList) {
	for _, e := range node.Elements {
		if c.Visit(e).Error != nil {
			return
		}
	}
}

func (c *Compiler) VisitBinaryExpression(node *ast.BinaryExpressionNode) {
	if node.Op == ast.OP_ASS {
		c.visitAssign(node)
		return
	}

	var (
		a = c.Visit(node.Left).Result
		b = c.Visit(node.Right).Result
	)

	c.Result = c.tempVar()

	switch node.Op {
	case ast.OP_ADD:
		c.Assembler.Add(c.Result, a, b)
	case ast.OP_MIN:
		c.Assembler.Sub(c.Result, a, b)
	case ast.OP_MUL:
		c.Assembler.Mul(c.Result, a, b)
	case ast.OP_DIV:
		c.Assembler.Div(c.Result, a, b)
	case ast.OP_EQL:
		c.Assembler.Eq(c.Result, a, b)
	case ast.OP_LT:
		c.Assembler.Lt(c.Result, a, b)
	case ast.OP_GT:
		c.Assembler.Lt(c.Result, b, a)
	case ast.OP_LTE:
		c.Assembler.Lte(c.Result, a, b)
	case ast.OP_GTE:
		c.Assembler.Lte(c.Result, b, a)
	default:
		panic(fmt.Sprintf("Unhandled binary operator: 0x%x", node.Op))
	}
}

func (c *Compiler) VisitUnaryExpression(node *ast.UnaryExpressionNode) {
}

func (c *Compiler) VisitVector(node *ast.VectorNode) {
}

func (c *Compiler) VisitMap(node *ast.MapNode) {
}

func (c *Compiler) VisitRecord(node *ast.RecordNode) {
}

func (c *Compiler) VisitPair(node *ast.PairNode) {
}

func (c *Compiler) VisitFunctionApplication(node *ast.FunctionApplicationNode) {
}

func (c *Compiler) VisitKeyAccess(node *ast.KeyAccessNode) {
}

func (c *Compiler) VisitIfThenElse(node *ast.IfThenElseNode) {
}

func (c *Compiler) VisitCaseExpression(node *ast.CaseExpressionNode) {
}

func (c *Compiler) VisitCatchExpression(node *ast.CatchExpressionNode) {
}

func (c *Compiler) VisitThrow(node *ast.ThrowNode) {
}
