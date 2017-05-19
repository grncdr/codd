package codd

type JoinType string

func (j JoinType) Kind() string {
	return "JoinType"
}

func (j JoinType) Compile(builder Compiler) {
	builder.PushText(string(j))
}

func (j JoinType) Join(left RelationNode, right RelationNode, cond Node) Join {
	return Join{left, j, right, cond}
}

const (
	Inner JoinType = " JOIN "
	Left  JoinType = " LEFT JOIN "
	Right JoinType = " RIGHT JOIN "
	Outer JoinType = " FULL OUTER JOIN "
	Union JoinType = " UNION "
)
