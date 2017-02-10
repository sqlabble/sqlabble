package chunk_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/chunk"
	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/internal/grammar"
)

func TestOrderByClauseIsCaluse(t *testing.T) {
	if _, ok := interface{}(chunk.OrderBy{}).(grammar.Clause); !ok {
		t.Errorf("chunk.OrderByClause doesn't implement grammar.Clause")
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
			chunk.NewOrderBy(
				chunk.NewColumn("foo").Asc(),
			),
			"ORDER BY foo ASC",
			`> ORDER BY
>   foo ASC
`,
			[]interface{}{},
		},
		{
			chunk.NewOrderBy(
				chunk.NewColumn("foo").Desc(),
				chunk.NewColumn("bar").Asc(),
				chunk.NewColumn("baz").Desc(),
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
