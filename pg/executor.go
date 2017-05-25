package pg

import (
	"database/sql"

	"github.com/grncdr/codd"
	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

type Tx struct {
	tx *sql.Tx
}

type handle interface {
	Query(text string, params ...interface{}) (*sql.Rows, error)
}

func Open(connString string) (DB, error) {
	db, err := sql.Open("postgres", connString)
	return DB{db}, err
}

func (db DB) Query(query codd.Query) (*sql.Rows, error) {
	return exec(db.db, query)
}

func (db DB) Begin() (Tx, error) {
	tx, err := db.db.Begin()
	return Tx{tx}, err
}

func (tx Tx) Query(query codd.Query) (*sql.Rows, error) {
	return exec(tx.tx, query)
}

func (tx Tx) Commit() error {
	return tx.tx.Commit()
}

func (tx Tx) Rollback() error {
	return tx.tx.Rollback()
}

func exec(handle handle, query codd.Query) (*sql.Rows, error) {
	compiler := &codd.BaseCompiler{}
	compiler.Push(query)
	return handle.Query(compiler.String(), compiler.ParamValues()...)
}
