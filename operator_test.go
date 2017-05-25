package codd

func ExampleExpression() {
	x := Add(Person.ID, Person.ID)
	y := Sub(Person.ID, Person.ID)
	z := GT(Mul(x, y), Param(3)).As("z")
	PrintQuery(Select(Project(z)))
	// Output: SELECT (person.id + person.id) * (person.id - person.id) > $1 AS z FROM person [3]
}
