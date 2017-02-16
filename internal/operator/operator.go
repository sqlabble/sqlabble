package operator

type Operator string

const (
	And Operator = "AND"
	Or  Operator = "OR"
	Not Operator = "NOT"

	Eq     Operator = "="
	NotEq  Operator = "!="
	Gt     Operator = ">"
	Gte    Operator = ">="
	Lt     Operator = "<"
	Lte    Operator = "<="
	Like   Operator = "LIKE"
	RegExp Operator = "REGEXP"

	Between Operator = "BETWEEN"

	In    Operator = "IN"
	NotIn Operator = "NOT IN"

	IsNull    Operator = "IS NULL"
	IsNotNull Operator = "IS NOT NULL"

	As Operator = "AS"
)
