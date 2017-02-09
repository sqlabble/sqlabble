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

func TestOrderByClauseIsCaluse(t *testing.T) {
	if _, ok := interface{}(token.OrderBy{}).(grammar.Clause); !ok {
		t.Errorf("token.OrderByClause doesn't implement grammar.Clause")
	}
}

func TestOrderByClause(t *testing.T) {
	for i, c := range []struct {
		statement grammar.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			token.NewOrderBy(
				token.NewColumn("foo").Asc(),
			),
			"ORDER BY foo ASC",
			`> ORDER BY
>   foo ASC
`,
			[]interface{}{},
		},
		{
			token.NewOrderBy(
				token.NewColumn("foo").Desc(),
				token.NewColumn("bar").Asc(),
				token.NewColumn("baz").Desc(),
			),
			"ORDER BY foo DESC, bar ASC, baz DESC",
			`> ORDER BY
>   foo DESC
>   , bar ASC
>   , baz DESC
`,
			[]interface{}{},
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
