package stmt

import "github.com/sqlabble/sqlabble/keyword"

func NewIFNULL(expr1, expr2 ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.IFNULL,
		args: Args{expr1, expr2},
	}
}
