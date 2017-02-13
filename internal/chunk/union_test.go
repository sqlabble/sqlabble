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

func TestUnionType(t *testing.T) {
	if c, ok := interface{}(chunk.Union{}).(grammar.Clause); !ok {
		t.Errorf("%T should implement grammar.Clause", c)
	}
}

func TestUnionSQL(t *testing.T) {
	for i, c := range []struct {
		statement grammar.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			chunk.NewUnion(
				chunk.NewSelect(
					chunk.NewColumn("emp_id"),
				).From(
					chunk.NewTable("employee"),
				).Where(
					chunk.NewAnd(
						chunk.NewColumn("assigned_branch_id").Eq(2),
						chunk.NewOr(
							chunk.NewColumn("title").Eq("Teller"),
							chunk.NewColumn("title").Eq("Head Teller"),
						),
					),
				),
				chunk.NewSelectDistinct(
					chunk.NewColumn("open_emp_id"),
				).From(
					chunk.NewTable("account"),
				).Where(
					chunk.NewColumn("open_branch_id").Eq(2),
				),
			),
			"(SELECT emp_id FROM employee WHERE assigned_branch_id = ? AND (title = ? OR title = ?)) UNION (SELECT DISTINCT open_emp_id FROM account WHERE open_branch_id = ?)",
			`> (
>   SELECT
>     emp_id
>   FROM
>     employee
>   WHERE
>     assigned_branch_id = ?
>     AND (
>       title = ?
>       OR title = ?
>     )
> )
> UNION
> (
>   SELECT DISTINCT
>     open_emp_id
>   FROM
>     account
>   WHERE
>     open_branch_id = ?
> )
`,
			[]interface{}{
				2,
				"Teller",
				"Head Teller",
				2,
			},
		},
		{
			chunk.NewUnion(
				chunk.NewUnion(
					chunk.NewSelect(chunk.NewColumn("a")),
					chunk.NewSelect(chunk.NewColumn("b")),
				),
				chunk.NewSelect(chunk.NewColumn("c")),
			),
			"((SELECT a) UNION (SELECT b)) UNION (SELECT c)",
			`> (
>   (
>     SELECT
>       a
>   )
>   UNION
>   (
>     SELECT
>       b
>   )
> )
> UNION
> (
>   SELECT
>     c
> )
`,
			[]interface{}{},
		},
		{
			chunk.NewUnion(
				chunk.NewSelect(chunk.NewColumn("a")),
				chunk.NewUnion(
					chunk.NewSelect(chunk.NewColumn("b")),
					chunk.NewSelect(chunk.NewColumn("c")),
				),
			),
			"(SELECT a) UNION ((SELECT b) UNION (SELECT c))",
			`> (
>   SELECT
>     a
> )
> UNION
> (
>   (
>     SELECT
>       b
>   )
>   UNION
>   (
>     SELECT
>       c
>   )
> )
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

func TestUnionSQLWithMySQL4(t *testing.T) {
	for i, c := range []struct {
		statement grammar.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			chunk.NewUnion(
				chunk.NewSelect(
					chunk.NewColumn("emp_id"),
				).From(
					chunk.NewTable("employee"),
				).Where(
					chunk.NewAnd(
						chunk.NewColumn("assigned_branch_id").Eq(2),
						chunk.NewOr(
							chunk.NewColumn("title").Eq("Teller"),
							chunk.NewColumn("title").Eq("Head Teller"),
						),
					),
				),
				chunk.NewSelectDistinct(
					chunk.NewColumn("open_emp_id"),
				).From(
					chunk.NewTable("account"),
				).Where(
					chunk.NewColumn("open_branch_id").Eq(2),
				),
			),
			"(SELECT emp_id FROM employee WHERE assigned_branch_id = ? AND (title = ? OR title = ?)) UNION (SELECT DISTINCT open_emp_id FROM account WHERE open_branch_id = ?)",
			`> (
>   SELECT
>     emp_id
>   FROM
>     employee
>   WHERE
>     assigned_branch_id = ?
>     AND (
>       title = ?
>       OR title = ?
>     )
> )
> UNION
> (
>   SELECT DISTINCT
>     open_emp_id
>   FROM
>     account
>   WHERE
>     open_branch_id = ?
> )
`,
			[]interface{}{
				2,
				"Teller",
				"Head Teller",
				2,
			},
		},
		{
			chunk.NewUnion(
				chunk.NewUnion(
					chunk.NewSelect(chunk.NewColumn("a")),
					chunk.NewSelect(chunk.NewColumn("b")),
				),
				chunk.NewSelect(chunk.NewColumn("c")),
			),
			"(SELECT a) UNION (SELECT b) UNION (SELECT c)",
			`> (
>   SELECT
>     a
> )
> UNION
> (
>   SELECT
>     b
> )
> UNION
> (
>   SELECT
>     c
> )
`,
			[]interface{}{},
		},
		{
			chunk.NewUnion(
				chunk.NewSelect(chunk.NewColumn("a")),
				chunk.NewUnion(
					chunk.NewSelect(chunk.NewColumn("b")),
					chunk.NewSelect(chunk.NewColumn("c")),
				),
			),
			"(SELECT a) UNION (SELECT b) UNION (SELECT c)",
			`> (
>   SELECT
>     a
> )
> UNION
> (
>   SELECT
>     b
> )
> UNION
> (
>   SELECT
>     c
> )
`,
			[]interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			b := sqlabble.NewBuilderForMySQL4("", "")
			sql, values := b.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			b := sqlabble.NewBuilderForMySQL4("> ", "  ")
			sql, values := b.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}
