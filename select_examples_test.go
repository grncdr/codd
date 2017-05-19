package codd

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	InitPersonTable()
	m.Run()
}

func ExampleSelect() {
	fmt.Println(Select(Person))
	// Output: SELECT person.* FROM person []
}

func ExampleProject() {
	fmt.Println(Select(Project(Person.ID, Person.Email)))
	// Output: SELECT person.id, person.email FROM person []
}

// This example shows how Project will create a "FROM" list containing all
// referenced tables.
func ExampleProject_referencingMultipleTables() {
	BestFriend := Person.As("best_friend")
	BestFriendID := Person.ID.Of(BestFriend)
	p := Project(
		Person.ID,
		BestFriendID,
	)

	fmt.Println(Select(p))
	// Output: SELECT person.id, best_friend.id FROM person, person AS best_friend []
}
