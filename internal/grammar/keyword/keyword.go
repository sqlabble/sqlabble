package keyword

type Clause string

const (
	Select  Clause = "SELECT"
	From    Clause = "FROM"
	Where   Clause = "WHERE"
	OrderBy Clause = "ORDER BY"
)

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
	CreateTable = "CREATE TABLE"
	InsertInto  = "INSERT INTO"
	Values      = "VALUES"
	Update      = "UPDATE"
	Set         = "SET"
)
