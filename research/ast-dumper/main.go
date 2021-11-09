package main

import (
	"fmt"
	"os"

	"github.com/wu-xian/go/research/ast-dumper/syntax"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:")
		fmt.Println("\tastdump [filename]")
		os.Exit(1)
	}
	fname := os.Args[1]
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println("cannot open file: %v", err)
		os.Exit(1)
	}
	defer file.Close()
	errh := syntax.ErrorHandler(
		func(err error) {
			fmt.Println(err)
		})

	mode := syntax.AllowGenerics

	ast, _ := syntax.ParseFile(
		fname,
		errh,
		nil,
		mode)

	f, _ := os.Create(fname + ".ast")
	defer f.Close()
	syntax.Fdump(f, ast) //<--this prints out your AST nicely

}
