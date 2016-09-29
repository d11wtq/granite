package main

import (
	"./parser"
	"./parser/ast"
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
		fmt.Println(err)
		return
	}

	err, result := parser.Parse(
		bufio.NewReader(file),
		os.Args[1],
	)

	if err != nil {
		fmt.Println("Parse error", err)
		return
	}

	ast.Dump(result)
}
