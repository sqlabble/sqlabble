package node

// Node is a block of queries that make up SQL.
// It constructs a character string for constructing
// an SQL statement based on Context.
type Node interface {
	ToSQL(Context) (string, []interface{})
}
