package vm

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	OP_RETURN = iota
	OP_LOADK
	OP_ADD
	OP_SUB
	OP_MUL
	OP_DIV
	OP_PRINT
	OP_SLEEP
)

// Endianness of the virtual machine.
var ByteOrder = binary.LittleEndian

// An instance of the virtual machine.
type VM struct {
	// Raw bytecode to process
	Code *bytes.Buffer
	// Constants defined in the program
	Constants []Value
	// Virtual machine registers
	Registers []Value
	// Executable code instructions
	Instructions []uint32
	// Current instruction pointer
	IP uint64
}

//Create a new instance of the virtual machine.
func NewVM(code []byte) *VM {
	return &VM{
		Code:         bytes.NewBuffer(code),
		Constants:    make([]Value, 0, 256),
		Registers:    make([]Value, 256),
		Instructions: make([]uint32, 0, 1024),
		IP:           0,
	}
}

// Execute the loaded bytecode.
func (vm *VM) Run() {
	vm.loadConstants()
	vm.loadInstructions()
	vm.loop()
}

func (vm *VM) loadConstants() {
	var (
		numConsts uint32
		valueType uint8
		intValue  int64
		strLength uint32
	)

	binary.Read(vm.Code, ByteOrder, &numConsts)
	for numConsts > 0 {
		binary.Read(vm.Code, ByteOrder, &valueType)
		switch valueType {
		case V_INT:
			binary.Read(vm.Code, ByteOrder, &intValue)
			vm.Constants = append(vm.Constants, Integer(intValue))
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

func (vm *VM) loop() {
	var (
		op uint8
		ax uint32
		bx uint16
		cx uint8
	)

	for {
		op = uint8(vm.Instructions[vm.IP]>>26) & 0x3F
		switch op {
		case OP_LOADK:
			decodeAxBx(vm.Instructions[vm.IP], &ax, &bx)
			vm.Registers[ax] = vm.Constants[bx]
			vm.IP++
		case OP_ADD:
			decodeAxBxCx(vm.Instructions[vm.IP], &ax, &bx, &cx)
			vm.Registers[ax] = vm.Registers[bx].Add(vm.Registers[cx])
			vm.IP++
		case OP_SUB:
			decodeAxBxCx(vm.Instructions[vm.IP], &ax, &bx, &cx)
			vm.Registers[ax] = vm.Registers[bx].Sub(vm.Registers[cx])
			vm.IP++
		case OP_MUL:
			decodeAxBxCx(vm.Instructions[vm.IP], &ax, &bx, &cx)
			vm.Registers[ax] = vm.Registers[bx].Mul(vm.Registers[cx])
			vm.IP++
		case OP_DIV:
			decodeAxBxCx(vm.Instructions[vm.IP], &ax, &bx, &cx)
			vm.Registers[ax] = vm.Registers[bx].Div(vm.Registers[cx])
			vm.IP++
		case OP_PRINT:
			decodeAx(vm.Instructions[vm.IP], &ax)
			fmt.Println(vm.Registers[ax])
			vm.IP++
		case OP_RETURN:
			return
		default:
			panic(fmt.Sprintf("Unhandled opcode: 0x%x", op))
		}
	}

	return
}
