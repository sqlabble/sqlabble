package statement

import (
	"github.com/minodisk/sqlabble/node"
	"github.com/minodisk/sqlabble/operator"
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

func (o JoinOperation) node() node.Node {
	ns := make([]node.Node, len(o.ops))
	for i, op := range o.ops {
		ns[i] = op.node()
	}
	return node.NewJoinOperation(
		o.operator(),
		ns...,
	)
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

func (o Not) node() node.Node {
	return node.NewOpParentheses(
		o.operator(),
		node.NewParentheses(o.operation.node()),
	)
}

func (o Not) operator() operator.Operator {
	return operator.Not
}

func (o Not) operations() []ComparisonOrLogicalOperation {
	return []ComparisonOrLogicalOperation{o.operation}
}

type ComparisonOperation struct {
	op  operator.Operator
	col ColumnOrSubquery
	val interface{}
}

func NewEq(val interface{}) ComparisonOperation {
	return ComparisonOperation{
		op:  operator.Eq,
		val: val,
	}
}

func NewNotEq(val interface{}) ComparisonOperation {
	return ComparisonOperation{
		op:  operator.NotEq,
		val: val,
	}
}

func NewGt(val interface{}) ComparisonOperation {
	return ComparisonOperation{
		op:  operator.Gt,
		val: val,
	}
}

func NewGte(val interface{}) ComparisonOperation {
	return ComparisonOperation{
		op:  operator.Gte,
		val: val,
	}
}

func NewLt(val interface{}) ComparisonOperation {
	return ComparisonOperation{
		op:  operator.Lt,
		val: val,
	}
}

func NewLte(val interface{}) ComparisonOperation {
	return ComparisonOperation{
		op:  operator.Lte,
		val: val,
	}
}

func NewLike(val interface{}) ComparisonOperation {
	return ComparisonOperation{
		op:  operator.Like,
		val: val,
	}
}

func NewRegExp(val interface{}) ComparisonOperation {
	return ComparisonOperation{
		op:  operator.RegExp,
		val: val,
	}
}

func (o ComparisonOperation) node() node.Node {
	var n1, n2 node.Node

	if o.col != nil {
		switch col := o.col.(type) {
		case Column:
			n1 = col.expression()
		default:
			n1 = o.col.node()
		}
	}

	if o.val != nil {
		switch val := o.val.(type) {
		case Subquery:
			n2 = val.node()
		case Statement:
			n2 = NewSubquery(val).node()
		default:
			n2 = node.ValuesToExpression(val)
		}
	}

	op := node.NewExpression(string(o.operator()))
	return joinExpressionLikes(n1, n2, op)
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

func (o Between) node() node.Node {
	post := node.JoinExpressions(
		node.NewExpression(string(o.operator())),
		node.ValuesToExpression(o.from),
		node.NewExpression(string(operator.And)),
		node.ValuesToExpression(o.to),
	)
	if o.col == nil {
		return post
	}
	switch col := o.col.(type) {
	case Column:
		return node.JoinExpressions(
			col.expression(),
			post,
		)
	default:
		return node.NewNodes(
			o.col.node(),
			post,
		)
	}
}

func (o Between) operator() operator.Operator {
	return operator.Between
}

type ContainingOperation struct {
	op   operator.Operator
	col  ColumnOrSubquery
	vals []interface{}
}

func NewIn(vals ...interface{}) ContainingOperation {
	return ContainingOperation{
		op:   operator.In,
		vals: vals,
	}
}

func NewNotIn(vals ...interface{}) ContainingOperation {
	return ContainingOperation{
		op:   operator.NotIn,
		vals: vals,
	}
}

func (o ContainingOperation) node() node.Node {
	var n1, n2 node.Node

	if o.col != nil {
		switch col := o.col.(type) {
		case Column:
			n1 = col.expression()
		default:
			n1 = o.col.node()
		}
	}

	if o.vals != nil && len(o.vals) == 1 {
		switch val := o.vals[0].(type) {
		case Subquery:
			n2 = val.node()
		case Statement:
			n2 = NewSubquery(val).node()
		}
	}
	if n2 == nil {
		n2 = node.JoinExpressions(
			node.ValuesToExpression(o.vals...).
				WrapSQL("(", ")"),
		)

	}

	op := node.NewExpression(string(o.operator()))
	return joinExpressionLikes(n1, n2, op)
}

func joinExpressionLikes(n1, n2 node.Node, op node.Expression) node.Node {
	e1, ok1 := n1.(node.Expression)
	e2, ok2 := n2.(node.Expression)
	if ok1 && ok2 {
		return node.JoinExpressions(e1, op, e2)
	}
	if ok1 {
		return node.NewNodes(e1.Append(op), n2)
	}
	if ok2 {
		return node.NewNodes(n1, e2.Prepend(op))
	}
	return node.NewNodes(n1, op, n2)
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

func (o NullOperation) node() node.Node {
	post := node.NewExpression(string(o.operator()))
	if o.col == nil {
		return post
	}
	switch col := o.col.(type) {
	case Column:
		return node.JoinExpressions(
			col.expression(),
			post,
		)
	default:
		return node.NewNodes(
			o.col.node(),
			post,
		)
	}
}

func (o NullOperation) operator() operator.Operator {
	return o.op
}
