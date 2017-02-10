package keyword

type JoinType string

const (
	Join      JoinType = "JOIN"
	InnerJoin JoinType = "INNER JOIN"
	LeftJoin  JoinType = "LEFT JOIN"
	RightJoin JoinType = "RIGHT JOIN"
)

type JoinCondition string

const (
	On    JoinCondition = "ON"
	Using JoinCondition = "USING"
)

const (
	Select      = "SELECT"
	From        = "FROM"
	Where       = "WHERE"
	GroupBy     = "GROUP BY"
	OrderBy     = "ORDER BY"
	Limit       = "LIMIT"
	InsertInto  = "INSERT INTO"
	Values      = "VALUES"
	Update      = "UPDATE"
	Set         = "SET"
	Delete      = "DELETE"
	CreateTable = "CREATE TABLE"
)
