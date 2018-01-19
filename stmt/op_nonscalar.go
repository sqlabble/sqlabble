package stmt

import (
	"github.com/sqlabble/sqlabble/keyword"
	"github.com/sqlabble/sqlabble/token"
	"github.com/sqlabble/sqlabble/tokenizer"
)

type NonScalarOperation struct {
	op     keyword.Operator
	column ValOrColOrFuncOrSub
	sub    Subquery
}

func NewEqAll(sub Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:  keyword.EqAll,
		sub: sub,
	}
}

func NewNotEqAll(sub Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:  keyword.NotEqAll,
		sub: sub,
	}
}

func NewGtAll(sub Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:  keyword.GtAll,
		sub: sub,
	}
}

func NewGteAll(sub Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:  keyword.GteAll,
		sub: sub,
	}
}

func NewLtAll(sub Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:  keyword.LtAll,
		sub: sub,
	}
}

func NewLteAll(sub Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:  keyword.LteAll,
		sub: sub,
	}
}

func NewEqAny(sub Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:  keyword.EqAny,
		sub: sub,
	}
}

func NewNotEqAny(sub Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:  keyword.NotEqAny,
		sub: sub,
	}
}

func NewGtAny(sub Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:  keyword.GtAny,
		sub: sub,
	}
}

func NewGteAny(sub Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:  keyword.GteAny,
		sub: sub,
	}
}

func NewLtAny(sub Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:  keyword.LtAny,
		sub: sub,
	}
}

func NewLteAny(sub Subquery) NonScalarOperation {
	return NonScalarOperation{
		op:  keyword.LteAny,
		sub: sub,
	}
}

func (n NonScalarOperation) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t2, v2 := n.sub.nodeize()
	if n.column == nil {
		return t2.Prepend(
			token.Word(n.keyword()),
		), v2
	}

	t1, v1 := n.column.nodeize()
	return tokenizer.ConcatTokenizers(
		t1,
		t2,
		tokenizer.NewLine(
			token.Word(n.keyword()),
		),
	), append(v1, v2...)
}

func (n NonScalarOperation) keyword() keyword.Operator {
	return n.op
}

type ExistanceOperation struct {
	op  keyword.Operator
	sub Subquery
}

func NewExists(sub Subquery) ExistanceOperation {
	return ExistanceOperation{
		op:  keyword.Exists,
		sub: sub,
	}
}

func NewNotExists(sub Subquery) ExistanceOperation {
	return ExistanceOperation{
		op:  keyword.NotExists,
		sub: sub,
	}
}

func (e ExistanceOperation) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t2, v2 := e.sub.nodeize()
	return t2.Prepend(
		token.Word(e.keyword()),
	), v2
}

func (e ExistanceOperation) keyword() keyword.Operator {
	return e.op
}
