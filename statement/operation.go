package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/operator"
	"github.com/minodisk/sqlabble/token"
)

type JoinOperation struct {
	op  operator.Operator
	ops []ComparisonOrLogicalOperation
}

func NewAnd(ops ...ComparisonOrLogicalOperation) JoinOperation {
	return JoinOperation{
		op:  operator.And,
		ops: ops,
	}
}

func NewOr(ops ...ComparisonOrLogicalOperation) JoinOperation {
	return JoinOperation{
		op:  operator.Or,
		ops: ops,
	}
}

func (o JoinOperation) nodeize() (token.Tokenizer, []interface{}) {
	ts := make(token.Tokenizers, len(o.ops))
	values := []interface{}{}
	for i, op := range o.ops {
		t, vals := op.nodeize()
		if _, ok := op.(JoinOperation); ok {
			t = token.NewParentheses(t)
		}
		ts[i] = t
		values = append(values, vals...)
	}
	return token.NewTokenizers(ts...).Prefix(
		token.Word(o.operator()),
		token.Space,
	), values
}

func (o JoinOperation) operator() operator.Operator {
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

func (o Not) nodeize() (token.Tokenizer, []interface{}) {
	middle, values := o.operation.nodeize()
	return token.NewParentheses(
		middle,
	).Prepend(
		token.Word(o.operator()),
		token.Space,
	), values
}

func (o Not) operator() operator.Operator {
	return operator.Not
}

func (o Not) operations() []ComparisonOrLogicalOperation {
	return []ComparisonOrLogicalOperation{o.operation}
}

type ComparisonOperation struct {
	op     operator.Operator
	column ColumnOrSubquery
	param  ParamOrSubquery
}

func NewEq(param ParamOrSubquery) ComparisonOperation {
	return ComparisonOperation{
		op:    operator.Eq,
		param: param,
	}
}

func NewNotEq(param ParamOrSubquery) ComparisonOperation {
	return ComparisonOperation{
		op:    operator.NotEq,
		param: param,
	}
}

func NewGt(param ParamOrSubquery) ComparisonOperation {
	return ComparisonOperation{
		op:    operator.Gt,
		param: param,
	}
}

func NewGte(param ParamOrSubquery) ComparisonOperation {
	return ComparisonOperation{
		op:    operator.Gte,
		param: param,
	}
}

func NewLt(param ParamOrSubquery) ComparisonOperation {
	return ComparisonOperation{
		op:    operator.Lt,
		param: param,
	}
}

func NewLte(param ParamOrSubquery) ComparisonOperation {
	return ComparisonOperation{
		op:    operator.Lte,
		param: param,
	}
}

func NewLike(param ParamOrSubquery) ComparisonOperation {
	return ComparisonOperation{
		op:    operator.Like,
		param: param,
	}
}

func NewRegExp(param ParamOrSubquery) ComparisonOperation {
	return ComparisonOperation{
		op:    operator.RegExp,
		param: param,
	}
}

func (o ComparisonOperation) nodeize() (token.Tokenizer, []interface{}) {
	t1, v1 := o.column.nodeize()
	t2, v2 := o.param.nodeize()
	return token.ConcatTokenizers(
		t1,
		t2,
		token.NewLine(
			token.Space,
			token.Word(o.operator()),
			token.Space,
		),
	), append(v1, v2...)
}

func (o ComparisonOperation) operator() operator.Operator {
	return o.op
}

type Between struct {
	column   ColumnOrSubquery
	from, to ParamOrSubquery
}

func NewBetween(from, to ParamOrSubquery) Between {
	return Between{
		from: from,
		to:   to,
	}
}

func (o Between) nodeize() (token.Tokenizer, []interface{}) {
	t1, v1 := o.column.nodeize()
	t2, v2 := o.from.nodeize()
	t3, v3 := o.to.nodeize()
	return token.ConcatTokenizers(
		token.ConcatTokenizers(
			t1,
			t2,
			token.NewLine(
				token.Space,
				token.Word(o.operator()),
				token.Space,
			),
		),
		t3,
		token.NewLine(
			token.Space,
			token.Word(operator.And),
			token.Space,
		),
	), append(append(v1, v2...), v3...)
}

func (o Between) operator() operator.Operator {
	return operator.Between
}

type ContainingOperation struct {
	op     operator.Operator
	column ColumnOrSubquery
	params ParamsOrSubquery
}

func NewIn(params ParamsOrSubquery) ContainingOperation {
	return ContainingOperation{
		op:     operator.In,
		params: params,
	}
}

func NewNotIn(vals ParamsOrSubquery) ContainingOperation {
	return ContainingOperation{
		op:     operator.NotIn,
		params: vals,
	}
}

func (o ContainingOperation) nodeize() (token.Tokenizer, []interface{}) {
	t1, v1 := o.column.nodeize()
	t2, v2 := o.params.nodeize()
	return token.ConcatTokenizers(
		t1,
		t2,
		token.NewLine(
			token.Space,
			token.Word(o.operator()),
			token.Space,
		),
	), append(v1, v2...)
}

func (o ContainingOperation) operator() operator.Operator {
	return o.op
}

type NullOperation struct {
	op     operator.Operator
	column ColumnOrSubquery
}

func NewIsNull() NullOperation {
	return NullOperation{
		op: operator.Is,
	}
}

func NewIsNotNull() NullOperation {
	return NullOperation{
		op: operator.IsNot,
	}
}

func (o NullOperation) nodeize() (token.Tokenizer, []interface{}) {
	t1, v1 := o.column.nodeize()
	return token.ConcatTokenizers(
		t1,
		token.NewLine(token.Word(keyword.Null)),
		token.NewLine(
			token.Space,
			token.Word(o.operator()),
			token.Space,
		),
	), v1
}

func (o NullOperation) operator() operator.Operator {
	return o.op
}
