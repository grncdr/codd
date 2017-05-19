package codd

type List struct {
	nodes  []Node
	kind   string
	glue   string
	before string
	after  string
}

func NewList(kind, before, glue, after string) List {
	return List{kind: kind, before: before, glue: glue, after: after}
}

func (list List) Kind() string {
	return list.kind
}

func (list List) Compile(builder Compiler) {
	for i, node := range list.nodes {
		if i != 0 && list.glue != "" {
			builder.PushText(list.glue)
		}
		builder.Push(node)
	}
}

func (list *List) Push(nodes ...Node) {
	list.nodes = append(list.nodes, nodes...)
}
