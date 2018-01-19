package stmt

import "github.com/sqlabble/sqlabble/keyword"

func NewCount(col ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Count,
		args: Args{col},
	}
}
