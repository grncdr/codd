package codd

type FromList struct {
	from Node
}

func (fl FromList) Kind() string {
	return "FromList"
}

func (fl FromList) Compile(builder Compiler) {
	builder.PushText(" FROM ")
	builder.Push(fl.from)
}
