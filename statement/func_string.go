package statement

func NewConcat(strs ...ValOrColOrFuncOrSub) Func {
	return Func{
		name: "CONCAT",
		args: ValOrColOrFuncOrSubs(strs).ToArgs(),
	}
}
