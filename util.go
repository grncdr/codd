package codd

type commas []Node

func (nodeList commas) Kind() string {
	return "Commas"
}

func (nodeList commas) SQL(builder Compiler) {
	for i := 0; i < len(nodeList); i++ {
		if i != 0 {
			builder.PushText(",")
		}
		builder.Push(nodeList[i])
	}
}

func Commas(nodes ...Node) commas {
	return commas(nodes)
}
