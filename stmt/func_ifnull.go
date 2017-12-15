package stmt

import "github.com/minodisk/sqlabble/keyword"

func NewIFNULL(v1, v2 ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.IFNULL,
		args: Args{v1, v2},
	}
}
