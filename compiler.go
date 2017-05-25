package codd

import "fmt"
import "strings"

// Node is the base type implemented by all AST types that can be compiled.
type Node interface {
	Kind() string
	Compile(builder Compiler)
}

// Compiler is the interface used by nodes to push text and/or child nodes.
type Compiler interface {
	Push(node Node)
	PushText(text string)
	Quote(ident Name) string
	Param(value interface{}) (placeholder string)
	Context() []Node
	ContextMatches(pattern string) bool
}

type SQLCompiler interface {
	Compiler
	String() string
	ParamValues() []interface{}
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
	node.Compile(b)
	b.context = b.context[0 : len(b.context)-1]
}

func (b *BaseCompiler) PushText(text string) {
	b.chunks = append(b.chunks, text)
}

// Quote a name, the base compiler does not perform any quoting.
func (b *BaseCompiler) Quote(name Name) string {
	return string(name)
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
