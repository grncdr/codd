package codd

import (
	"fmt"
)

func ExampleSort() {
	fmt.Println(Select(Sort(Person, Person.Name, Ascending)))
	// Output: SELECT person.* FROM person ORDER BY person.name ASC []
}

// You can sort before or after projection, it doesn't matter
func ExampleSort_combinators() {
	fmt.Println(Select(ProjectFrom(
		Sort(Person, Person.Name, Ascending),
		Person.ID,
		Person.Name,
	)))
	// Output: SELECT person.id, person.name FROM person ORDER BY person.name ASC []
	fmt.Println(Select(Sort(
		Project(Person.ID, Person.Name),
		Person.Name,
		Ascending,
	)))
	// SELECT person.id, person.name FROM person ORDER BY person.name ASC []
}
