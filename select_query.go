package codd

type SelectQuery struct {
	Relation
}

func Select(rel Relation) SelectQuery {
	return SelectQuery{rel}
}

func (stmt SelectQuery) Kind() string {
	return "SelectStatement"
}

func (stmt SelectQuery) IsQuery() {}

func (stmt SelectQuery) Compile(builder Compiler) {
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
