package sqlabble

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/operator"
)

type joinOperation struct {
	op  operator.Operator
	ops []comparisonOrLogicalOperation
}

func newAnd(ops ...comparisonOrLogicalOperation) joinOperation {
	return joinOperation{
		op:  operator.And,
		ops: ops,
	}
}

func newOr(ops ...comparisonOrLogicalOperation) joinOperation {
	return joinOperation{
		op:  operator.Or,
		ops: ops,
	}
}

func (a joinOperation) node() generator.Node {
	ns := make([]generator.Node, len(a.ops))
	for i, op := range a.ops {
		ns[i] = op.node()
	}
	return generator.NewOperator(
		a.operator(),
		ns...,
	)
}

func (a joinOperation) operator() operator.Operator {
	return a.op
}

func (a joinOperation) operations() []comparisonOrLogicalOperation {
	return a.ops
}

type not struct {
	operation comparisonOrLogicalOperation
}

func newNot(operation comparisonOrLogicalOperation) not {
	return not{operation: operation}
}

func (n not) node() generator.Node {
	return generator.NewNot(n.operation.node())
}

func (n not) operator() operator.Operator {
	return operator.Not
}

func (n not) operations() []comparisonOrLogicalOperation {
	return []comparisonOrLogicalOperation{n.operation}
}

type comparisonOperation struct {
	op  operator.Operator
	col column
	val interface{}
}

func newEq(col column, val interface{}) comparisonOperation {
	return comparisonOperation{
		op:  operator.Eq,
		col: col,
		val: val,
	}
}

func newNotEq(col column, val interface{}) comparisonOperation {
	return comparisonOperation{
		op:  operator.NotEq,
		col: col,
		val: val,
	}
}

func newGt(col column, val interface{}) comparisonOperation {
	return comparisonOperation{
		op:  operator.Gt,
		col: col,
		val: val,
	}
}

func newGte(col column, val interface{}) comparisonOperation {
	return comparisonOperation{
		op:  operator.Gte,
		col: col,
		val: val,
	}
}

func newLt(col column, val interface{}) comparisonOperation {
	return comparisonOperation{
		op:  operator.Lt,
		col: col,
		val: val,
	}
}

func newLte(col column, val interface{}) comparisonOperation {
	return comparisonOperation{
		op:  operator.Lte,
		col: col,
		val: val,
	}
}

func newLike(col column, val interface{}) comparisonOperation {
	return comparisonOperation{
		op:  operator.Like,
		col: col,
		val: val,
	}
}

func newRegExp(col column, val interface{}) comparisonOperation {
	return comparisonOperation{
		op:  operator.RegExp,
		col: col,
		val: val,
	}
}

func (c comparisonOperation) node() generator.Node {
	return generator.JoinExpressions(
		generator.NewExpression(c.col.name),
		generator.NewExpression(string(c.operator())),
		generator.ValuesToExpression(c.val),
	)
}

func (c comparisonOperation) operator() operator.Operator {
	return c.op
}

type between struct {
	col      column
	from, to interface{}
}

func newBetween(col column, from, to interface{}) between {
	return between{
		col:  col,
		from: from,
		to:   to,
	}
}

func (b between) node() generator.Node {
	return generator.JoinExpressions(
		generator.NewExpression(b.col.name),
		generator.NewExpression(string(b.operator())),
		generator.ValuesToExpression(b.from),
		generator.NewExpression(string(operator.And)),
		generator.ValuesToExpression(b.to),
	)
}

func (b between) operator() operator.Operator {
	return operator.Between
}

type containingOperation struct {
	op   operator.Operator
	col  column
	vals []interface{}
}

func newIn(col column, vals ...interface{}) containingOperation {
	return containingOperation{
		op:   operator.In,
		col:  col,
		vals: vals,
	}
}

func newNotIn(col column, vals ...interface{}) containingOperation {
	return containingOperation{
		op:   operator.NotIn,
		col:  col,
		vals: vals,
	}
}

func (o containingOperation) node() generator.Node {
	return generator.JoinExpressions(
		o.col.expression(),
		generator.NewExpression(string(o.operator())),
		generator.ValuesToExpression(o.vals...).
			Wrap("(", ")"),
	)
}

func (o containingOperation) operator() operator.Operator {
	return o.op
}

type nullyOperation struct {
	op  operator.Operator
	col column
}

func newIsNull(col column) nullyOperation {
	return nullyOperation{
		op:  operator.IsNull,
		col: col,
	}
}

func newIsNotNull(col column) nullyOperation {
	return nullyOperation{
		op:  operator.IsNotNull,
		col: col,
	}
}

func (o nullyOperation) node() generator.Node {
	return generator.JoinExpressions(
		o.col.expression(),
		generator.NewExpression(string(o.operator())),
	)
}

func (o nullyOperation) operator() operator.Operator {
	return o.op
}
