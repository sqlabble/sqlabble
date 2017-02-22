package statement

import (
	"github.com/minodisk/sqlabble/operator"
	"github.com/minodisk/sqlabble/token"
)

type NonScalarOperation struct {
	op     operator.Operator
	column ColumnOrSubquery
	params Subquery
}

func NewEqAll(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     operator.EqAll,
		params: params,
	}
}

func NewNotEqAll(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     operator.NotEqAll,
		params: params,
	}
}

func NewGtAll(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     operator.GtAll,
		params: params,
	}
}

func NewGteAll(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     operator.GteAll,
		params: params,
	}
}

func NewLtAll(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     operator.LtAll,
		params: params,
	}
}

func NewLteAll(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     operator.LteAll,
		params: params,
	}
}

func NewEqAny(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     operator.EqAny,
		params: params,
	}
}

func NewNotEqAny(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     operator.NotEqAny,
		params: params,
	}
}

func NewGtAny(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     operator.GtAny,
		params: params,
	}
}

func NewGteAny(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     operator.GteAny,
		params: params,
	}
}

func NewLtAny(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     operator.LtAny,
		params: params,
	}
}

func NewLteAny(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     operator.LteAny,
		params: params,
	}
}

func (n NonScalarOperation) nodeize() (token.Tokenizer, []interface{}) {
	t2, v2 := n.params.nodeize()
	if n.column == nil {
		return t2.Prepend(
			token.Word(n.operator()),
			token.Space,
		), v2
	}

	t1, v1 := n.column.nodeize()
	return token.ConcatTokenizers(
		t1,
		t2,
		token.NewLine(
			token.Space,
			token.Word(n.operator()),
			token.Space,
		),
	), append(v1, v2...)
}

func (n NonScalarOperation) operator() operator.Operator {
	return n.op
}

type ExistanceOperation struct {
	op     operator.Operator
	params Subquery
}

func NewExists(params Subquery) ExistanceOperation {
	return ExistanceOperation{
		op:     operator.Exists,
		params: params,
	}
}

func NewNotExists(params Subquery) ExistanceOperation {
	return ExistanceOperation{
		op:     operator.NotExists,
		params: params,
	}
}

func (e ExistanceOperation) nodeize() (token.Tokenizer, []interface{}) {
	t2, v2 := e.params.nodeize()
	return t2.Prepend(
		token.Word(e.operator()),
		token.Space,
	), v2
}

func (e ExistanceOperation) operator() operator.Operator {
	return e.op
}
