package codd

// Node is the base type impleented by all AST node types.
type Node interface {
	Kind() string
	SQL(builder Compiler)
}
