package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type JoinOperation struct {
	op  keyword.Operator
	ops []ComparisonOrLogicalOperation
}

func NewAnd(ops ...ComparisonOrLogicalOperation) JoinOperation {
	return JoinOperation{
		op:  keyword.And,
		ops: ops,
	}
}

func NewOr(ops ...ComparisonOrLogicalOperation) JoinOperation {
	return JoinOperation{
		op:  keyword.Or,
		ops: ops,
	}
}

func (o JoinOperation) nodeize() (tokenizer.Tokenizer, []interface{}) {
	ts := make(tokenizer.Tokenizers, len(o.ops))
	values := []interface{}{}
	for i, op := range o.ops {
		t, vals := op.nodeize()
		if _, ok := op.(JoinOperation); ok {
			t = tokenizer.NewParentheses(t)
		}
		ts[i] = t
		values = append(values, vals...)
	}
	return tokenizer.NewTokenizers(ts...).Prefix(
		token.Word(o.keyword()),
		token.Space,
	), values
}

func (o JoinOperation) keyword() keyword.Operator {
	return o.op
}

func (o JoinOperation) operations() []ComparisonOrLogicalOperation {
	return o.ops
}

type Not struct {
	operation ComparisonOrLogicalOperation
}

func NewNot(operation ComparisonOrLogicalOperation) Not {
	return Not{operation: operation}
}

func (o Not) nodeize() (tokenizer.Tokenizer, []interface{}) {
	middle, values := o.operation.nodeize()
	return tokenizer.NewParentheses(
		middle,
	).Prepend(
		token.Word(o.keyword()),
		token.Space,
	), values
}

func (o Not) keyword() keyword.Operator {
	return keyword.Not
}

func (o Not) operations() []ComparisonOrLogicalOperation {
	return []ComparisonOrLogicalOperation{o.operation}
}

type ComparisonOperation struct {
	op     keyword.Operator
	column ValOrColOrFuncOrSub
	param  ValOrColOrFuncOrSub
}

func NewEq(param ValOrColOrFuncOrSub) ComparisonOperation {
	return ComparisonOperation{
		op:    keyword.Eq,
		param: param,
	}
}

func NewNotEq(param ValOrColOrFuncOrSub) ComparisonOperation {
	return ComparisonOperation{
		op:    keyword.NotEq,
		param: param,
	}
}

func NewGt(param ValOrColOrFuncOrSub) ComparisonOperation {
	return ComparisonOperation{
		op:    keyword.Gt,
		param: param,
	}
}

func NewGte(param ValOrColOrFuncOrSub) ComparisonOperation {
	return ComparisonOperation{
		op:    keyword.Gte,
		param: param,
	}
}

func NewLt(param ValOrColOrFuncOrSub) ComparisonOperation {
	return ComparisonOperation{
		op:    keyword.Lt,
		param: param,
	}
}

func NewLte(param ValOrColOrFuncOrSub) ComparisonOperation {
	return ComparisonOperation{
		op:    keyword.Lte,
		param: param,
	}
}

func NewLike(param ValOrColOrFuncOrSub) ComparisonOperation {
	return ComparisonOperation{
		op:    keyword.Like,
		param: param,
	}
}

func NewRegExp(param ValOrColOrFuncOrSub) ComparisonOperation {
	return ComparisonOperation{
		op:    keyword.RegExp,
		param: param,
	}
}

func (o ComparisonOperation) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t1, v1 := o.column.nodeize()
	t2, v2 := o.param.nodeize()
	return tokenizer.ConcatTokenizers(
		t1,
		t2,
		tokenizer.NewLine(
			token.Space,
			token.Word(o.keyword()),
			token.Space,
		),
	), append(v1, v2...)
}

func (o ComparisonOperation) keyword() keyword.Operator {
	return o.op
}

type Between struct {
	column   ValOrColOrFuncOrSub
	from, to ValOrColOrFuncOrSub
}

func NewBetween(from, to ValOrColOrFuncOrSub) Between {
	return Between{
		from: from,
		to:   to,
	}
}

func (o Between) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t1, v1 := o.column.nodeize()
	t2, v2 := o.from.nodeize()
	t3, v3 := o.to.nodeize()
	return tokenizer.ConcatTokenizers(
		tokenizer.ConcatTokenizers(
			t1,
			t2,
			tokenizer.NewLine(
				token.Space,
				token.Word(o.keyword()),
				token.Space,
			),
		),
		t3,
		tokenizer.NewLine(
			token.Space,
			token.Word(keyword.And),
			token.Space,
		),
	), append(append(v1, v2...), v3...)
}

func (o Between) keyword() keyword.Operator {
	return keyword.Between
}

type ContainingOperation struct {
	op     keyword.Operator
	column ValOrColOrFuncOrSub
	params ValsOrSub
}

func NewIn(params ValsOrSub) ContainingOperation {
	return ContainingOperation{
		op:     keyword.In,
		params: params,
	}
}

func NewNotIn(vals ValsOrSub) ContainingOperation {
	return ContainingOperation{
		op:     keyword.NotIn,
		params: vals,
	}
}

func (o ContainingOperation) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t1, v1 := o.column.nodeize()
	t2, v2 := o.params.nodeize()
	return tokenizer.ConcatTokenizers(
		t1,
		t2,
		tokenizer.NewLine(
			token.Space,
			token.Word(o.keyword()),
			token.Space,
		),
	), append(v1, v2...)
}

func (o ContainingOperation) keyword() keyword.Operator {
	return o.op
}

type NullOperation struct {
	op     keyword.Operator
	column ColOrSub
}

func NewIsNull() NullOperation {
	return NullOperation{
		op: keyword.Is,
	}
}

func NewIsNotNull() NullOperation {
	return NullOperation{
		op: keyword.IsNot,
	}
}

func (o NullOperation) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t1, v1 := o.column.nodeize()
	return tokenizer.ConcatTokenizers(
		t1,
		tokenizer.NewLine(token.Word(keyword.Null)),
		tokenizer.NewLine(
			token.Space,
			token.Word(o.keyword()),
			token.Space,
		),
	), v1
}

func (o NullOperation) keyword() keyword.Operator {
	return o.op
}
