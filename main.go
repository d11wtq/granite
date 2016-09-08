package main

import (
	"./parser"
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

	err, ast := parser.Parse(
		bufio.NewReader(file),
		os.Args[1],
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", ast)
}
