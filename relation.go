package codd

// Relation is the interface implemented by all "table-shaped things"
type Relation interface {
	//distinct() Node
	ReferencesTables
	Projection() []Projectable
	FromList() Node
	Restriction() Boolean
	Ordering() Node
	//group() Node
}

type RelationNode interface {
	Node
	Relation
}
