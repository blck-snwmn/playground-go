package main

import (
	"fmt"

	"github.com/blck-snwmn/playground-go/deadcodesample/module"
)

func main() {
	str := module.ExportedFunc()
	fmt.Println(str)

	_ = []module.Executor{
		&module.PrintExecutor{}, // PrintExecutor.Execute() is not deadcode
		// &module.OnlyReturnExecutor{}, // OnlyReturnExecutor.Execute() is deadcode
	}

	module.ExportedFunc4()

	s := module.SampleStruct{}
	s.ExportedMethod()

	ps := &module.PSampleStruct{}
	ps.ExportedMethodP()
}
