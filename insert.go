package codd

import "fmt"

type InsertTarget interface {
	Node
	Named
	Columns() []Column
}

type insertStatement struct {
	target    InsertTarget
	returning []Node
	data      []map[Name]interface{}
}

func Insert(t InsertTarget, returning []Column, rows ...map[Name]interface{}) (string, []interface{}, error) {
	if len(rows) == 0 {
		return "", nil, fmt.Errorf("Insert requires at least one row")
	}

	stmt := insertStatement{target: t}
	var unreturnable []Name
	for _, col := range returning {
		if col.DBTable().DBName() != t.DBName() {
			unreturnable = append(unreturnable, col.DBName())
		} else {
			stmt.returning = append(stmt.returning, col.DBName())
		}
	}
	if len(unreturnable) > 0 {
		return "", nil, fmt.Errorf("Cannot return columns %s from table %q", returning, t.DBName())
	}
	unreturnable = nil

	builder := BaseCompiler{}
	builder.Push(stmt)
	return builder.String(), builder.ParamValues(), nil
}

func (stmt insertStatement) Kind() string {
	return "InsertStatement"
}

func (stmt insertStatement) Compile(builder Compiler) {
	builder.PushText("INSERT INTO ")
	builder.Push(stmt.target)
	builder.PushText("(")
	columns := stmt.target.Columns()
	for i, col := range columns {
		if i != 0 {
			builder.PushText(", ")
		}
		builder.Push(col)
	}
	builder.PushText(") VALUES ")
	for i, rowData := range stmt.data {
		builder.PushText("(")
		if i != 0 {
			builder.PushText(",")
		}
		for j, col := range columns {
			if j != 0 {
				builder.PushText(",")
			}
			if val, ok := rowData[col.DBName()]; ok {
				builder.PushText(builder.Param(val))
			} else {
				builder.PushText("DEFAULT")
			}
		}
		builder.PushText(")")
	}
}
