package codd

type selectStatement struct {
	Relation
}

func Select(source Relation) (string, []interface{}) {
	builder := DefaultSQLCompiler()
	builder.Push(selectStatement{source})
	return builder.String(), builder.ParamValues()
}

func (stmt selectStatement) Kind() string {
	return "SelectStatement"
}

func (stmt selectStatement) Compile(builder Compiler) {
	builder.PushText("SELECT ")

	for i, field := range stmt.Projection() {
		if i != 0 {
			builder.PushText(", ")
		}
		builder.Push(field)
	}

	builder.Push(FromList{stmt.FromList()})

	if r := stmt.Restriction(); r != nil {
		builder.PushText(" WHERE ")
		builder.Push(r)
	}

	if o := stmt.Ordering(); o != nil {
		builder.PushText(" ORDER BY ")
		builder.Push(o)
	}
}
