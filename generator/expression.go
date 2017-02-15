package generator

import (
	"fmt"
	"strings"
)

var (
	EmptyExpression = NewExpression("")
)

func JoinExpressions(es ...Expression) Expression {
	var exp Expression
	for i, e := range es {
		if i == 0 {
			exp = e
			continue
		}
		exp = exp.Append(e)
	}
	return exp
}

type Expression struct {
	sql    string
	values []interface{}
}

func NewExpression(sql string, values ...interface{}) Expression {
	if len(values) == 0 {
		values = []interface{}{}
	}
	return Expression{
		sql:    sql,
		values: values,
	}
}

func (e Expression) ToSQL(ctx Context) (string, []interface{}) {
	h := ctx.Head()
	ctx = ctx.ClearHead()
	if ctx.Breaking() {
		p := ctx.Prefix()
		return fmt.Sprintf("%s%s%s\n", p, h, e.sql), e.values
	}
	return fmt.Sprintf("%s%s", h, e.sql), e.values
}

func (e Expression) Prepend(exp Expression) Expression {
	if exp.Empty() {
		return e
	}
	if e.Empty() {
		return exp
	}
	e.sql = exp.sql + " " + e.sql
	e.values = append(e.values, exp.values...)
	return e
}

func (e Expression) Append(exp Expression) Expression {
	if exp.Empty() {
		return e
	}
	if e.Empty() {
		return exp
	}
	e.sql = e.sql + " " + exp.sql
	e.values = append(e.values, exp.values...)
	return e
}

func (e Expression) WrapSQL(pre, post string) Expression {
	e.sql = pre + e.sql + post
	return e
}

func (e Expression) Empty() bool {
	return e.sql == "" && len(e.values) == 0
}

func ArrayToExpression(es ...Expression) Expression {
	l := len(es)
	if l == 0 {
		return EmptyExpression
	}
	sqls := make([]string, l)
	values := []interface{}{}
	for i, e := range es {
		sqls[i] = e.sql
		values = append(values, e.values...)
	}
	return NewExpression(
		strings.Join(sqls, ", "),
		values...,
	).WrapSQL("(", ")")
}

func ValuesToExpression(values ...interface{}) Expression {
	return NewExpression(
		placeholders(len(values)),
		values...,
	)
}
