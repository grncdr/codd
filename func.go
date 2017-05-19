package codd

type FuncCall struct {
	name string
	args []Node
}

func Func(name string, arity int) func(...Node) FuncCall {
	return func(nodes ...Node) FuncCall {
		return FuncCall{name, nodes}
	}
}

var (
	Count = Func("COUNT", -1)
)
