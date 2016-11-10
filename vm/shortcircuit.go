package vm

import (
	"../parser/ast"
)

// Compiles 'and' and 'or' expressions.
// The right-hand-side is only executed if the left-hand-side passes the predicate.
func compileShortCircuit(c *Compiler, flag bool, left Operand, right ast.ASTNode) {
	if c.Error != nil {
		return
	}

	var (
		r    = c.tempVar()
		then = c.ASM.GenLabel()
		done = c.ASM.GenLabel()
	)

	c.ASM.JmpIf(left, Jmp(then))

	if flag == true {
		c.ASM.Move(r, left)
		c.ASM.Jmp(Jmp(done))
		c.ASM.SetLabel(then)
		c.ASM.Move(r, c.Visit(right).Result)
		c.ASM.SetLabel(done)
	} else {
		c.ASM.Move(r, c.Visit(right).Result)
		c.ASM.Jmp(Jmp(done))
		c.ASM.SetLabel(then)
		c.ASM.Move(r, left)
		c.ASM.SetLabel(done)
	}

	c.Result = r
}
