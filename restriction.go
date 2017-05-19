package codd

type Restrictable interface {
	ReferencesTables
}

type Restriction struct {
	source    Relation
	condition Boolean
}

func Restrict(source Relation, condition Boolean) Restriction {
	return Restriction{source, condition}
}

func (r Restriction) Kind() string {
	return "Restriction"
}

func (r Restriction) SQL(compiler Compiler) {
	compiler.Push(r.condition)
}

func (r Restriction) ReferencedTables() []Table {
	return r.source.ReferencedTables()
}

func (r Restriction) Projection() []Projectable {
	return r.source.Projection()
}

func (r Restriction) FromList() Node {
	return r.source.FromList()
}

func (r Restriction) Restriction() Boolean {
	before := r.source.Restriction()
	if before != nil {
		return And(before, r.condition)
	}
	return r.condition
}
