package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	// only
	// preix=`// Code generated `
	// sufix=` DO NOT EDIT.`
	{
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, "generated.go", nil, parser.AllErrors|parser.ParseComments)
		if err != nil {
			panic(err)
		}
		fmt.Println(ast.IsGenerated(node))
	}
	{
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, "not_generated_1.go", nil, parser.AllErrors|parser.ParseComments)
		if err != nil {
			panic(err)
		}
		fmt.Println(ast.IsGenerated(node))
	}
	{
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, "not_generated_2.go", nil, parser.AllErrors|parser.ParseComments)
		if err != nil {
			panic(err)
		}
		fmt.Println(ast.IsGenerated(node))
	}
}
