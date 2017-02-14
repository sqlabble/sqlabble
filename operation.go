package sqlabble

import (
	"fmt"

	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/operator"
)

type and struct {
	ops []comparisonOrLogicalOperation
}

func newAnd(cs ...comparisonOrLogicalOperation) and {
	return and{ops: cs}
}

func (a and) node() generator.Node {
	fs := make([]generator.Node, len(a.ops))
	for i, o := range a.ops {
		fs[i] = o.node()
	}
	return generator.NewOperator(a.operator(), fs...)
}

func (a and) operator() operator.Operator {
	return operator.And
}

func (a and) operations() []comparisonOrLogicalOperation {
	return a.ops
}

type or struct {
	ops []comparisonOrLogicalOperation
}

func newOr(conditions ...comparisonOrLogicalOperation) or {
	return or{ops: conditions}
}

func (a or) node() generator.Node {
	fs := make([]generator.Node, len(a.ops))
	for i, o := range a.ops {
		fs[i] = o.node()
	}
	return generator.NewOperator(a.operator(), fs...)
}

func (o or) operator() operator.Operator {
	return operator.Or
}

func (o or) operations() []comparisonOrLogicalOperation {
	return o.ops
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

type eq struct {
	col   column
	value interface{}
}

func newEq(col column, value interface{}) eq {
	return eq{
		col:   col,
		value: value,
	}
}

func (c eq) node() generator.Node {
	return generator.NewExpression(
		fmt.Sprintf("%s %s ?", c.col.name, c.operator()),
		c.value,
	)
}

func (c eq) operator() operator.Operator {
	return operator.Equal
}

type notEq struct {
	col   column
	value interface{}
}

func newNotEq(col column, value interface{}) notEq {
	return notEq{
		col:   col,
		value: value,
	}
}

func (c notEq) node() generator.Node {
	return generator.NewExpression(
		fmt.Sprintf("%s %s ?", c.col.name, c.operator()),
		c.value,
	)
}

func (c notEq) operator() operator.Operator {
	return operator.NotEqual
}

type gt struct {
	col   column
	value interface{}
}

func newGt(col column, value interface{}) gt {
	return gt{
		col:   col,
		value: value,
	}
}

func (c gt) node() generator.Node {
	return generator.NewExpression(
		fmt.Sprintf("%s %s ?", c.col.name, c.operator()),
		c.value,
	)
}

func (c gt) operator() operator.Operator {
	return operator.Gt
}

type gte struct {
	col   column
	value interface{}
}

func newGte(col column, value interface{}) gte {
	return gte{
		col:   col,
		value: value,
	}
}

func (c gte) node() generator.Node {
	return generator.NewExpression(
		fmt.Sprintf("%s %s ?", c.col.name, c.operator()),
		c.value,
	)
}

func (c gte) operator() operator.Operator {
	return operator.Gte
}

type lt struct {
	col   column
	value interface{}
}

func newLt(col column, value interface{}) lt {
	return lt{
		col:   col,
		value: value,
	}
}

func (c lt) node() generator.Node {
	return generator.NewExpression(
		fmt.Sprintf("%s %s ?", c.col.name, c.operator()),
		c.value,
	)
}

func (c lt) operator() operator.Operator {
	return operator.Lt
}

type lte struct {
	col   column
	value interface{}
}

func newLte(col column, value interface{}) lte {
	return lte{
		col:   col,
		value: value,
	}
}

func (c lte) node() generator.Node {
	return generator.NewExpression(
		fmt.Sprintf("%s %s ?", c.col.name, c.operator()),
		c.value,
	)
}

func (c lte) operator() operator.Operator {
	return operator.Lte
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

func (c between) node() generator.Node {
	return generator.NewExpression(
		fmt.Sprintf("%s %s ? %s ?", c.col.name, c.operator(), operator.And),
		c.from, c.to,
	)
}

func (c between) operator() operator.Operator {
	return operator.Between
}

type in struct {
	col    column
	values []interface{}
}

func newIn(col column, values ...interface{}) in {
	return in{
		col:    col,
		values: values,
	}
}

func (c in) node() generator.Node {
	return generator.NewExpression(
		fmt.Sprintf(
			"%s %s (%s)",
			c.col.name,
			c.operator(),
			placeholders(len(c.values)),
		),
		c.values...,
	)
}

func (c in) operator() operator.Operator {
	return operator.In
}

type notIn struct {
	col    column
	values []interface{}
}

func newNotIn(col column, values ...interface{}) notIn {
	return notIn{
		col:    col,
		values: values,
	}
}

func (c notIn) node() generator.Node {
	return generator.NewExpression(
		fmt.Sprintf(
			"%s %s (%s)",
			c.col.name,
			c.operator(),
			placeholders(len(c.values)),
		),
		c.values...,
	)
}

func (c notIn) operator() operator.Operator {
	return operator.NotIn
}

type like struct {
	col column
	lik string
}

func newLike(col column, lik string) like {
	return like{
		col: col,
		lik: lik,
	}
}

func (l like) node() generator.Node {
	return generator.NewExpression(
		fmt.Sprintf("%s %s ?", l.col.name, l.operator()),
		l.lik,
	)
}

func (l like) operator() operator.Operator {
	return operator.Like
}

type regExp struct {
	col    column
	regexp string
}

func newRegExp(col column, regexp string) regExp {
	return regExp{
		col:    col,
		regexp: regexp,
	}
}

func (l regExp) node() generator.Node {
	return generator.NewExpression(
		fmt.Sprintf("%s %s ?", l.col.name, l.operator()),
		l.regexp,
	)
}

func (l regExp) operator() operator.Operator {
	return operator.RegExp
}

type isNull struct {
	col column
}

func newIsNull(col column) isNull {
	return isNull{
		col: col,
	}
}

func (l isNull) node() generator.Node {
	return generator.NewExpression(
		fmt.Sprintf("%s %s", l.col.name, l.Operator()),
	)
}

func (l isNull) Operator() operator.Operator {
	return operator.IsNull
}

type isNotNull struct {
	col column
}

func newIsNotNull(col column) isNotNull {
	return isNotNull{
		col: col,
	}
}

func (l isNotNull) node() generator.Node {
	return generator.NewExpression(
		fmt.Sprintf("%s %s", l.col.name, l.Operator()),
	)
}

func (l isNotNull) Operator() operator.Operator {
	return operator.IsNotNull
}
