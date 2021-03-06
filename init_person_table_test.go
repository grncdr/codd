package codd

var (
	Person PersonTable
)

// This type (and it's corresponding Columns method and InitPersonTable func) are
// generated by reflecting a database schema.
type PersonTable struct {
	TableConfig
	ID           IntegerColumn
	Name         TextColumn
	Email        TextColumn
	BestFriendID TextColumn
}

// This InitPersonTable function is generated via schema reflection
func InitPersonTable() {
	Person.TableConfig.Name = "person"
	Person.TableConfig.Self = &Person
	Person.ID.Table = &Person
	Person.ID.Self = &Person.ID
	Person.ID.Name = "id"
	Person.Name.Table = &Person
	Person.Name.Self = &Person.Name
	Person.Name.Name = "name"
	Person.Email.Table = &Person
	Person.Email.Self = &Person.Email
	Person.Email.Name = "email"
	Person.BestFriendID.Table = &Person
	Person.BestFriendID.Self = &Person.BestFriendID
	Person.BestFriendID.Name = "bestfriend_id"
}
