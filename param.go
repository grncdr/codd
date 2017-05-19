package codd

type Parameter struct {
	value interface{}
}

func Param(v interface{}) Parameter {
	return Parameter{v}
}

func (p Parameter) Kind() string {
	return "Parameter"
}

func (p Parameter) SQL(compiler Compiler) {
	compiler.PushText(compiler.Param(p.value))
}

func (p Parameter) Precedence() int {
	return 0
}

func (p Parameter) ReferencedColumns() []Column {
	return []Column{}
}
