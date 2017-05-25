package codd

type Named interface {
	DBName() Name
}

type Name string

func (name Name) Kind() string {
	return "Identifier"
}

func (name Name) Name() Name {
	return name
}

func (name Name) Compile(compiler Compiler) {
	compiler.PushText(compiler.Quote(name))
}
