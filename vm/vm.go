package vm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

// Endianness of the virtual machine.
var ByteOrder = binary.LittleEndian

// An instance of the virtual machine.
type VM struct {
	// Number of bytes loaded
	BinSize int
	// Raw bytecode to process
	Code *bytes.Buffer
	// Constants defined in the program
	Constants []Value
	// Virtual machine registers
	Registers []Value
	// Executable code instructions
	Instructions []uint32
	// Current instruction pointer
	IP int64
}

//Create a new instance of the virtual machine.
func NewVM(code []byte) *VM {
	vm := &VM{
		Code:         bytes.NewBuffer(code),
		BinSize:      len(code),
		Constants:    make([]Value, 0, 256),
		Registers:    make([]Value, 256),
		Instructions: make([]uint32, 0, 1024),
		IP:           0,
	}
	vm.load()
	return vm
}

// Execute the loaded bytecode.
func (vm *VM) Run() {
	vm.loop()
}

func (vm *VM) load() {
	vm.loadConstants()
	vm.loadInstructions()
}

func (vm *VM) loadConstants() {
	var (
		numConsts uint32
		valueType uint8
		intValue  int64
		strLength uint32
		boolValue uint8
	)

	binary.Read(vm.Code, ByteOrder, &numConsts)
	for numConsts > 0 {
		binary.Read(vm.Code, ByteOrder, &valueType)
		switch valueType {
		case V_NIL:
			vm.Constants = append(vm.Constants, Nil)
		case V_INT:
			binary.Read(vm.Code, ByteOrder, &intValue)
			vm.Constants = append(vm.Constants, Integer(intValue))
		case V_BLN:
			binary.Read(vm.Code, ByteOrder, &boolValue)
			vm.Constants = append(vm.Constants, Boolean(boolValue != 0))
		case V_STR:
			binary.Read(vm.Code, ByteOrder, &strLength)
			b := make([]byte, strLength)
			vm.Code.Read(b)
			vm.Constants = append(vm.Constants, String(b))
		}
		numConsts--
	}
}

func (vm *VM) loadInstructions() {
	var (
		numInsts uint64
		inst     uint32
	)

	binary.Read(vm.Code, ByteOrder, &numInsts)
	for numInsts > 0 {
		binary.Read(vm.Code, ByteOrder, &inst)
		vm.Instructions = append(vm.Instructions, inst)
		numInsts--
	}
}

func (vm *VM) handleError(reason uint32, bx uint16, cx uint8) {
	switch reason {
	case E_BADMATCH:
		fmt.Fprintln(os.Stderr, &BadMatch{vm.Registers[bx], vm.Registers[cx]})
	}
}

func (vm *VM) loop() {
	var (
		inst uint32
		ax   uint32
		bx   uint16
		cx   uint8
	)

	for {
		inst = vm.Instructions[vm.IP]

		switch (inst >> 26) & 0x3F {
		case OP_RETURN:
			return
		case OP_ERR:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			vm.handleError(ax, bx, cx)
			return
		case OP_JMP:
			decodeAx(inst, &ax)
			vm.IP += int64(ax)
		case OP_JMPIF:
			decodeAxBx(inst, &ax, &bx)
			switch vm.Registers[ax] {
			case Nil, Boolean(false):
				vm.IP += 0 // noop
			default:
				vm.IP += int64(bx)
			}
		case OP_MOVE:
			decodeAxBx(inst, &ax, &bx)
			vm.Registers[ax] = vm.Registers[bx]
		case OP_LOADK:
			decodeAxBx(inst, &ax, &bx)
			vm.Registers[ax] = vm.Constants[bx]
		case OP_ADD:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			vm.Registers[ax] = vm.Registers[bx].Add(vm.Registers[cx])
		case OP_SUB:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			vm.Registers[ax] = vm.Registers[bx].Sub(vm.Registers[cx])
		case OP_MUL:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			vm.Registers[ax] = vm.Registers[bx].Mul(vm.Registers[cx])
		case OP_DIV:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			vm.Registers[ax] = vm.Registers[bx].Div(vm.Registers[cx])
		case OP_EQ:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			vm.Registers[ax] = Boolean(vm.Registers[bx].Eq(vm.Registers[cx]))
		case OP_LT:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			vm.Registers[ax] = Boolean(vm.Registers[bx].Lt(vm.Registers[cx]))
		case OP_LTE:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			vm.Registers[ax] = Boolean(vm.Registers[bx].Lte(vm.Registers[cx]))
		case OP_PRINT:
			decodeAx(inst, &ax)
			fmt.Println(vm.Registers[ax])
		default:
			panic(fmt.Sprintf("Unhandled opcode: 0x%x", (inst>>26)&0x3F))
		}
		vm.IP++
	}

	return
}
