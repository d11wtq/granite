package vm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sort"
)

// Convert boolean v to a uint8.
func btoi(v bool) uint8 {
	if v {
		return 1
	}

	return 0
}

// Operand of the three address instruction.
// Resolution of the actual register value is defered to code gen time.
type Operand interface {
	// Resolve this operand to a specific register value.
	// The arguments are the assember and the IP of this instruction.
	Resolve(*ASM, uint64) uint32
}

// 32-bit instruction.
// Resolution of the actual register values is defered to code gen time.
type Instruction interface {
	// Resolve the operands in this instruction to compute its final value.
	// The arguments are the assember and the IP of this instruction.
	Resolve(*ASM, uint64) uint32
	// Return the operands that are Vars
	ActiveVars() []Var
}

// A fixed operand with a value known ahead of time.
type Reg uint32

func (o Reg) Resolve(*ASM, uint64) uint32 {
	return uint32(o)
}

// A jmp offset to a label, not known until code gen time.
type Jmp string

func (o Jmp) Resolve(a *ASM, i uint64) uint32 {
	j, ok := a.Labels[string(o)]
	if ok == false {
		panic(fmt.Sprintf("invalid JMP label `%s'", string(o)))
	}
	return uint32(j - i) // could wrap, and that's fine
}

// Reference to a variable, resolved to a register at code gen time.
type Var string

func (o Var) Resolve(a *ASM, i uint64) uint32 {
	return a.Locals[o]
}

// The location of a constant in the constants pool.
type Constant struct {
	v Value
}

func (o *Constant) Resolve(a *ASM, i uint64) uint32 {
	return a.DefConst(o.v)
}

// A three-operand instruction.
type AxBxCx struct {
	op      uint32
	a, b, c Operand
}

func (o *AxBxCx) Resolve(a *ASM, i uint64) uint32 {
	return encodeAxBxCx(
		o.op,
		o.a.Resolve(a, i),
		o.b.Resolve(a, i),
		o.c.Resolve(a, i),
	)
}

func (o *AxBxCx) ActiveVars() []Var {
	a, aok := o.a.(Var)
	b, bok := o.b.(Var)
	c, cok := o.c.(Var)

	if aok || bok || cok {
		v := make([]Var, 0, 3)
		if aok {
			v = append(v, a)
		}
		if bok {
			v = append(v, b)
		}
		if cok {
			v = append(v, c)
		}
		return v
	}

	return []Var(nil)
}

// A two-operand instruction.
type AxBx struct {
	op   uint32
	a, b Operand
}

func (o *AxBx) Resolve(a *ASM, i uint64) uint32 {
	return encodeAxBx(
		o.op,
		o.a.Resolve(a, i),
		o.b.Resolve(a, i),
	)
}

func (o *AxBx) ActiveVars() []Var {
	a, aok := o.a.(Var)
	b, bok := o.b.(Var)

	if aok || bok {
		v := make([]Var, 0, 2)
		if aok {
			v = append(v, a)
		}
		if bok {
			v = append(v, b)
		}
		return v
	}

	return []Var(nil)
}

// A single operand instruction.
type Ax struct {
	op uint32
	a  Operand
}

func (o *Ax) Resolve(a *ASM, i uint64) uint32 {
	return encodeAx(
		o.op,
		o.a.Resolve(a, i),
	)
}

func (o *Ax) ActiveVars() []Var {
	a, aok := o.a.(Var)

	if aok {
		return []Var{a}
	}

	return []Var(nil)
}

// An assembler for producing register machine byte code.
// The interface closely maps to what the textual ASM form would look like.
type ASM struct {
	Constants    []Value
	Instructions []Instruction
	Labels       map[string]uint64
	Locals       map[Var]uint32
	constMap     map[Value]uint32
	labelNum     uint32
}

// Create a new assembler.
func NewASM() *ASM {
	return &ASM{
		make([]Value, 0, 256),
		make([]Instruction, 0, 1024),
		make(map[string]uint64),
		make(map[Var]uint32),
		make(map[Value]uint32),
		0,
	}
}

// Generate a unique name for a label within this ASM.
func (asm *ASM) GenLabel() string {
	asm.labelNum++
	return fmt.Sprintf(".label%d", asm.labelNum)
}

// Generate byte-code for the assembled program.
// All instructions defined are resolved and converted to a []byte.
func (asm *ASM) GetCode() []byte {
	asm.allocateLocals()

	b := new(bytes.Buffer)

	// Write instructions
	binary.Write(b, ByteOrder, uint64(len(asm.Instructions)))

	for x, i := range asm.Instructions {
		binary.Write(b, ByteOrder, i.Resolve(asm, uint64(x)))
	}

	// Write constants
	binary.Write(b, ByteOrder, uint32(len(asm.Constants)))
	for _, v := range asm.Constants {
		binary.Write(b, ByteOrder, v.Type())
		switch t := v.(type) {
		case *NilType, *Vector:
			// noop
		case Integer:
			binary.Write(b, ByteOrder, v)
		case Boolean:
			binary.Write(b, ByteOrder, btoi(bool(t)))
		case String:
			s := []byte(string(t))
			binary.Write(b, ByteOrder, uint32(len(s)))
			b.Write(s)
		}
	}

	return b.Bytes()
}

// Define a constant value in the constant pool and return its index.
// Constants are never duplicated.
func (asm *ASM) DefConst(v Value) uint32 {
	if _, ok := asm.constMap[v]; ok == false {
		asm.constMap[v] = uint32(len(asm.Constants))
		asm.Constants = append(asm.Constants, v)
	}

	return asm.constMap[v]
}

