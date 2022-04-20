package acr

import (
	"bufio"
	"os"
	"strings"
)

type Table struct {
	Table string
}

type Field struct {
	Field, Table, Ident string
}

type Template struct {
	Template, Pattern, Sep string
}

type Xref struct {
	Xref, Table string
}

type DB struct {
	Table    []Table
	Field    []Field
	Template []Template
	Xref     []Xref
}

func (db *DB) readFile(fi *os.File) error {
	s := bufio.NewScanner(fi)
	for s.Scan() {
		header := strings.Split(s.Text(), "\t")
		for s.Scan() {
			row := strings.Split(s.Text(), "\t")
			if row[0] == "" {
				break
			}

			db.readRow(header, row)
		}
	}

	return nil
}

func (db *DB) readRow(header, row []string) error {
	switch header[0] {
	case "table":
		db.Table = append(db.Table, Table{Table: row[0]})
	case "field":
		field := strings.Split(row[0], ".")
		db.Field = append(db.Field, Field{Field: row[0], Table: field[0], Ident: field[1]})
	case "template":
		db.Template = append(db.Template, Template{})
	case "xref":
		db.Xref = append(db.Xref, Xref{Xref: row[0]})
	}
	return nil
}

func NewDB(files ...*os.File) (*DB, error) {
	db := &DB{}
	for _, file := range files {
		if err := db.readFile(file); err != nil {
			return nil, err
		}
	}

	return db, nil
}
