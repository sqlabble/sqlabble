package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type NonScalarOperation struct {
	op     keyword.Operator
	column ValOrColOrFuncOrSub
	params Subquery
}

func NewEqAll(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     keyword.EqAll,
		params: params,
	}
}

func NewNotEqAll(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     keyword.NotEqAll,
		params: params,
	}
}

func NewGtAll(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     keyword.GtAll,
		params: params,
	}
}

func NewGteAll(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     keyword.GteAll,
		params: params,
	}
}

func NewLtAll(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     keyword.LtAll,
		params: params,
	}
}

func NewLteAll(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     keyword.LteAll,
		params: params,
	}
}

func NewEqAny(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     keyword.EqAny,
		params: params,
	}
}

func NewNotEqAny(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     keyword.NotEqAny,
		params: params,
	}
}

func NewGtAny(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     keyword.GtAny,
		params: params,
	}
}

func NewGteAny(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     keyword.GteAny,
		params: params,
	}
}

func NewLtAny(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     keyword.LtAny,
		params: params,
	}
}

func NewLteAny(params Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:     keyword.LteAny,
		params: params,
	}
}

func (n NonScalarOperation) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t2, v2 := n.params.nodeize()
	if n.column == nil {
		return t2.Prepend(
			token.Word(n.keyword()),
			token.Space,
		), v2
	}

	t1, v1 := n.column.nodeize()
	return tokenizer.ConcatTokenizers(
		t1,
		t2,
		tokenizer.NewLine(
			token.Space,
			token.Word(n.keyword()),
			token.Space,
		),
	), append(v1, v2...)
}

func (n NonScalarOperation) keyword() keyword.Operator {
	return n.op
}

type ExistanceOperation struct {
	op     keyword.Operator
	params Subquery
}

func NewExists(params Subquery) ExistanceOperation {
	return ExistanceOperation{
		op:     keyword.Exists,
		params: params,
	}
}

func NewNotExists(params Subquery) ExistanceOperation {
	return ExistanceOperation{
		op:     keyword.NotExists,
		params: params,
	}
}

func (e ExistanceOperation) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t2, v2 := e.params.nodeize()
	return t2.Prepend(
		token.Word(e.keyword()),
		token.Space,
	), v2
}

func (e ExistanceOperation) keyword() keyword.Operator {
	return e.op
}
