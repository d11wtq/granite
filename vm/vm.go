package vm

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Endianness of the virtual machine.
var ByteOrder = binary.LittleEndian

type Frame struct {
	Registers [256]Value
	// Stack of POP addresses to return to
	LinkRegister []int64
	IP           int64
}

// Create a new stack frame.
func NewFrame(ip int64) *Frame {
	return &Frame{IP: ip, LinkRegister: make([]int64, 0, 16)}
}

func (f *Frame) pushLink(offset int64) {
	f.LinkRegister = append(f.LinkRegister, f.IP+offset)
}

func (f *Frame) popLink() {
	f.IP = f.LinkRegister[len(f.LinkRegister)-1]
	f.LinkRegister = f.LinkRegister[:len(f.LinkRegister)-1]
}

// An instance of the virtual machine.
type VM struct {
	// Number of bytes loaded
	BinSize int
	// Constants defined in the program
	Constants []Value
	// Executable code instructions
	Instructions []uint32
}

//Create a new instance of the virtual machine.
func NewVM(code []byte) *VM {
	vm := &VM{
		BinSize:      len(code),
		Constants:    make([]Value, 0, 256),
		Instructions: make([]uint32, 0, 1024),
	}
	vm.load(bytes.NewBuffer(code))
	return vm
}

// Execute the loaded bytecode.
func (vm *VM) Run(f *Frame) (Value, error) {
	var (
		inst    uint32
		ax      uint32
		bx      uint32
		cx      uint8
		closure *Closure
		err     error
	)

Loop:
	for {
		inst = vm.Instructions[f.IP]

		switch inst & 0x3F {
		case OP_RETURN:
			decodeAx(inst, &ax)
			return f.Registers[ax], nil
		case OP_ASSERT:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			if !f.Registers[ax].Eq(f.Registers[bx]) {
				err = &BadMatch{f.Registers[cx]}
			}
		case OP_JMP:
			decodeAx(inst, &ax)
			f.IP += int64(sAx(ax))
			continue Loop
		case OP_JMPIF:
			decodeAxBx(inst, &ax, &bx)
			switch f.Registers[ax] {
			case Nil, Boolean(false):
			default:
				f.IP += int64(sBx(bx))
				continue Loop
			}
		case OP_PUSH:
			decodeAx(inst, &ax)
			f.pushLink(int64(sAx(ax)))
		case OP_POP:
			f.popLink()
			continue Loop
		case OP_MOVE:
			decodeAxBx(inst, &ax, &bx)
			f.Registers[ax] = f.Registers[bx]
		case OP_LOADK:
			decodeAxBx(inst, &ax, &bx)
			f.Registers[ax] = vm.Constants[bx]
		case OP_ISA:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			f.Registers[ax] = Boolean(f.Registers[bx].Type() == cx)
		case OP_TYPE:
			decodeAxBx(inst, &ax, &bx)
			f.Registers[ax] = Integer(f.Registers[bx].Type())
		case OP_ADD:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			f.Registers[ax], err = f.Registers[bx].Add(f.Registers[cx])
		case OP_SUB:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			f.Registers[ax], err = f.Registers[bx].Sub(f.Registers[cx])
		case OP_MUL:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			f.Registers[ax], err = f.Registers[bx].Mul(f.Registers[cx])
		case OP_DIV:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			f.Registers[ax], err = f.Registers[bx].Div(f.Registers[cx])
		case OP_EQ:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			f.Registers[ax] = Boolean(f.Registers[bx].Eq(f.Registers[cx]))
		case OP_LT:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			f.Registers[ax] = Boolean(f.Registers[bx].Lt(f.Registers[cx]))
		case OP_LTE:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			f.Registers[ax] = Boolean(f.Registers[bx].Lte(f.Registers[cx]))
		case OP_LEN:
			decodeAxBx(inst, &ax, &bx)
			f.Registers[ax], err = f.Registers[bx].Len()
		case OP_APPEND:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			f.Registers[ax], err = f.Registers[bx].Append(f.Registers[cx])
		case OP_GET:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			f.Registers[ax], err = f.Registers[bx].Get(f.Registers[cx])
		case OP_FN:
			decodeAxBx(inst, &ax, &bx)
			closure = NewClosure(vm, f.IP+1)
			f.Registers[ax] = closure
			copy(closure.Env[:], f.Registers[:])
			f.IP += int64(bx)
		case OP_CALL:
			decodeAxBx(inst, &ax, &bx)
			f.Registers[ax], err = f.Registers[bx].Call(Nil)
		case OP_PRINT:
			decodeAx(inst, &ax)
			fmt.Println(f.Registers[ax])
		default:
			panic(fmt.Sprintf("Unhandled opcode: 0x%x", inst&0x3F))
		}

		if err != nil {
			return nil, err
		}

		f.IP++
	}

	panic("vm exited prematurely")
}

func (vm *VM) load(b *bytes.Buffer) {
	vm.loadInstructions(b)
	vm.loadConstants(b)
}

func (vm *VM) loadConstants(b *bytes.Buffer) {
	var (
		numConsts uint32
		valueType uint8
		intValue  int64
		strLength uint32
		boolValue uint8
	)

	binary.Read(b, ByteOrder, &numConsts)
	for numConsts > 0 {
		binary.Read(b, ByteOrder, &valueType)
		switch valueType {
		case V_NIL:
			vm.Constants = append(vm.Constants, Nil)
		case V_INT:
			binary.Read(b, ByteOrder, &intValue)
			vm.Constants = append(vm.Constants, Integer(intValue))
		case V_BLN:
			binary.Read(b, ByteOrder, &boolValue)
			vm.Constants = append(vm.Constants, Boolean(boolValue != 0))
		case V_STR:
			binary.Read(b, ByteOrder, &strLength)
			s := make([]byte, strLength)
			b.Read(s)
			vm.Constants = append(vm.Constants, String(s))
		case V_VEC:
			vm.Constants = append(vm.Constants, EmptyVector)
		}
		numConsts--
	}
}

func (vm *VM) loadInstructions(b *bytes.Buffer) {
	var (
		numInsts uint64
		inst     uint32
	)

	binary.Read(b, ByteOrder, &numInsts)
	for numInsts > 0 {
		binary.Read(b, ByteOrder, &inst)
		vm.Instructions = append(vm.Instructions, inst)
		numInsts--
	}
}
