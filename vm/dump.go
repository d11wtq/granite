package vm

import (
	"fmt"
)

// Dump loaded Bytecode in human-readable form.
// Useful for debugging.
func Dump(vm *VM) {
	var (
		ax uint32
		bx uint32
		cx uint8
	)

	fmt.Println(
		fmt.Sprintf(
			"; <main :: size %dB, %d constants, %d instructions>",
			vm.BinSize,
			len(vm.Constants),
			len(vm.Instructions),
		),
	)
	fmt.Println("; constants")
	for _, v := range vm.Constants {
		fmt.Println("  .const", v.String())
	}

	fmt.Println("; instructions")
	for _, inst := range vm.Instructions {
		switch (inst >> 26) & 0x3F {
		case OP_RETURN:
			decodeAx(inst, &ax)
			fmt.Println("  RETURN", ax)
		case OP_ASSERT:
			decodeAxBx(inst, &ax, &bx)
			fmt.Println("  ASSERT", ax, bx)
		case OP_JMP:
			decodeAx(inst, &ax)
			fmt.Println("  JMP", ax)
		case OP_JMPIF:
			decodeAxBx(inst, &ax, &bx)
			fmt.Println("  JMPIF", ax, bx)
		case OP_MOVE:
			decodeAxBx(inst, &ax, &bx)
			fmt.Println("  MOVE", ax, bx)
		case OP_LOADK:
			decodeAxBx(inst, &ax, &bx)
			fmt.Println("  LOADK", ax, bx)
		case OP_ISA:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			fmt.Println("  ISA", ax, bx, cx)
		case OP_ADD:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			fmt.Println("  ADD", ax, bx, cx)
		case OP_SUB:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			fmt.Println("  SUB", ax, bx, cx)
		case OP_MUL:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			fmt.Println("  MUL", ax, bx, cx)
		case OP_DIV:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			fmt.Println("  DIV", ax, bx, cx)
		case OP_EQ:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			fmt.Println("  EQ", ax, bx, cx)
		case OP_LT:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			fmt.Println("  LT", ax, bx, cx)
		case OP_LTE:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			fmt.Println("  LTE", ax, bx, cx)
		case OP_LEN:
			decodeAxBx(inst, &ax, &bx)
			fmt.Println("  LEN", ax, bx)
		case OP_APPEND:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			fmt.Println("  APPEND", ax, bx, cx)
		case OP_GET:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			fmt.Println("  GET", ax, bx, cx)
		case OP_FN:
			decodeAxBx(inst, &ax, &bx)
			fmt.Println("  FN", ax, bx)
		case OP_CALL:
			decodeAxBx(inst, &ax, &bx)
			fmt.Println("  CALL", ax, bx)
		case OP_PRINT:
			decodeAx(inst, &ax)
			fmt.Println("  PRINT", ax)
		}
	}
}
