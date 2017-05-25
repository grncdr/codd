package pg

import (
	"fmt"
	"github.com/grncdr/codd"
)

type Compiler struct {
	codd.BaseCompiler
}

func (c *Compiler) Push(node codd.Node) {
	if node.Kind() == "Name" {
		c.PushText(fmt.Sprintf("%q", node.(codd.Name)))
		return
	}
	c.BaseCompiler.Push(node)
}

func (c *Compiler) Quote(name codd.Name) string {
	return fmt.Sprintf("%q")
}
