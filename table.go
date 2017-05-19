package codd

// Table should be implemented by types in your program representing tables in
// your database. In addition to being a Relation, a table is also a Node (meaning
// it can be compiled into a SQL string), has a name, and can be aliased to a new
// name.
type Table interface {
	Relation
	Node
	Named
	As(name Name) TableAlias
}

type ReferencesTables interface {
	ReferencedTables() []Table
}

// TableConfig implements the table interface and can be embedded in other
// structs to colocate the table with it's columns.
type TableConfig struct {
	//Schema    *Schema
	Self Table
	Name Name
}

func (t TableConfig) Kind() string {
	return "Table"
}

func (t TableConfig) FromList() Node {
	return t.Self
}

func (t TableConfig) ReferencedTables() []Table {
	return []Table{t.Self}
}

func (t TableConfig) Compile(builder Compiler) {
	builder.Push(t.Name)
}

func (t TableConfig) DBName() Name {
	return t.Name
}

// Projection returns the a single item list containing the star projection of this
// table. Use codd.Project(col1, col2) to select individual columns.
func (t TableConfig) Projection() []Projectable {
	return []Projectable{Star(t.Self)}
}

// Restriction returns nil, as tables don't have any implied "WHERE" clause.
func (t TableConfig) Restriction() Boolean {
	return nil
}

// Ordering returns nil, as tables don't have any implied "ORDER BY" clause.
func (t TableConfig) Ordering() Node {
	return nil
}

// As creates a TableAlias for this table with the given name.
func (t TableConfig) As(ident Name) TableAlias {
	return TableAlias{t.Self, ident}
}

type UnsafeTable struct {
	TableConfig
	columns []Column
}

func (t UnsafeTable) UnsafeColumn(name Name) Column {
	if name == "*" {
		return Star(t)
	}
	for _, column := range t.columns {
		if column.DBName() == name {
			return column
		}
	}
	return nil
}

func (t UnsafeTable) Columns() []Column {
	result := make([]Column, len(t.columns))
	i := 0
	for _, col := range t.columns {
		result[i] = col
		i++
	}

	return result
}

type TableAlias struct {
	Table
	alias Name
}

func (t TableAlias) Kind() string {
	return "TableAlias"
}

func (t TableAlias) FromList() Node {
	return t
}

func (t TableAlias) ReferencedTables() []Table {
	return []Table{t}
}

func (t TableAlias) DBName() Name {
	return t.alias
}

func (t TableAlias) Compile(builder Compiler) {
	if builder.ContextMatches("FromList") && !builder.ContextMatches("JoinCondition") {
		builder.Push(t.Table)
		builder.PushText(" AS ")
	}
	builder.Push(t.alias)
}

func (t TableAlias) Projection() []Projectable {
	return []Projectable{Star(t)}
}

func (t TableAlias) Column(c Column) Column {
	return c.Of(t)
}
