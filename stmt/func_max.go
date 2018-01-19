package stmt

import "github.com/sqlabble/sqlabble/keyword"

func NewMax(col ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Max,
		args: Args{col},
	}
}
