package codd

type Join struct {
	left  RelationNode
	typ   JoinType
	right RelationNode
	cond  Node
}

func (spec Join) joinable() {}

func (spec Join) Kind() string {
	return "Join"
}

func (join Join) ReferencedTables() []Table {
	var result []Table
	result = append(result, join.left.ReferencedTables()...)
	result = append(result, join.right.ReferencedTables()...)
	return result
}

func (join Join) Projection() []Projectable {
	result := join.left.Projection()
	result = append(result, join.right.Projection()...)
	return result
}

func (join Join) FromList() Node {
	return join
}

func (join Join) Restriction() Boolean {
	left := join.left.Restriction()
	right := join.right.Restriction()
	if left != nil && right != nil {
		return And(left, right)
	} else if left != nil {
		return left
	} else if right != nil {
		return right
	}
	return nil
}

func (join Join) Ordering() Node {
	return join.left.Ordering()
}

func (join Join) SQL(builder Compiler) {
	builder.Push(join.left)
	builder.Push(join.typ)
	builder.Push(join.right)
	if join.cond != nil {
		builder.PushText(" ON ")
		builder.Push(JoinCondition{join.cond})
	}
}

type JoinCondition struct {
	Node
}

func (cond JoinCondition) Kind() string {
	return "JoinCondition"
}
