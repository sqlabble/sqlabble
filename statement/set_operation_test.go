package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestUnionSQL(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewUnion(
				statement.NewSelect(statement.NewColumn("a")).From(statement.NewTable("aaa")),
				statement.NewSelect(statement.NewColumn("b")),
			).OrderBy(
				statement.NewColumn("foo").Asc(),
			),
			"(SELECT a FROM aaa) UNION (SELECT b) ORDER BY foo ASC",
			`> (
>   SELECT
>     a
>   FROM
>     aaa
> )
> UNION (
>   SELECT
>     b
> )
> ORDER BY
>   foo ASC
`,
			nil,
		},
		{
			statement.NewUnion(
				statement.NewSelect(statement.NewColumn("a")),
				statement.NewUnion(
					statement.NewUnion(
						statement.NewSelect(statement.NewColumn("b")),
						statement.NewSelect(statement.NewColumn("c")),
					),
					statement.NewSelect(statement.NewColumn("d")),
					statement.NewSelect(statement.NewColumn("e")),
				),
			),
			"(SELECT a) UNION (((SELECT b) UNION (SELECT c)) UNION (SELECT d) UNION (SELECT e))",
			`> (
>   SELECT
>     a
> )
> UNION (
>   (
>     (
>       SELECT
>         b
>     )
>     UNION (
>       SELECT
>         c
>     )
>   )
>   UNION (
>     SELECT
>       d
>   )
>   UNION (
>     SELECT
>       e
>   )
> )
`,
			nil,
		},
	} {
		t.Run(fmt.Sprintf("%d Builder", i), func(t *testing.T) {
			sql, values := b.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := bi.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

func TestUnionAllSQL(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewUnionAll(
				statement.NewSelect(statement.NewColumn("a")),
				statement.NewUnionAll(
					statement.NewUnionAll(
						statement.NewSelect(statement.NewColumn("b")),
						statement.NewSelect(statement.NewColumn("c")),
					),
					statement.NewSelect(statement.NewColumn("d")),
					statement.NewSelect(statement.NewColumn("e")),
				),
			),
			"(SELECT a) UNION ALL (((SELECT b) UNION ALL (SELECT c)) UNION ALL (SELECT d) UNION ALL (SELECT e))",
			`> (
>   SELECT
>     a
> )
> UNION ALL (
>   (
>     (
>       SELECT
>         b
>     )
>     UNION ALL (
>       SELECT
>         c
>     )
>   )
>   UNION ALL (
>     SELECT
>       d
>   )
>   UNION ALL (
>     SELECT
>       e
>   )
> )
`,
			nil,
		},
	} {
		t.Run(fmt.Sprintf("%d Builder", i), func(t *testing.T) {
			sql, values := b.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := bi.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

func TestIntersectSQL(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewIntersect(
				statement.NewSelect(statement.NewColumn("a")),
				statement.NewIntersect(
					statement.NewIntersect(
						statement.NewSelect(statement.NewColumn("b")),
						statement.NewSelect(statement.NewColumn("c")),
					),
					statement.NewSelect(statement.NewColumn("d")),
					statement.NewSelect(statement.NewColumn("e")),
				),
			),
			"(SELECT a) INTERSECT (((SELECT b) INTERSECT (SELECT c)) INTERSECT (SELECT d) INTERSECT (SELECT e))",
			`> (
>   SELECT
>     a
> )
> INTERSECT (
>   (
>     (
>       SELECT
>         b
>     )
>     INTERSECT (
>       SELECT
>         c
>     )
>   )
>   INTERSECT (
>     SELECT
>       d
>   )
>   INTERSECT (
>     SELECT
>       e
>   )
> )
`,
			nil,
		},
	} {
		t.Run(fmt.Sprintf("%d Builder", i), func(t *testing.T) {
			sql, values := b.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := bi.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

func TestIntersectAllSQL(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewIntersectAll(
				statement.NewSelect(statement.NewColumn("a")),
				statement.NewIntersectAll(
					statement.NewIntersectAll(
						statement.NewSelect(statement.NewColumn("b")),
						statement.NewSelect(statement.NewColumn("c")),
					),
					statement.NewSelect(statement.NewColumn("d")),
					statement.NewSelect(statement.NewColumn("e")),
				),
			),
			"(SELECT a) INTERSECT ALL (((SELECT b) INTERSECT ALL (SELECT c)) INTERSECT ALL (SELECT d) INTERSECT ALL (SELECT e))",
			`> (
>   SELECT
>     a
> )
> INTERSECT ALL (
>   (
>     (
>       SELECT
>         b
>     )
>     INTERSECT ALL (
>       SELECT
>         c
>     )
>   )
>   INTERSECT ALL (
>     SELECT
>       d
>   )
>   INTERSECT ALL (
>     SELECT
>       e
>   )
> )
`,
			nil,
		},
	} {
		t.Run(fmt.Sprintf("%d Builder", i), func(t *testing.T) {
			sql, values := b.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := bi.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

func TestExceptSQL(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewExcept(
				statement.NewSelect(statement.NewColumn("a")),
				statement.NewExcept(
					statement.NewExcept(
						statement.NewSelect(statement.NewColumn("b")),
						statement.NewSelect(statement.NewColumn("c")),
					),
					statement.NewSelect(statement.NewColumn("d")),
					statement.NewSelect(statement.NewColumn("e")),
				),
			),
			"(SELECT a) EXCEPT (((SELECT b) EXCEPT (SELECT c)) EXCEPT (SELECT d) EXCEPT (SELECT e))",
			`> (
>   SELECT
>     a
> )
> EXCEPT (
>   (
>     (
>       SELECT
>         b
>     )
>     EXCEPT (
>       SELECT
>         c
>     )
>   )
>   EXCEPT (
>     SELECT
>       d
>   )
>   EXCEPT (
>     SELECT
>       e
>   )
> )
`,
			nil,
		},
	} {
		t.Run(fmt.Sprintf("%d Builder", i), func(t *testing.T) {
			sql, values := b.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := bi.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

func TestExceptAllSQL(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewExceptAll(
				statement.NewSelect(statement.NewColumn("a")),
				statement.NewExceptAll(
					statement.NewExceptAll(
						statement.NewSelect(statement.NewColumn("b")),
						statement.NewSelect(statement.NewColumn("c")),
					),
					statement.NewSelect(statement.NewColumn("d")),
					statement.NewSelect(statement.NewColumn("e")),
				),
			),
			"(SELECT a) EXCEPT ALL (((SELECT b) EXCEPT ALL (SELECT c)) EXCEPT ALL (SELECT d) EXCEPT ALL (SELECT e))",
			`> (
>   SELECT
>     a
> )
> EXCEPT ALL (
>   (
>     (
>       SELECT
>         b
>     )
>     EXCEPT ALL (
>       SELECT
>         c
>     )
>   )
>   EXCEPT ALL (
>     SELECT
>       d
>   )
>   EXCEPT ALL (
>     SELECT
>       e
>   )
> )
`,
			nil,
		},
	} {
		t.Run(fmt.Sprintf("%d Builder", i), func(t *testing.T) {
			sql, values := b.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := bi.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}
