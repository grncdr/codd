package codd

type ColumnType int32

const (
	ColumnTypeText ColumnType = iota
	ColumnTypeNumeric
	ColumnTypeInteger
	ColumnTypeBigInteger
	ColumnTypeBoolean
	ColumnTypeTime
)

type ReferencesColumns interface {
	ReferencedColumns() []Column
}

type Column interface {
	Named
	Expression
	DBTable() Table
	DBType() ColumnType
	Of(t Table) Column
}

type ColumnConfig struct {
	Self  Column
	Table Table
	Name  Name
}

func (c ColumnConfig) Kind() string {
	return "Column"
}

func (c ColumnConfig) Precedence() int {
	return 0
}

func (c ColumnConfig) DBName() Name {
	return c.Name
}

func (c ColumnConfig) DBTable() Table {
	return c.Table
}

func (c ColumnConfig) Compile(compiler Compiler) {
	// TODO use context here to decide when to print fully qualified
	compiler.Push(c.Table)
	compiler.PushText(".")
	compiler.Push(c.Name)
}

func (c ColumnConfig) As(ident Name) ExprAlias {
	return ExprAlias{c.Self, ident}
}

func (c ColumnConfig) ReferencedColumns() []Column {
	return []Column{c.Self}
}

type StarColumn struct {
	ColumnConfig
}

func Star(t Table) (c StarColumn) {
	c.Table = t
	c.Name = "*"
	c.Self = &c
	return
}

func (c StarColumn) Of(t Table) Column {
	return Star(t)
}

func (c StarColumn) DBType() ColumnType {
	return 0
}

type TextColumn struct {
	ColumnConfig
}

func (c TextColumn) DBType() ColumnType {
	return ColumnTypeText
}

func (c TextColumn) Of(t Table) Column {
	r := TextColumn{}
	r.Table = t
	r.Name = c.DBName()
	r.Self = &r
	return r
}

type NumericColumn struct {
	ColumnConfig
}

func (c NumericColumn) DBType() ColumnType {
	return ColumnTypeNumeric
}

func (c NumericColumn) Of(t Table) Column {
	r := NumericColumn{}
	r.Table = t
	r.Name = c.DBName()
	r.Self = &r
	return r
}

type IntegerColumn struct {
	ColumnConfig
}

func (c IntegerColumn) DBType() ColumnType {
	return ColumnTypeInteger
}

func (c IntegerColumn) Of(t Table) Column {
	r := IntegerColumn{}
	r.Table = t
	r.Name = c.DBName()
	r.Self = &r
	return r
}

type BooleanColumn struct {
	ColumnConfig
}

func (c BooleanColumn) DBType() ColumnType {
	return ColumnTypeBoolean
}

func (c BooleanColumn) Of(t Table) Column {
	r := BooleanColumn{}
	r.Table = t
	r.Name = c.DBName()
	r.Self = &r
	return r
}

func (c BooleanColumn) IsBoolean() {}

type TimeColumn struct {
	ColumnConfig
}

func (c TimeColumn) DBType() ColumnType {
	return ColumnTypeTime
}

func (c TimeColumn) Of(t Table) Column {
	r := TimeColumn{}
	r.Table = t
	r.Name = c.DBName()
	r.Self = &r
	return r
}
