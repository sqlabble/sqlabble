package stmt

import "github.com/minodisk/sqlabble/keyword"

func NewCount(col ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Count,
		args: Args{col},
	}
}
