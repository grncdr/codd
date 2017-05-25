package pg

import "github.com/grncdr/codd"

type JSONColumn struct {
	codd.ColumnConfig
}

func (c JSONColumn) DBType() codd.ColumnType {
	return -1
}

func (c JSONColumn) Of(t codd.Table) codd.Column {
	r := JSONColumn{}
	r.Table = t
	r.Name = c.DBName()
	r.Self = &r
	return r
}
