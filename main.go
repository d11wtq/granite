package main

import (
	"./parser"
	"./parser/ast"
	"./vm"
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Filename required")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	err, result := parser.Parse(
		bufio.NewReader(file),
		os.Args[1],
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if os.Getenv("DUMPAST") != "" {
		ast.Dump(result)
	}

	code, err := vm.Compile(result)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	vmx := vm.NewVM(code)

	if os.Getenv("DUMPASM") != "" {
		vm.Dump(vmx)
	}

	vmx.Run()
}
