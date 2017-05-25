package codd

import (
	"fmt"
)

func ExampleSort() {
	PrintQuery(Select(Sort(Person, Person.Name, Ascending)))
	// Output: SELECT person.* FROM person ORDER BY person.name ASC []
}

func ExampleSort_orderIndependence() {
	fmt.Println()
	PrintQuery(Select(ProjectFrom(
		Sort(Person, Person.Name, Ascending),
		Person.ID,
		Person.Name,
	)))

	PrintQuery(Select(Sort(
		Project(Person.ID, Person.Name),
		Person.Name,
		Ascending,
	)))

	// Output:
	// SELECT person.id, person.name FROM person ORDER BY person.name ASC []
	// SELECT person.id, person.name FROM person ORDER BY person.name ASC []
}
