package codd

type Projection struct {
	from RelationNode
	cols []Projectable
}

type Projectable interface {
	Node
	Named
	ReferencesColumns
}

// Project the given expressions. Selecting this projection yields a "FROM" list
// containing every Table referenced by the expressions.
func Project(expressions ...Projectable) Projection {
	return Projection{nil, expressions}
}

// ProjectFrom projects the given expressions from a Selectable source such as a
// Join. If source is nil this is equivalent to Project(expressions...).
func ProjectFrom(from RelationNode, expressions ...Projectable) Projection {
	return Projection{from, expressions}
}

// Kind returns "Projection"
func (p Projection) Kind() string {
	return "Projection"
}

func (p Projection) Projection() []Projectable {
	return p.cols
}

func (p Projection) FromList() Node {
	if p.from != nil {
		return p.from.FromList()
	}
	tableList := NewList("TableList", "", ", ", "")
	for _, table := range p.ReferencedTables() {
		tableList.Push(table)
	}
	return tableList
}

func (p Projection) Restriction() Boolean {
	if p.from != nil {
		return p.from.Restriction()
	}
	return nil
}

func (p Projection) Ordering() Node {
	if p.from != nil {
		return p.from.Ordering()
	}
	return nil
}

func (p Projection) ReferencedTables() []Table {
	tables := map[Name]Named{}
	tableList := []Table{}
	for _, x := range p.cols {
		for _, col := range x.ReferencedColumns() {
			table := col.DBTable()
			if _, ok := tables[table.DBName()]; !ok {
				tables[table.DBName()] = table
				tableList = append(tableList, table)
			}
		}
	}
	return tableList
}
