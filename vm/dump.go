package vm

import (
	"fmt"
)

// Human-readable value types
var typeNames = map[uint8]string{
	V_NIL: "NIL",
	V_INT: "INT",
	V_STR: "STR",
}

// Dump loaded Bytecode in human-readable form.
// Useful for debugging.
func Dump(vm *VM) {
	var (
		ax uint32
		bx uint16
		cx uint8
	)

	fmt.Println("; Begin Code Dump")

	fmt.Println("; Constants:")
	for _, v := range vm.Constants {
		fmt.Println(fmt.Sprintf(".%s", typeNames[v.Type()]), v.String())
	}

	fmt.Println("; Code")
	for _, inst := range vm.Instructions {
		switch (inst >> 26) & 0x3F {
		case OP_RETURN:
			fmt.Println("RETURN")
		case OP_LOADK:
			decodeAxBx(inst, &ax, &bx)
			fmt.Println("LOADK", ax, bx)
		case OP_ADD:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			fmt.Println("ADD", ax, bx, cx)
		case OP_SUB:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			fmt.Println("SUB", ax, bx, cx)
		case OP_MUL:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			fmt.Println("MUL", ax, bx, cx)
		case OP_DIV:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			fmt.Println("DIV", ax, bx, cx)
		case OP_EQ:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			fmt.Println("EQ", ax, bx, cx)
		case OP_LT:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			fmt.Println("LT", ax, bx, cx)
		case OP_LTE:
			decodeAxBxCx(inst, &ax, &bx, &cx)
			fmt.Println("LTE", ax, bx, cx)
		case OP_PRINT:
			decodeAx(inst, &ax)
			fmt.Println("PRINT", ax)
		}
	}

	fmt.Println("; End Code Dump")
}
