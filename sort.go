package codd

type OrderDirection string

const (
	Ascending  OrderDirection = " ASC"
	Descending OrderDirection = " DESC"
)

type OrderBy struct {
	rel  Relation
	expr Expression
	dir  OrderDirection
}

func Sort(rel Relation, expr Expression, dir OrderDirection) OrderBy {
	return OrderBy{rel, expr, dir}
}

func (o OrderBy) Kind() string {
	return "OrderBy"
}

func (o OrderBy) Compile(compiler Compiler) {
	compiler.Push(o.expr)
	compiler.PushText(string(o.dir))
	if secondary := o.rel.Ordering(); secondary != nil {
		compiler.PushText(", ")
		compiler.Push(secondary)
	}
}

func (o OrderBy) ReferencedTables() []Table {
	return o.rel.ReferencedTables()
}

func (o OrderBy) Projection() []Projectable {
	return o.rel.Projection()
}

func (o OrderBy) FromList() Node {
	return o.rel.FromList()
}

func (o OrderBy) Restriction() Boolean {
	return o.rel.Restriction()
}

func (o OrderBy) Ordering() Node {
	return o
}
