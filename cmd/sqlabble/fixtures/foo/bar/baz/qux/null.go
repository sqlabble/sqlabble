package qux

import (
	"database/sql"
)

// +db:"articles"
type Article struct {
	Bool    sql.NullBool    `db:"bool"`
	Float64 sql.NullFloat64 `db:"float64"`
	Int64   sql.NullInt64   `db:"int64"`
	String  sql.NullString  `db:"string"`
}
