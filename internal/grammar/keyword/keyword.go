package keyword

type Clause string

const (
	Select      Clause = "SELECT"
	From        Clause = "FROM"
	Where       Clause = "WHERE"
	OrderBy     Clause = "ORDER BY"
	InsertInto  Clause = "INSERT INTO"
	CreateTable Clause = "CREATE TABLE"
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
	Values = "VALUES"
)
