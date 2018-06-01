package sqlabble

import (
	"database/sql"

	"github.com/sqlabble/sqlabble/builder"
	"github.com/sqlabble/sqlabble/stmt"
)

// Methods exported to make statements.
var (
	CreateTable            = stmt.NewCreateTable
	CreateTableIfNotExists = stmt.NewCreateTableIfNotExists
	Select                 = stmt.NewSelect
	SelectDistinct         = stmt.NewSelectDistinct
	InsertInto             = stmt.NewInsertInto
	Update                 = stmt.NewUpdate
	Delete                 = stmt.NewDelete

	SimpleCase   = stmt.NewSimpleCase
	SimpleWhen   = stmt.NewSimpleWhen
	SearchedCase = stmt.NewSearchedCase
	SearchedWhen = stmt.NewSearchedWhen

	Column   = stmt.NewColumn
	Table    = stmt.NewTable
	Val      = stmt.NewVal
	Vals     = stmt.NewVals
	Subquery = stmt.NewSubquery

	Assign    = stmt.NewAssign
	And       = stmt.NewAnd
	Or        = stmt.NewOr
	Not       = stmt.NewNot
	Exists    = stmt.NewExists
	NotExists = stmt.NewNotExists

	Union        = stmt.NewUnion
	UnionAll     = stmt.NewUnionAll
	Intersect    = stmt.NewIntersect
	IntersectAll = stmt.NewIntersectAll
	Except       = stmt.NewExcept
	ExceptAll    = stmt.NewExceptAll

	Wildcard = stmt.NewWildcard()

	Count = stmt.NewCount
	Max   = stmt.NewMax

	Add        = stmt.NewAdd
	Sub        = stmt.NewSub
	Mul        = stmt.NewMul
	Div        = stmt.NewDiv
	IntegerDiv = stmt.NewIntegerDiv
	Mod        = stmt.NewMod

	IFNULL = stmt.NewIFNULL
)

// Session wraps sql.DB
type Session struct {
	DBInterface
	Builder *builder.Builder
}

// DBInterface *sql.DB or *sql.Tx
type DBInterface interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// NewSession create a Session instance
func NewSession(builderName string, db DBInterface) (*Session, error) {
	var b *builder.Builder
	switch builderName {
	case "standard":
		b = &builder.Standard
	case "standard_indented":
		b = &builder.StandardIndented
	case "mysql":
		b = &builder.MySQL
	case "mysql_indented":
		b = &builder.MySQLIndented
	default:
		return nil, builder.NewErrBuilderNotSupported(builderName)
	}
	return &Session{
		DBInterface: db,
		Builder:     b,
	}, nil
}
