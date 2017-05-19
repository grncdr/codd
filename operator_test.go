package codd

import (
	"testing"
)

func TestArithmetic(t *testing.T) {
	x := Add(Person.ID, Person.ID)
	y := Sub(Person.ID, Person.ID)
	z := GT(Mul(x, y), Param(3)).As("z")
	t.Log(Select(Project(z)))
}
