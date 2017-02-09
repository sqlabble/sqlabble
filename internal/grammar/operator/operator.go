package operator

type Operator string

const (
	And Operator = "AND"
	Or  Operator = "OR"
	Not Operator = "NOT"

	Equal    Operator = "="
	NotEqual Operator = "!="
	Gt       Operator = ">"
	Gte      Operator = ">="
	Lt       Operator = "<"
	Lte      Operator = "<="
	Between  Operator = "BETWEEN"

	In    Operator = "IN"
	NotIn Operator = "NOT IN"

	Like   Operator = "LIKE"
	RegExp Operator = "REGEXP"

	IsNull    Operator = "IS NULL"
	IsNotNull Operator = "IS NOT NULL"
)
