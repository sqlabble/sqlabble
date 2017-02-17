package node

import "strings"

type Quote struct {
	sql    string
	split  bool
	values []interface{}
}

func NewQuote(sql string, split bool, values ...interface{}) Quote {
	return Quote{
		sql:    sql,
		split:  split,
		values: values,
	}
}

func (q Quote) ToSQL(ctx Context) (string, []interface{}) {
	if q.split {
		return ctx.Quote + q.sql + ctx.Quote, q.values
	}

	sqls := strings.Split(q.sql, ".")
	for i, s := range sqls {
		sqls[i] = ctx.Quote + s + ctx.Quote
	}
	return strings.Join(sqls, "."), q.values
}
