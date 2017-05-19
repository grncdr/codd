package codd

type columnList []Node

func (columns columnList) Kind() string {
	return "ColumnList"
}

func (columns columnList) SQL(builder Compiler) {
	for i, col := range columns {
		if i != 0 {
			builder.PushText(", ")
		}
		builder.Push(col)
	}
}
