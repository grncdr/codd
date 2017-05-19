package codd

import "fmt"

var Things ThingTable

type ThingTable struct {
	TableConfig
	ID IntegerColumn
}

func ExampleTableConfig() {
	// This sort of setup code can be generated from a DB schema
	Things.Name = "things"
	Things.Self = &Things
	Things.ID.Table = &Things
	Things.ID.Name = "id"
	Things.ID.Self = &Things.ID

	fmt.Println(Select(Things))
	// Output: SELECT things.* FROM things []
	fmt.Println(Select(Project(Things.ID)))
	// SELECT things.id FROM things []
}
