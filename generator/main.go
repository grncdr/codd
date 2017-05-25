package gen

import (
	"fmt"
	"io"
	"strings"
	"text/template"
)

type Config struct {
	PackageName string
	Imports     []string
	Tables      []Table
	Writer      io.Writer
	ColumnType  func(dbType string) string
}

type Table struct {
	Name    string
	Columns []Column
}

type Column struct {
	Name string
	Type string
}

var rootTemplate = `
{{define "package"}}
package {{.PackageName}}

import (
	"github.com/grncdr/codd"
	{{range .Imports}}"{{.}}"
	{{end}}
)

var (
	{{range .Tables}}{{toIdent .Name}} {{toIdent .Name}}Table
	{{end}}
)

{{range .Tables}}
{{template "tableType" .}}
{{end}}

func InitTables() {
	{{range .Tables}}
	{{template "initTable" .}}
	{{end}}
}

{{end}}

{{define "tableType"}}
type {{toIdent .Name}}Table struct {
	codd.TableConfig
	{{range .Columns -}}{{toIdent .Name}} {{columnType .Type}} {{columnTag .Name .Type}}
	{{end}}
}{{end}}

{{define "initTable"}}
{{ $tableVar := toIdent .Name }}
{{$tableVar}}.TableConfig.Name = {{printf "%q" .Name}}
{{$tableVar}}.TableConfig.Self = &{{$tableVar}}
{{- range .Columns -}}
{{$columnVar := toIdent .Name }}
{{$tableVar}}.{{$columnVar}}.Table = &{{$tableVar}}
{{$tableVar}}.{{$columnVar}}.Self = &{{$tableVar}}.{{$columnVar}}
{{$tableVar}}.{{$columnVar}}.Name = {{printf "%q" .Name}}{{end}}
{{end}}
`

func Render(config Config) error {
	root, err := template.New("root").
		Funcs(template.FuncMap{
			"toIdent": func(name string) string {
				spaced := strings.Replace(name, "_", " ", -1)
				titled := strings.Title(spaced)
				return strings.Replace(titled, " ", "", -1)
			},
			"columnTag": func(name, ty string) string {
				return fmt.Sprintf("`codd:\"%s,%s\"`", name, ty)
			},
			"columnType": config.ColumnType,
		}).
		Parse(rootTemplate)
	if err != nil {
		return err
	}
	return root.ExecuteTemplate(config.Writer, "package", config)
}
