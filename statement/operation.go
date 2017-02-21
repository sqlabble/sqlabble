package statement

import (
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
		var vals []interface{}
		ts[i], vals = op.nodeize()
		values = append(values, vals...)
	}
	if len(values) == 0 {
		values = nil
	}
	return ts.Prefix(
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
	return token.NewContainer(
		token.NewLine(
			token.Word(o.operator()),
			token.Space,
			token.ParenthesesStart,
		),
	).SetMiddle(
		middle,
	).SetLast(
		token.NewLine(
			token.ParenthesesEnd,
		),
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
	values := append(v1, v2...)
	return token.ConcatTokenizers(
		t1,
		t2,
		token.NewLine(
			token.Space,
			token.Word(o.operator()),
			token.Space,
		),
	), values
}

func (o ComparisonOperation) operator() operator.Operator {
	return o.op
}

type Between struct {
	col      ColumnOrSubquery
	from, to interface{}
}

func NewBetween(from, to interface{}) Between {
	return Between{
		from: from,
		to:   to,
	}
}

func (o Between) nodeize() (token.Tokenizer, []interface{}) {
	line := token.NewLine(
		token.Word(o.operator()),
		token.Space,
		token.PlaceholderTokens(1),
		token.Space,
		token.Word(operator.And),
		token.Space,
		token.PlaceholderTokens(1),
	)
	values := []interface{}{o.from, o.to}

	if o.col == nil {
		return line, values
	}

	return line, values
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
	return token.ConcatTokenizers(t1, t2, token.NewLine(
		token.Space,
		token.Word(o.operator()),
		token.Space,
	)), append(v1, v2...)
}

func (o ContainingOperation) operator() operator.Operator {
	return o.op
}

type NullOperation struct {
	op  operator.Operator
	col ColumnOrSubquery
}

func NewIsNull() NullOperation {
	return NullOperation{
		op: operator.IsNull,
	}
}

func NewIsNotNull() NullOperation {
	return NullOperation{
		op: operator.IsNotNull,
	}
}

func (o NullOperation) nodeize() (token.Tokenizer, []interface{}) {
	return token.NewLine(), nil
}

func (o NullOperation) operator() operator.Operator {
	return o.op
}
