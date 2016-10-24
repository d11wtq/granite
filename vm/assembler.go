package vm

import (
	"bytes"
	"encoding/binary"
)

// Convert boolean v to a uint8.
func btoi(v bool) uint8 {
	if v {
		return 1
	}

	return 0
}

// Bytecode assember for Bijou code
type Assembler struct {
	// Program constant pool
	Constants []Value
	// List of executable instructions
	Instructions []uint32
	// Locations of constants in the constant pool
	ConstantMap map[Value]uint32
}

// Create a new bytecode assembler.
func NewAssembler() *Assembler {
	return &Assembler{
		Constants:    make([]Value, 0),
		Instructions: make([]uint32, 0),
		ConstantMap:  make(map[Value]uint32),
	}
}

// Produce bytecode from the built ASM.
func (a *Assembler) GetCode() []byte {
	b := new(bytes.Buffer)
	a.encodeConstants(b)
	a.encodeInstructions(b)
	return b.Bytes()
}

// Add a constant for use with LOADK.
// The index of the constant in the pool is returned.
func (a *Assembler) AddConstant(v Value) uint32 {
	if _, ok := a.ConstantMap[v]; ok == false {
		a.ConstantMap[v] = uint32(len(a.Constants))
		a.Constants = append(a.Constants, v)
	}

	return a.ConstantMap[v]
}

// Append an instruction to the list of instructions.
func (a *Assembler) AddInstruction(inst uint32) {
	a.Instructions = append(a.Instructions, inst)
}

// Append a placeholder instruction and return its location.
func (a *Assembler) ReserveInstruction(op uint32) int {
	idx := len(a.Instructions)
	a.AddInstruction(encodeAx(op, 0))
	return idx
}

// Replace an existing instruction with another.
func (a *Assembler) SetInstruction(idx int, inst uint32) {
	a.Instructions[idx] = inst
}

func (a *Assembler) encodeConstants(b *bytes.Buffer) {
	binary.Write(b, ByteOrder, uint32(len(a.Constants)))
	for _, v := range a.Constants {
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
}

func (a *Assembler) encodeInstructions(b *bytes.Buffer) {
	binary.Write(b, ByteOrder, uint64(len(a.Instructions)+2))

	for _, inst := range a.Instructions {
		binary.Write(b, ByteOrder, inst)
	}

	binary.Write(b, ByteOrder, encodeAx(OP_PRINT, 0))
	binary.Write(b, ByteOrder, encodeAx(OP_RETURN, 0))
}
