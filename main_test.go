package codd

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	InitPersonTable()
	m.Run()
}

func PrintQuery(query Query) {
	compiler := &BaseCompiler{}
	compiler.Push(query)
	fmt.Println(compiler.String(), compiler.ParamValues())
}
