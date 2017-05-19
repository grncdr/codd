package codd

import (
	"fmt"
)

type Schema struct {
	Name   Name
	tables []UnsafeTable
}

// UnsafeTable gets an UnsafeTable instance from the schema by name. This method
// panics if a table with the given name does not exist. Generally you should only
// call this once per-table in a package init method.
func (s Schema) UnsafeTable(name Name) UnsafeTable {
	for _, table := range s.tables {
		if table.DBName() == name {
			return table
		}
	}
	panic(fmt.Errorf("Table %q is not defined in schema %q", name, s.Name))
}
