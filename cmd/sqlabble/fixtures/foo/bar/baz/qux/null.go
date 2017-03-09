package qux

import "database/sql"

// +db:"articles"
type Article struct {
	Bool    sql.NullBool
	Float64 sql.NullFloat64
	Int64   sql.NullInt64
	String  sql.NullString
}
