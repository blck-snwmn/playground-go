package module

import "fmt"

func ExportedFunc() string {
	return "exported"
}

func ExportedFunc2() string {
	return "exported2"
}

func ExportedFunc3() string {
	return "exported3"
}

func internalFunc() string {
	return "internal"
}

func internalFunc2() string {
	return "internal2"
}

func ExportedFunc4() string {
	if true {
		return internalFunc()
	}
	// internalFunc2() is not deadcode
	return internalFunc2()
}

type Executor interface {
	Execute() string
}

type PrintExecutor struct{}

func (p *PrintExecutor) Execute() string {
	fmt.Println("print")
	return "print"
}

func (p *PrintExecutor) Execute2() string {
	// Execute2() is not deadcode
	// but not called from main.go
	return "print2"
}

type OnlyReturnExecutor struct{}

func (p *OnlyReturnExecutor) Execute() string {
	return "only return"
}

type SampleStruct struct{}

func (s SampleStruct) ExportedMethod() string {
	return "exported method"
}

func (s SampleStruct) ExportedMethod2() string {
	return "exported method2"
}

type PSampleStruct struct{}

func (s *PSampleStruct) ExportedMethodP() string {
	return "exported method"
}

func (s *PSampleStruct) ExportedMethodP2() string {
	return "exported method2"
}
