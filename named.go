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

func (name Name) Compile(builder Compiler) {
	builder.PushText(string(name))
}
