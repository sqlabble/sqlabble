package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestUnionSQL(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewUnion(
				stmt.NewSelect(stmt.NewColumn("a")),
				stmt.NewSelect(stmt.NewColumn("b")),
			).OrderBy(
				stmt.NewColumn("foo").Asc(),
			),
			`(SELECT "a") UNION (SELECT "b") ORDER BY "foo" ASC`,
			`> (
>   SELECT
>     "a"
> )
> UNION (
>   SELECT
>     "b"
> )
> ORDER BY
>   "foo" ASC
`,
			nil,
		},
		{
			stmt.NewUnion(
				stmt.NewSelect(stmt.NewColumn("a")),
				stmt.NewUnion(
					stmt.NewUnion(
						stmt.NewSelect(stmt.NewColumn("b")),
						stmt.NewSelect(stmt.NewColumn("c")),
					),
					stmt.NewSelect(stmt.NewColumn("d")),
					stmt.NewSelect(stmt.NewColumn("e")),
				),
			),
			`(SELECT "a") UNION (((SELECT "b") UNION (SELECT "c")) UNION (SELECT "d") UNION (SELECT "e"))`,
			`> (
>   SELECT
>     "a"
> )
> UNION (
>   (
>     (
>       SELECT
>         "b"
>     )
>     UNION (
>       SELECT
>         "c"
>     )
>   )
>   UNION (
>     SELECT
>       "d"
>   )
>   UNION (
>     SELECT
>       "e"
>   )
> )
`,
			nil,
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d Builder", i), func(t *testing.T) {
			t.Parallel()
			sql, values := b.Build(c.stmt)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			t.Parallel()
			sql, values := bi.Build(c.stmt)
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
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewUnionAll(
				stmt.NewSelect(stmt.NewColumn("a")),
				stmt.NewUnionAll(
					stmt.NewUnionAll(
						stmt.NewSelect(stmt.NewColumn("b")),
						stmt.NewSelect(stmt.NewColumn("c")),
					),
					stmt.NewSelect(stmt.NewColumn("d")),
					stmt.NewSelect(stmt.NewColumn("e")),
				),
			),
			`(SELECT "a") UNION ALL (((SELECT "b") UNION ALL (SELECT "c")) UNION ALL (SELECT "d") UNION ALL (SELECT "e"))`,
			`> (
>   SELECT
>     "a"
> )
> UNION ALL (
>   (
>     (
>       SELECT
>         "b"
>     )
>     UNION ALL (
>       SELECT
>         "c"
>     )
>   )
>   UNION ALL (
>     SELECT
>       "d"
>   )
>   UNION ALL (
>     SELECT
>       "e"
>   )
> )
`,
			nil,
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d Builder", i), func(t *testing.T) {
			t.Parallel()
			sql, values := b.Build(c.stmt)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			t.Parallel()
			sql, values := bi.Build(c.stmt)
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
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewIntersect(
				stmt.NewSelect(stmt.NewColumn("a")),
				stmt.NewIntersect(
					stmt.NewIntersect(
						stmt.NewSelect(stmt.NewColumn("b")),
						stmt.NewSelect(stmt.NewColumn("c")),
					),
					stmt.NewSelect(stmt.NewColumn("d")),
					stmt.NewSelect(stmt.NewColumn("e")),
				),
			),
			`(SELECT "a") INTERSECT (((SELECT "b") INTERSECT (SELECT "c")) INTERSECT (SELECT "d") INTERSECT (SELECT "e"))`,
			`> (
>   SELECT
>     "a"
> )
> INTERSECT (
>   (
>     (
>       SELECT
>         "b"
>     )
>     INTERSECT (
>       SELECT
>         "c"
>     )
>   )
>   INTERSECT (
>     SELECT
>       "d"
>   )
>   INTERSECT (
>     SELECT
>       "e"
>   )
> )
`,
			nil,
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d Builder", i), func(t *testing.T) {
			t.Parallel()
			sql, values := b.Build(c.stmt)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			t.Parallel()
			sql, values := bi.Build(c.stmt)
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
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewIntersectAll(
				stmt.NewSelect(stmt.NewColumn("a")),
				stmt.NewIntersectAll(
					stmt.NewIntersectAll(
						stmt.NewSelect(stmt.NewColumn("b")),
						stmt.NewSelect(stmt.NewColumn("c")),
					),
					stmt.NewSelect(stmt.NewColumn("d")),
					stmt.NewSelect(stmt.NewColumn("e")),
				),
			),
			`(SELECT "a") INTERSECT ALL (((SELECT "b") INTERSECT ALL (SELECT "c")) INTERSECT ALL (SELECT "d") INTERSECT ALL (SELECT "e"))`,
			`> (
>   SELECT
>     "a"
> )
> INTERSECT ALL (
>   (
>     (
>       SELECT
>         "b"
>     )
>     INTERSECT ALL (
>       SELECT
>         "c"
>     )
>   )
>   INTERSECT ALL (
>     SELECT
>       "d"
>   )
>   INTERSECT ALL (
>     SELECT
>       "e"
>   )
> )
`,
			nil,
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d Builder", i), func(t *testing.T) {
			t.Parallel()
			sql, values := b.Build(c.stmt)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			t.Parallel()
			sql, values := bi.Build(c.stmt)
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
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewExcept(
				stmt.NewSelect(stmt.NewColumn("a")),
				stmt.NewExcept(
					stmt.NewExcept(
						stmt.NewSelect(stmt.NewColumn("b")),
						stmt.NewSelect(stmt.NewColumn("c")),
					),
					stmt.NewSelect(stmt.NewColumn("d")),
					stmt.NewSelect(stmt.NewColumn("e")),
				),
			),
			`(SELECT "a") EXCEPT (((SELECT "b") EXCEPT (SELECT "c")) EXCEPT (SELECT "d") EXCEPT (SELECT "e"))`,
			`> (
>   SELECT
>     "a"
> )
> EXCEPT (
>   (
>     (
>       SELECT
>         "b"
>     )
>     EXCEPT (
>       SELECT
>         "c"
>     )
>   )
>   EXCEPT (
>     SELECT
>       "d"
>   )
>   EXCEPT (
>     SELECT
>       "e"
>   )
> )
`,
			nil,
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d Builder", i), func(t *testing.T) {
			t.Parallel()
			sql, values := b.Build(c.stmt)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			t.Parallel()
			sql, values := bi.Build(c.stmt)
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
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewExceptAll(
				stmt.NewSelect(stmt.NewColumn("a")),
				stmt.NewExceptAll(
					stmt.NewExceptAll(
						stmt.NewSelect(stmt.NewColumn("b")),
						stmt.NewSelect(stmt.NewColumn("c")),
					),
					stmt.NewSelect(stmt.NewColumn("d")),
					stmt.NewSelect(stmt.NewColumn("e")),
				),
			),
			`(SELECT "a") EXCEPT ALL (((SELECT "b") EXCEPT ALL (SELECT "c")) EXCEPT ALL (SELECT "d") EXCEPT ALL (SELECT "e"))`,
			`> (
>   SELECT
>     "a"
> )
> EXCEPT ALL (
>   (
>     (
>       SELECT
>         "b"
>     )
>     EXCEPT ALL (
>       SELECT
>         "c"
>     )
>   )
>   EXCEPT ALL (
>     SELECT
>       "d"
>   )
>   EXCEPT ALL (
>     SELECT
>       "e"
>   )
> )
`,
			nil,
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d Builder", i), func(t *testing.T) {
			t.Parallel()
			sql, values := b.Build(c.stmt)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			t.Parallel()
			sql, values := bi.Build(c.stmt)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}
