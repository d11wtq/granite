package vm

import (
	"../parser/ast"
)

// Compile the given AST into bytecode.
func Compile(prog ast.ASTNode) ([]byte, error) {
	return NewCompiler().Visit(prog).GetCode()
}

// Bytecode compiler for Bijou code
type Compiler struct {
	// Bytecode assembler
	ASM *ASM
	// Location of the result
	Result Operand
	// Table of local vars to registers
	Symbols *SymbolTable
	// Compile error
	Error error
}

// Create a new bytecode compiler.
func NewCompiler() *Compiler {
	return &Compiler{
		ASM:     NewASM(),
		Symbols: NewSymbolTable(),
	}
}

// Compile the given AST node and return the compiler object.
// The result register will be left in Result.
// If a compile error occurs, it is found in Error.
func (c *Compiler) Visit(node ast.ASTNode) *Compiler {
	node.Accept(c)
	return c
}

// Produce bytecode for the processed AST
func (c *Compiler) GetCode() ([]byte, error) {
	if c.Error != nil {
		return nil, c.Error
	}

	c.ASM.Return(c.Result)

	return c.ASM.GetCode(), nil
}

func (c *Compiler) tempVar() Operand {
	return Var(c.ASM.GenLabel())
}

func (c *Compiler) loadConstant(v Value) {
	c.Result = c.tempVar()
	c.ASM.LoadK(c.Result, &Constant{v})
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
	var exists bool
	c.Result, exists = c.Symbols.Get(node.Name)
	if !exists {
		c.Error = NameError(node.Name)
	}
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
		if call, ok := node.Left.(*ast.CallNode); ok {
			if name, ok := call.Target.(*ast.IdentifierNode); ok {
				if _, exists := c.Symbols.Get(name.Name); !exists {
					// FIXME: Function cases must be collated.
					// FIXME: Parameters are not handled
					// FIXME: Closed value access
					// FIXME: Hoisting
					var (
						r    = c.tempVar()
						done = c.ASM.GenLabel()
					)

					c.Symbols.Set(name.Name, r)
					c.ASM.Fn(r, Jmp(done))
					c.Symbols.BeginScope()
					c.Visit(node.Right)
					c.ASM.SetLabel(done)
					c.ASM.Return(c.Result)
					c.Symbols.EndScope()

					c.Result = r
					return
				}
			}
		}

		compilePattern(
			c,
			node.Left,
			c.Visit(node.Right).Result,
		)
		return
	}

	if node.Op == ast.OP_AND || node.Op == ast.OP_OR {
		compileShortCircuit(
			c,
			node.Op == ast.OP_AND,
			c.Visit(node.Left).Result,
			node.Right,
		)
		return
	}

	var (
		a = c.Visit(node.Left).Result
		b = c.Visit(node.Right).Result
	)

	if c.Error != nil {
		return
	}

	c.Result = c.tempVar()

	switch node.Op {
	case ast.OP_ADD:
		c.ASM.Add(c.Result, a, b)
	case ast.OP_MIN:
		c.ASM.Sub(c.Result, a, b)
	case ast.OP_MUL:
		c.ASM.Mul(c.Result, a, b)
	case ast.OP_DIV:
		c.ASM.Div(c.Result, a, b)
	case ast.OP_EQL:
		c.ASM.Eq(c.Result, a, b)
	case ast.OP_LT:
		c.ASM.Lt(c.Result, a, b)
	case ast.OP_GT:
		c.ASM.Lt(c.Result, b, a)
	case ast.OP_LTE:
		c.ASM.Lte(c.Result, a, b)
	case ast.OP_GTE:
		c.ASM.Lte(c.Result, b, a)
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

func (c *Compiler) VisitCall(node *ast.CallNode) {
	var (
		f = c.Visit(node.Target).Result
	)

	c.Result = c.tempVar()
	c.ASM.Call(c.Result, f)
}

func (c *Compiler) VisitKeyAccess(node *ast.KeyAccessNode) {
}

func (c *Compiler) VisitIfThenElse(node *ast.IfThenElseNode) {
	c.Symbols.BeginScope()
	defer c.Symbols.EndScope()

	var (
		a    = c.Visit(node.Expression).Result
		r    = c.tempVar()
		then = c.ASM.GenLabel()
		done = c.ASM.GenLabel()
	)

	if c.Error != nil {
		return
	}

	c.ASM.JmpIf(a, Jmp(then))
	c.ASM.Move(r, c.Visit(node.Else).Result)
	c.ASM.Jmp(Jmp(done))
	c.ASM.SetLabel(then)
	c.ASM.Move(r, c.Visit(node.Then).Result)
	c.ASM.SetLabel(done)

	c.Result = r
}

func (c *Compiler) VisitCaseExpression(node *ast.CaseExpressionNode) {
}

func (c *Compiler) VisitCatchExpression(node *ast.CatchExpressionNode) {
}

func (c *Compiler) VisitThrow(node *ast.ThrowNode) {
}
