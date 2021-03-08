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

type (
	xxxx struct {
		int
		string
	}
	yyyy struct {
		xi int
		xs string
	}
)

func (xxxx)Test(){}

const template = "Hello, world.\nHello %s.\n"

var defaultName = "bob"

func greet(name string) {
	if name == "" {
		name = defaultName
	}
	fmt.Printf(template, name)
}

func main() {
	var name1 = "sara"
	const name2 = "tom"
	greet(name1)
	innerFunc := func(name string){
		greet(name)
	}
	innerFunc(name2)
	innerFunc("")
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
	names := scope.Names()
	fmt.Printf("scope's names = %v\n", names)
	for _, n := range names {
		obj := scope.Lookup(n)
		fmt.Printf("lookup `%s` = %v[%v]\n", n, obj, obj.Type())
	}
	fmt.Printf("lookup `mainx` = %v\n", scope.Lookup("mainx"))
	fmt.Println()

	mainFunc, _ := scope.Lookup("main").(*types.Func)
	fmt.Printf("mainFunc: String() = %v\n", mainFunc.String())
	fmt.Printf("mainFunc: FullName() = %v\n", mainFunc.FullName())
	fmt.Printf("mainFunc: Type() = %v\n", mainFunc.Type())
	fmt.Printf("mainFunc: Scope() = %v\n", mainFunc.Scope())
	fmt.Println()

	greetFunc, _ := scope.Lookup("greet").(*types.Func)
	fmt.Printf("greetFunc: Type() = %v\n", greetFunc.Type())
	fmt.Println()

	templateConst, _ := scope.Lookup("template").(*types.Const)
	fmt.Printf("templateConst: String() = %v\n", templateConst.String())
	fmt.Printf("templateConst: Type() = %v\n", templateConst.Type())
	fmt.Printf("templateConst: Val() = %v\n", templateConst.Val())
	fmt.Println()

	xxxxTypeName, _ := scope.Lookup("xxxx").(*types.TypeName)
	fmt.Printf("xxxxStruct: Type() = %v\n", xxxxTypeName.Type())
	fmt.Printf("xxxxStruct: Parent() = %v\n", xxxxTypeName.Parent())
	fmt.Printf("xxxxStruct: Exported() = %v\n", xxxxTypeName.Exported())
	fmt.Println()

	xxxxNames, _ := xxxxTypeName.Type().(*types.Named)
	fmt.Printf("xxxxNames: NumMethods() = %v\n", xxxxNames.NumMethods())
	fmt.Printf("xxxxNames: Underlying() = %v\n", xxxxNames.Underlying())
	fmt.Printf("xxxxNames: Obj() = %v\n", xxxxNames.Obj())
	fmt.Println()

	s, _ := xxxxNames.Underlying().(*types.Struct)
	fmt.Printf("xxxxNames: NumFields() = %v\n", s.NumFields())
	fmt.Printf("xxxxNames: Field(0) = %v (Anonymous? %t)\n", s.Field(0), s.Field(0).Anonymous())
	fmt.Printf("xxxxNames: Field(1) = %v (Anonymous? %t)\n", s.Field(1), s.Field(1).Anonymous())
	fmt.Println()

	yyyyTypeName, _ := scope.Lookup("yyyy").(*types.TypeName)
	yyyyNames, _ := yyyyTypeName.Type().(*types.Named)
	syyyy, _ := yyyyNames.Underlying().(*types.Struct)
	fmt.Printf("xxxxNames: NumFields() = %v\n", syyyy.NumFields())
	fmt.Printf("xxxxNames: Field(0) = %v (Anonymous? %t)\n", syyyy.Field(0), syyyy.Field(0).Anonymous())
	fmt.Printf("xxxxNames: Field(1) = %v (Anonymous? %t)\n", syyyy.Field(1), syyyy.Field(1).Anonymous())
	fmt.Println()
}
