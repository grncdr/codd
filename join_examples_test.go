package codd

import (
	"fmt"
)

func ExampleJoin() {
	BestFriend := Person.As("best_friend")
	BestFriendID := Person.ID.Of(BestFriend)

	fmt.Println(Select(Left.Join(Person, BestFriend, nil)))
	// Output: SELECT person.*, best_friend.* FROM person JOIN person AS best_friend []
	fmt.Println(Select(Left.Join(Person, BestFriend, EQ(Person.BestFriendID, BestFriendID))))
	// SELECT person.*, best_friend.* FROM person JOIN person AS best_friend ON person.best_friend_id = best_friend.id []
}
