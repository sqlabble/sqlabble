package token_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/token"
)

func TestWhereClauseIsCaluse(t *testing.T) {
	if _, ok := interface{}(token.Where{}).(grammar.Clause); !ok {
		t.Errorf("token.WhereClause doesn't implement grammar.Clause")
	}
}

func TestWhereClause(t *testing.T) {
	for i, c := range []struct {
		statement token.Where
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			token.NewWhere(
				token.NewColumn("foo").Eq(100),
			),
			"WHERE foo = ?",
			`> WHERE
>   foo = ?
`,
			[]interface{}{100},
		},
		{
			token.NewWhere(
				token.NewAnd(
					token.NewColumn("foo").Eq(100),
					token.NewColumn("bar").Eq("abc"),
				),
			),
			"WHERE foo = ? AND bar = ?",
			`> WHERE
>   foo = ?
>   AND bar = ?
`,
			[]interface{}{100, "abc"},
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := sqlabble.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := sqlabble.BuildIndent(c.statement, "> ", "  ")
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}
