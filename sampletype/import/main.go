package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
)

const bar = `package main

import "fmt"



const template = "Hello, world.\nHello %s.\n"

var defaultName = "bob"

func greet(name string) {
	if name == "" {
		name = defaultName
	}
	fmt.Printf(template, name)
}


func main() {
	greet("sara")
	greet("")
}`

func main() {
	fset := token.NewFileSet()

	f, _ := parser.ParseFile(fset, "foo.go", bar, 0)

	conf := types.Config{Importer: importer.Default()}

	pkg, _ := conf.Check("bbbb/aaaaa", fset, []*ast.File{f}, nil)

	fmt.Printf("Package  %q\n", pkg.Path())
	fmt.Printf("Name:    %s\n", pkg.Name())
	fmt.Printf("Imports: %s\n", pkg.Imports())
	fmt.Printf("Scope:   %s\n", pkg.Scope())
	fmt.Printf("Complete:   %t\n", pkg.Complete())
	fmt.Printf("String:   %s\n", pkg.String())
	fmt.Println()

	scope := pkg.Scope()
	fmt.Printf("scope's names = %v\n", scope.Names())
	fmt.Printf("lookup `main` = %v\n", scope.Lookup("main"))
	fmt.Printf("lookup `defaultName` = %v\n", scope.Lookup("defaultName"))
	fmt.Printf("lookup `greet` = %v\n", scope.Lookup("greet"))
	fmt.Printf("lookup `template` = %v\n", scope.Lookup("template"))
	fmt.Printf("lookup `mainx` = %v\n", scope.Lookup("mainx"))

}
