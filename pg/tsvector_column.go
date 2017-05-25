package pg

import "github.com/grncdr/codd"

type TSVectorColumn struct {
	codd.ColumnConfig
}

func (c TSVectorColumn) DBType() codd.ColumnType {
	return -1
}

func (c TSVectorColumn) Of(t codd.Table) codd.Column {
	r := TSVectorColumn{}
	r.Table = t
	r.Name = c.DBName()
	r.Self = &r
	return r
}

func (c TSVectorColumn) Holder() interface{} {
	holder := ""
	return &holder
}
