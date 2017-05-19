package codd

/*
// An identifier is a fully qualified name
type Identifier []Name

func Qualified(names ...Name) Identifier {
	return names
}

func (ident Identifier) Kind() string {
	return "Identifier"
}

func (ident Identifier) Name() Name {
	return ident[len(ident)-1]
}

func (ident Identifier) SQL(builder Compiler) {
	text := ""
	for i, name := range ident {
		if i != 0 {
			text += "."
		}
		text += builder.Quote(name)
	}
	builder.PushText(text)
}
*/
