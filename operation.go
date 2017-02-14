package sqlabble

import (
	"fmt"

	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/operator"
)

type and struct {
	ops []operationNode
}

func newAnd(cs ...operationNode) and {
	return and{ops: cs}
}

func (a and) generator() generator.Generator {
	fs := make([]generator.Generator, len(a.ops))
	for i, o := range a.ops {
		fs[i] = o.generator()
	}
	return generator.NewOperator(a.operator(), fs...)
}

func (a and) operator() operator.Operator {
	return operator.And
}

func (a and) operations() []operationNode {
	return a.ops
}

type or struct {
	ops []operationNode
}

func newOr(conditions ...operationNode) or {
	return or{ops: conditions}
}

func (a or) generator() generator.Generator {
	fs := make([]generator.Generator, len(a.ops))
	for i, o := range a.ops {
		fs[i] = o.generator()
	}
	return generator.NewOperator(a.operator(), fs...)
}

func (o or) operator() operator.Operator {
	return operator.Or
}

func (o or) operations() []operationNode {
	return o.ops
}

type not struct {
	operation operationNode
}

func newNot(operation operationNode) not {
	return not{operation: operation}
}

func (n not) generator() generator.Generator {
	return generator.NewNot(n.operation.generator())
}

func (n not) operator() operator.Operator {
	return operator.Not
}

func (n not) operations() []operationNode {
	return []operationNode{n.operation}
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

func (c eq) generator() generator.Generator {
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

func (c notEq) generator() generator.Generator {
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

func (c gt) generator() generator.Generator {
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

func (c gte) generator() generator.Generator {
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

func (c lt) generator() generator.Generator {
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

func (c lte) generator() generator.Generator {
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

func (c between) generator() generator.Generator {
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

func (c in) generator() generator.Generator {
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

func (c notIn) generator() generator.Generator {
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

func (l like) generator() generator.Generator {
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

func (l regExp) generator() generator.Generator {
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

func (l isNull) generator() generator.Generator {
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

func (l isNotNull) generator() generator.Generator {
	return generator.NewExpression(
		fmt.Sprintf("%s %s", l.col.name, l.Operator()),
	)
}

func (l isNotNull) Operator() operator.Operator {
	return operator.IsNotNull
}