// Define a label at the current position in the program.
// Jmp's can refer to this label to cleanly compute offsets at code gen time.
func (asm *ASM) SetLabel(label string) {
	asm.Labels[label] = uint64(len(asm.Instructions))
}

// Return a given register value from the current procedure.
// This causes the current stack frame to exit.
func (asm *ASM) Return(reg Operand) {
	asm.addInstruction(&Ax{OP_RETURN, reg})
}

// Assert two registers are equal, halt with a BadMatch if not
func (asm *ASM) Assert(a, b Operand) {
	asm.addInstruction(&AxBx{OP_ASSERT, a, b})
}

// Adjust the instruction pointer by offset.
func (asm *ASM) Jmp(offset Operand) {
	asm.addInstruction(&Ax{OP_JMP, offset})
}

// Adjust the instruction pointer by offset if the value in test is truthy,
func (asm *ASM) JmpIf(test, offset Operand) {
	asm.addInstruction(&AxBx{OP_JMPIF, test, offset})
}

// Copy the value in register a to register b.
func (asm *ASM) Move(a, b Operand) {
	asm.addInstruction(&AxBx{OP_MOVE, a, b})
}

// Load the constant idx from the constant pool to register dst.
func (asm *ASM) LoadK(dst, idx Operand) {
	asm.addInstruction(&AxBx{OP_LOADK, dst, idx})
}

// Test if the value in register test is of type t, placing the result in dst.
func (asm *ASM) IsA(dst, test, t Operand) {
	asm.addInstruction(&AxBxCx{OP_ISA, dst, test, t})
}

// Add the values in registers a and b, putting the result in dst.
func (asm *ASM) Add(dst, a, b Operand) {
	asm.addInstruction(&AxBxCx{OP_ADD, dst, a, b})
}

// Subtract the values in registers a and b, putting the result in dst.
func (asm *ASM) Sub(dst, a, b Operand) {
	asm.addInstruction(&AxBxCx{OP_SUB, dst, a, b})
}

// Multiply the values in registers a and b, putting the result in dst.
func (asm *ASM) Mul(dst, a, b Operand) {
	asm.addInstruction(&AxBxCx{OP_MUL, dst, a, b})
}

// Divide the values in registers a and b, putting the result in dst.
func (asm *ASM) Div(dst, a, b Operand) {
	asm.addInstruction(&AxBxCx{OP_DIV, dst, a, b})
}

// Compare the values in registers a and b, putting the result in dst.
func (asm *ASM) Eq(dst, a, b Operand) {
	asm.addInstruction(&AxBxCx{OP_EQ, dst, a, b})
}

// Compare the values in registers a and b, putting the result in dst.
func (asm *ASM) Lt(dst, a, b Operand) {
	asm.addInstruction(&AxBxCx{OP_LT, dst, a, b})
}

// Compare the values in registers a and b, putting the result in dst.
func (asm *ASM) Lte(dst, a, b Operand) {
	asm.addInstruction(&AxBxCx{OP_LTE, dst, a, b})
}

// Get the length of the value in register test and place the result in dst.
func (asm *ASM) Len(dst, test Operand) {
	asm.addInstruction(&AxBx{OP_LEN, dst, test})
}

// Append the value in register b to the vector in register a, putting the result in dst.
func (asm *ASM) Append(dst, a, b Operand) {
	asm.addInstruction(&AxBxCx{OP_APPEND, dst, a, b})
}

// Get the element looked up by the value in register k from the map or vector in register a.
func (asm *ASM) Get(dst, a, k Operand) {
	asm.addInstruction(&AxBxCx{OP_GET, dst, a, k})
}

// Create a closure of size length in register dst.
func (asm *ASM) Fn(dst, length Operand) {
	asm.addInstruction(&AxBx{OP_FN, dst, length})
}

// Call a function and return its result.
func (asm *ASM) Call(dst, target Operand) {
	asm.addInstruction(&AxBx{OP_CALL, dst, target})
}

// Dump the value in register a to stdout (debug use).
func (asm *ASM) Print(a Operand) {
	asm.addInstruction(&Ax{OP_PRINT, a})
}

func (asm *ASM) addInstruction(i Instruction) {
	asm.Instructions = append(asm.Instructions, i)
}

// Live Interval for a variable
type interval struct {
	beg, end int
	register uint32
}

// Sorting algorithm representation
type byEnd []*interval

func (a byEnd) Len() int {
	return len(a)
}

func (a byEnd) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byEnd) Less(i, j int) bool {
	return a[i].end < a[j].end
}

func (asm *ASM) allocateLocals() {
	// http://web.cs.ucla.edu/~palsberg/course/cs132/linearscan.pdf

	var (
		numRegisters  = uint32(256)
		liveIntervals = make([]*interval, 0, numRegisters)
		active        = make([]*interval, 0, numRegisters)
		seenVars      = make(map[Var]int)
		pool          = NewRegisterPool(numRegisters)
	)

	// assign live intervals
	for k, i := range asm.Instructions {
		for _, x := range i.ActiveVars() {
			idx, exists := seenVars[x]

			if exists {
				liveIntervals[idx].end = k
			} else {
				seenVars[x] = len(liveIntervals)
				liveIntervals = append(liveIntervals, &interval{k, k, 0})
			}
		}
	}

	// allocate registers
	for _, i := range liveIntervals {
		for _, j := range active {
			if j.end >= i.beg {
				break
			}

			active = active[1:]
			pool.Release(j.register)
		}

		i.register = pool.Reserve()
		active = append(active, i)
		sort.Sort(byEnd(active))
	}

	// associate with local vars
	for x, k := range seenVars {
		asm.Locals[x] = liveIntervals[k].register
	}
}
