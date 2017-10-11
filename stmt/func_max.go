package stmt

import "github.com/minodisk/sqlabble/keyword"

func NewMax(col ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Max,
		args: Args{col},
	}
}
