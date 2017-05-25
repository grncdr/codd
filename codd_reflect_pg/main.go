package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/grncdr/codd/generator"
	_ "github.com/lib/pq"
)

var connString string
var searchPath string

var intRegexp = regexp.MustCompile("^(big|small|tiny)?(int(eger|8|4|2)?|serial)$")
var timeRegexp = regexp.MustCompile("^timestamp with(out)? time zone$")
var jsonRegexp = regexp.MustCompile("^jsonb?")
var tsvecRegexp = regexp.MustCompile("tsvector")

func main() {
	flag.Parse()
	connString := flag.Arg(0)
	db, err := sql.Open("postgres", connString)
	fail(err, "Could not open database %q", connString)
	fail(db.Ping(), "Could not ping database %q", connString)
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = $1", "public")
	fail(err, "Could not query information_schema.tables")
	tables := []gen.Table{}
	for rows.Next() {
		table := gen.Table{}
		err := rows.Scan(&table.Name)
		if err != nil {
			fail(err, "Could not scan table row")
		}
		colRows, err := db.Query(
			"SELECT column_name, data_type FROM information_schema.columns WHERE table_schema = $1 AND table_name = $2",
			"public", table.Name,
		)
		if err != nil {
			fail(err, "Could not query columns for table %q", table.Name)
		}
		for colRows.Next() {
			column := gen.Column{}
			colRows.Scan(&column.Name, &column.Type)
			table.Columns = append(table.Columns, column)
		}
		tables = append(tables, table)
	}

	config := gen.Config{
		PackageName: "mypackage",
		Imports:     []string{"github.com/grncdr/pg"},
		Tables:      tables,
		Writer:      os.Stdout,
		ColumnType:  columnType,
	}

	fail(gen.Render(config), "Template rendering")
}

func columnType(dbType string) string {
	switch {
	case intRegexp.MatchString(dbType):
		return "codd.IntegerColumn"
	case timeRegexp.MatchString(dbType):
		return "codd.TimeColumn"
	case dbType == "boolean":
		return "codd.BooleanColumn"
	case dbType == "numeric":
		return "codd.NumericColumn"
	case dbType == "text" || dbType == `"char"`:
		return "codd.TextColumn"
	// Postgres specific types
	case jsonRegexp.MatchString(dbType):
		return "pg.JSONColumn"
	case dbType == "tsvector":
		return "pg.TSVectorColumn"
	default:
		return "codd.Column"
	}
}

func fail(err error, msg string, params ...interface{}) {
	if err == nil {
		return
	}
	params = append(params, err)
	fmt.Printf(msg+": %s\n", params...)
	os.Exit(3)
}
