package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/operator"
)

type joinOperation struct {
	op  operator.Operator
	ops []comparisonOrLogicalOperation
}

func NewAnd(ops ...comparisonOrLogicalOperation) joinOperation {
	return joinOperation{
		op:  operator.And,
		ops: ops,
	}
}

func NewOr(ops ...comparisonOrLogicalOperation) joinOperation {
	return joinOperation{
		op:  operator.Or,
		ops: ops,
	}
}

func (o joinOperation) node() generator.Node {
	ns := make([]generator.Node, len(o.ops))
	for i, op := range o.ops {
		ns[i] = op.node()
	}
	return generator.NewJoinOperation(
		o.operator(),
		ns...,
	)
}

func (o joinOperation) operator() operator.Operator {
	return o.op
}

func (o joinOperation) operations() []comparisonOrLogicalOperation {
	return o.ops
}

type not struct {
	operation comparisonOrLogicalOperation
}

func NewNot(operation comparisonOrLogicalOperation) not {
	return not{operation: operation}
}

func (o not) node() generator.Node {
	return generator.NewOpParenteses(
		o.operator(),
		generator.NewParentheses(o.operation.node()),
	)
}

func (o not) operator() operator.Operator {
	return operator.Not
}

func (o not) operations() []comparisonOrLogicalOperation {
	return []comparisonOrLogicalOperation{o.operation}
}

type comparisonOperation struct {
	op  operator.Operator
	col columnOrSubquery
	val interface{}
}

func NewEq(val interface{}) comparisonOperation {
	return comparisonOperation{
		op:  operator.Eq,
		val: val,
	}
}

func NewNotEq(val interface{}) comparisonOperation {
	return comparisonOperation{
		op:  operator.NotEq,
		val: val,
	}
}

func NewGt(val interface{}) comparisonOperation {
	return comparisonOperation{
		op:  operator.Gt,
		val: val,
	}
}

func NewGte(val interface{}) comparisonOperation {
	return comparisonOperation{
		op:  operator.Gte,
		val: val,
	}
}

func NewLt(val interface{}) comparisonOperation {
	return comparisonOperation{
		op:  operator.Lt,
		val: val,
	}
}

func NewLte(val interface{}) comparisonOperation {
	return comparisonOperation{
		op:  operator.Lte,
		val: val,
	}
}

func NewLike(val interface{}) comparisonOperation {
	return comparisonOperation{
		op:  operator.Like,
		val: val,
	}
}

func NewRegExp(val interface{}) comparisonOperation {
	return comparisonOperation{
		op:  operator.RegExp,
		val: val,
	}
}

func (o comparisonOperation) node() generator.Node {
	var n1, n2 generator.Node

	if o.col != nil {
		switch col := o.col.(type) {
		case column:
			n1 = col.expression()
		default:
			n1 = o.col.node()
		}
	}

	if o.val != nil {
		switch val := o.val.(type) {
		case sub:
			n2 = val.node()
		case Statement:
			n2 = NewSub(val).node()
		default:
			n2 = generator.ValuesToExpression(val)
		}
	}

	op := generator.NewExpression(string(o.operator()))
	return joinExpressionLikes(n1, n2, op)
}

func (o comparisonOperation) operator() operator.Operator {
	return o.op
}

type between struct {
	col      columnOrSubquery
	from, to interface{}
}

func NewBetween(from, to interface{}) between {
	return between{
		from: from,
		to:   to,
	}
}

func (o between) node() generator.Node {
	post := generator.JoinExpressions(
		generator.NewExpression(string(o.operator())),
		generator.ValuesToExpression(o.from),
		generator.NewExpression(string(operator.And)),
		generator.ValuesToExpression(o.to),
	)
	if o.col == nil {
		return post
	}
	switch col := o.col.(type) {
	case column:
		return generator.JoinExpressions(
			col.expression(),
			post,
		)
	default:
		return generator.NewNodes(
			o.col.node(),
			post,
		)
	}
}

func (o between) operator() operator.Operator {
	return operator.Between
}

type containingOperation struct {
	op   operator.Operator
	col  columnOrSubquery
	vals []interface{}
}

func NewIn(vals ...interface{}) containingOperation {
	return containingOperation{
		op:   operator.In,
		vals: vals,
	}
}

func NewNotIn(vals ...interface{}) containingOperation {
	return containingOperation{
		op:   operator.NotIn,
		vals: vals,
	}
}

func (o containingOperation) node() generator.Node {
	var n1, n2 generator.Node

	if o.col != nil {
		switch col := o.col.(type) {
		case column:
			n1 = col.expression()
		default:
			n1 = o.col.node()
		}
	}

	if o.vals != nil && len(o.vals) == 1 {
		switch val := o.vals[0].(type) {
		case sub:
			n2 = val.node()
		case Statement:
			n2 = NewSub(val).node()
		}
	}
	if n2 == nil {
		n2 = generator.JoinExpressions(
			generator.ValuesToExpression(o.vals...).
				WrapSQL("(", ")"),
		)

	}

	op := generator.NewExpression(string(o.operator()))
	return joinExpressionLikes(n1, n2, op)
}

func joinExpressionLikes(n1, n2 generator.Node, op generator.Expression) generator.Node {
	e1, ok1 := n1.(generator.Expression)
	e2, ok2 := n2.(generator.Expression)
	if ok1 && ok2 {
		return generator.JoinExpressions(e1, op, e2)
	}
	if ok1 {
		return generator.NewNodes(e1.Append(op), n2)
	}
	if ok2 {
		return generator.NewNodes(n1, e2.Prepend(op))
	}
	return generator.NewNodes(n1, op, n2)
}

func (o containingOperation) operator() operator.Operator {
	return o.op
}

type nullyOperation struct {
	op  operator.Operator
	col columnOrSubquery
}

func NewIsNull() nullyOperation {
	return nullyOperation{
		op: operator.IsNull,
	}
}

func NewIsNotNull() nullyOperation {
	return nullyOperation{
		op: operator.IsNotNull,
	}
}

func (o nullyOperation) node() generator.Node {
	post := generator.NewExpression(string(o.operator()))
	if o.col == nil {
		return post
	}
	switch col := o.col.(type) {
	case column:
		return generator.JoinExpressions(
			col.expression(),
			post,
		)
	default:
		return generator.NewNodes(
			o.col.node(),
			post,
		)
	}
}

func (o nullyOperation) operator() operator.Operator {
	return o.op
}
