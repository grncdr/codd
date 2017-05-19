package codd

/*
type SchemaBuilder struct {
	name   Name
	tables []*TableBuilder
}

type TableBuilder struct {
	name    Name
	columns []*ColumnBuilder
}

type ColumnBuilder struct {
	name   Name
	dbType ColumnType
}

func NewSchema(name Name) *SchemaBuilder {
	return &SchemaBuilder{name: name}
}

func (sb *SchemaBuilder) AddTable(name Name) *TableBuilder {
	// TODO - validate that table does not already exist
	table := &TableBuilder{
		name: name,
	}
	sb.tables = append(sb.tables, table)
	return table
}

func (tb *TableBuilder) AddColumn(name Name, typ ColumnType) *ColumnBuilder {
	col := &ColumnBuilder{
		name:   name,
		dbType: typ,
	}
	tb.columns = append(tb.columns, col)
	return col
}

func (tb *TableBuilder) Table(schema *Schema) UnsafeTable {
	t := UnsafeTable{TableConfig{nil, tb.name}, make([]Column, len(tb.columns))}
	t.Self = &t
	for i, col := range tb.columns {
		t.columns[i] = col.Column(&t)
	}
	return t
}

func (c *ColumnBuilder) Column(table *UnsafeTable) Column {
	return Column{table, c.name, c.dbType}
}

func (sb *SchemaBuilder) Schema() Schema {
	schema := Schema{sb.name, make([]UnsafeTable, len(sb.tables))}
	for i := range sb.tables {
		schema.tables[i] = sb.tables[i].Table(&schema)
	}
	return schema
}
*/
