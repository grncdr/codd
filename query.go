package codd

// Query is a marker interface that indicates a node can be compiled to an
// independent SQL query.
type Query interface {
	Node
	IsQuery()
}
