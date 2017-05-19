package codd

import "fmt"
import "strings"

type Compiler interface {
	Push(node Node)
	PushText(text string)
	Quote(name Name) string
	Param(value interface{}) (placeholder string)
	String() string
	ParamValues() []interface{}
	Context() []Node
	ContextMatches(pattern string) bool
}

type BaseCompiler struct {
	context []Node
	params  []interface{}
	chunks  []string
}

func (b *BaseCompiler) ContextMatches(pattern string) bool {
	for _, parent := range b.context {
		// TODO - more advanced pattern matching maybe
		if parent.Kind() == pattern {
			return true
		}
	}
	return false
}

func (b *BaseCompiler) Context() []Node {
	return b.context
}

func (b *BaseCompiler) Push(node Node) {
	b.context = append(b.context, node)
	// fmt.Printf("%s> %s %T\n", strings.Repeat("-", len(b.context)), node.Kind(), node)
	node.SQL(b)
	b.context = b.context[0 : len(b.context)-1]
}

func (b *BaseCompiler) PushText(text string) {
	b.chunks = append(b.chunks, text)
}

// this is incorrect in case of names containing double quotes
func (b *BaseCompiler) Quote(name Name) string {
	return fmt.Sprintf("%q", name)
}

func (b *BaseCompiler) Param(value interface{}) string {
	b.params = append(b.params, value)
	return fmt.Sprintf("$%d", len(b.params))
}

func (b *BaseCompiler) String() string {
	return strings.Join(b.chunks, "")
}

func (b *BaseCompiler) ParamValues() []interface{} {
	return b.params
}
