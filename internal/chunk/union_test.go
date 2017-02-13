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
	for _, s := range []grammar.Statement{
		chunk.SetOperation{},
	} {
		if _, ok := s.(grammar.Clause); !ok {
			t.Errorf("%T should implement grammar.Clause", s)
		}
	}
}

func TestUnionSQL(t *testing.T) {
	builder := sqlabble.NewBuilder("", "")
	builderIndent := sqlabble.NewBuilder("> ", "  ")
	builderMySQL4 := sqlabble.NewBuilderForMySQL4("", "")
	builderIndentMySQL4 := sqlabble.NewBuilderForMySQL4("> ", "  ")
	for i, c := range []struct {
		statement         grammar.Statement
		sql               string
		sqlIndent         string
		sqlForMySQL       string
		sqlIndentForMySQL string
		values            []interface{}
	}{
		{
			chunk.NewUnion(
				chunk.NewSelect(chunk.NewColumn("a")),
				chunk.NewUnion(
					chunk.NewUnion(
						chunk.NewSelect(chunk.NewColumn("b")),
						chunk.NewSelect(chunk.NewColumn("c")),
					),
					chunk.NewSelect(chunk.NewColumn("d")),
					chunk.NewSelect(chunk.NewColumn("e")),
				),
			),
			"(SELECT a) UNION (((SELECT b) UNION (SELECT c)) UNION (SELECT d) UNION (SELECT e))",
			`> (
>   SELECT
>     a
> )
> UNION
> (
>   (
>     (
>       SELECT
>         b
>     )
>     UNION
>     (
>       SELECT
>         c
>     )
>   )
>   UNION
>   (
>     SELECT
>       d
>   )
>   UNION
>   (
>     SELECT
>       e
>   )
> )
`,
			"(SELECT a) UNION (SELECT b) UNION (SELECT c) UNION (SELECT d) UNION (SELECT e)",
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
> UNION
> (
>   SELECT
>     d
> )
> UNION
> (
>   SELECT
>     e
> )
`,
			[]interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("%d Builder", i), func(t *testing.T) {
			sql, values := builder.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := builderIndent.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildForMySQL4", i), func(t *testing.T) {
			sql, values := builderMySQL4.Build(c.statement)
			if sql != c.sqlForMySQL {
				t.Error(diff.SQL(sql, c.sqlForMySQL))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndentForMySQL4", i), func(t *testing.T) {
			sql, values := builderIndentMySQL4.Build(c.statement)
			if sql != c.sqlIndentForMySQL {
				t.Error(diff.SQL(sql, c.sqlIndentForMySQL))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

func TestUnionAllSQL(t *testing.T) {
	builder := sqlabble.NewBuilder("", "")
	builderIndent := sqlabble.NewBuilder("> ", "  ")
	builderMySQL4 := sqlabble.NewBuilderForMySQL4("", "")
	builderIndentMySQL4 := sqlabble.NewBuilderForMySQL4("> ", "  ")
	for i, c := range []struct {
		statement         grammar.Statement
		sql               string
		sqlIndent         string
		sqlForMySQL       string
		sqlIndentForMySQL string
		values            []interface{}
	}{
		{
			chunk.NewUnionAll(
				chunk.NewSelect(chunk.NewColumn("a")),
				chunk.NewUnionAll(
					chunk.NewUnionAll(
						chunk.NewSelect(chunk.NewColumn("b")),
						chunk.NewSelect(chunk.NewColumn("c")),
					),
					chunk.NewSelect(chunk.NewColumn("d")),
					chunk.NewSelect(chunk.NewColumn("e")),
				),
			),
			"(SELECT a) UNION ALL (((SELECT b) UNION ALL (SELECT c)) UNION ALL (SELECT d) UNION ALL (SELECT e))",
			`> (
>   SELECT
>     a
> )
> UNION ALL
> (
>   (
>     (
>       SELECT
>         b
>     )
>     UNION ALL
>     (
>       SELECT
>         c
>     )
>   )
>   UNION ALL
>   (
>     SELECT
>       d
>   )
>   UNION ALL
>   (
>     SELECT
>       e
>   )
> )
`,
			"(SELECT a) UNION ALL (SELECT b) UNION ALL (SELECT c) UNION ALL (SELECT d) UNION ALL (SELECT e)",
			`> (
>   SELECT
>     a
> )
> UNION ALL
> (
>   SELECT
>     b
> )
> UNION ALL
> (
>   SELECT
>     c
> )
> UNION ALL
> (
>   SELECT
>     d
> )
> UNION ALL
> (
>   SELECT
>     e
> )
`,
			[]interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("%d Builder", i), func(t *testing.T) {
			sql, values := builder.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := builderIndent.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildForMySQL4", i), func(t *testing.T) {
			sql, values := builderMySQL4.Build(c.statement)
			if sql != c.sqlForMySQL {
				t.Error(diff.SQL(sql, c.sqlForMySQL))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndentForMySQL4", i), func(t *testing.T) {
			sql, values := builderIndentMySQL4.Build(c.statement)
			if sql != c.sqlIndentForMySQL {
				t.Error(diff.SQL(sql, c.sqlIndentForMySQL))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

func TestIntersectSQL(t *testing.T) {
	builder := sqlabble.NewBuilder("", "")
	builderIndent := sqlabble.NewBuilder("> ", "  ")
	builderMySQL4 := sqlabble.NewBuilderForMySQL4("", "")
	builderIndentMySQL4 := sqlabble.NewBuilderForMySQL4("> ", "  ")
	for i, c := range []struct {
		statement         grammar.Statement
		sql               string
		sqlIndent         string
		sqlForMySQL       string
		sqlIndentForMySQL string
		values            []interface{}
	}{
		{
			chunk.NewIntersect(
				chunk.NewSelect(chunk.NewColumn("a")),
				chunk.NewIntersect(
					chunk.NewIntersect(
						chunk.NewSelect(chunk.NewColumn("b")),
						chunk.NewSelect(chunk.NewColumn("c")),
					),
					chunk.NewSelect(chunk.NewColumn("d")),
					chunk.NewSelect(chunk.NewColumn("e")),
				),
			),
			"(SELECT a) INTERSECT (((SELECT b) INTERSECT (SELECT c)) INTERSECT (SELECT d) INTERSECT (SELECT e))",
			`> (
>   SELECT
>     a
> )
> INTERSECT
> (
>   (
>     (
>       SELECT
>         b
>     )
>     INTERSECT
>     (
>       SELECT
>         c
>     )
>   )
>   INTERSECT
>   (
>     SELECT
>       d
>   )
>   INTERSECT
>   (
>     SELECT
>       e
>   )
> )
`,
			"(SELECT a) INTERSECT (SELECT b) INTERSECT (SELECT c) INTERSECT (SELECT d) INTERSECT (SELECT e)",
			`> (
>   SELECT
>     a
> )
> INTERSECT
> (
>   SELECT
>     b
> )
> INTERSECT
> (
>   SELECT
>     c
> )
> INTERSECT
> (
>   SELECT
>     d
> )
> INTERSECT
> (
>   SELECT
>     e
> )
`,
			[]interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("%d Builder", i), func(t *testing.T) {
			sql, values := builder.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := builderIndent.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildForMySQL4", i), func(t *testing.T) {
			sql, values := builderMySQL4.Build(c.statement)
			if sql != c.sqlForMySQL {
				t.Error(diff.SQL(sql, c.sqlForMySQL))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndentForMySQL4", i), func(t *testing.T) {
			sql, values := builderIndentMySQL4.Build(c.statement)
			if sql != c.sqlIndentForMySQL {
				t.Error(diff.SQL(sql, c.sqlIndentForMySQL))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

func TestIntersectAllSQL(t *testing.T) {
	builder := sqlabble.NewBuilder("", "")
	builderIndent := sqlabble.NewBuilder("> ", "  ")
	builderMySQL4 := sqlabble.NewBuilderForMySQL4("", "")
	builderIndentMySQL4 := sqlabble.NewBuilderForMySQL4("> ", "  ")
	for i, c := range []struct {
		statement         grammar.Statement
		sql               string
		sqlIndent         string
		sqlForMySQL       string
		sqlIndentForMySQL string
		values            []interface{}
	}{
		{
			chunk.NewIntersectAll(
				chunk.NewSelect(chunk.NewColumn("a")),
				chunk.NewIntersectAll(
					chunk.NewIntersectAll(
						chunk.NewSelect(chunk.NewColumn("b")),
						chunk.NewSelect(chunk.NewColumn("c")),
					),
					chunk.NewSelect(chunk.NewColumn("d")),
					chunk.NewSelect(chunk.NewColumn("e")),
				),
			),
			"(SELECT a) INTERSECT ALL (((SELECT b) INTERSECT ALL (SELECT c)) INTERSECT ALL (SELECT d) INTERSECT ALL (SELECT e))",
			`> (
>   SELECT
>     a
> )
> INTERSECT ALL
> (
>   (
>     (
>       SELECT
>         b
>     )
>     INTERSECT ALL
>     (
>       SELECT
>         c
>     )
>   )
>   INTERSECT ALL
>   (
>     SELECT
>       d
>   )
>   INTERSECT ALL
>   (
>     SELECT
>       e
>   )
> )
`,
			"(SELECT a) INTERSECT ALL (SELECT b) INTERSECT ALL (SELECT c) INTERSECT ALL (SELECT d) INTERSECT ALL (SELECT e)",
			`> (
>   SELECT
>     a
> )
> INTERSECT ALL
> (
>   SELECT
>     b
> )
> INTERSECT ALL
> (
>   SELECT
>     c
> )
> INTERSECT ALL
> (
>   SELECT
>     d
> )
> INTERSECT ALL
> (
>   SELECT
>     e
> )
`,
			[]interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("%d Builder", i), func(t *testing.T) {
			sql, values := builder.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := builderIndent.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildForMySQL4", i), func(t *testing.T) {
			sql, values := builderMySQL4.Build(c.statement)
			if sql != c.sqlForMySQL {
				t.Error(diff.SQL(sql, c.sqlForMySQL))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndentForMySQL4", i), func(t *testing.T) {
			sql, values := builderIndentMySQL4.Build(c.statement)
			if sql != c.sqlIndentForMySQL {
				t.Error(diff.SQL(sql, c.sqlIndentForMySQL))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

func TestExceptSQL(t *testing.T) {
	builder := sqlabble.NewBuilder("", "")
	builderIndent := sqlabble.NewBuilder("> ", "  ")
	builderMySQL4 := sqlabble.NewBuilderForMySQL4("", "")
	builderIndentMySQL4 := sqlabble.NewBuilderForMySQL4("> ", "  ")
	for i, c := range []struct {
		statement         grammar.Statement
		sql               string
		sqlIndent         string
		sqlForMySQL       string
		sqlIndentForMySQL string
		values            []interface{}
	}{
		{
			chunk.NewExcept(
				chunk.NewSelect(chunk.NewColumn("a")),
				chunk.NewExcept(
					chunk.NewExcept(
						chunk.NewSelect(chunk.NewColumn("b")),
						chunk.NewSelect(chunk.NewColumn("c")),
					),
					chunk.NewSelect(chunk.NewColumn("d")),
					chunk.NewSelect(chunk.NewColumn("e")),
				),
			),
			"(SELECT a) EXCEPT (((SELECT b) EXCEPT (SELECT c)) EXCEPT (SELECT d) EXCEPT (SELECT e))",
			`> (
>   SELECT
>     a
> )
> EXCEPT
> (
>   (
>     (
>       SELECT
>         b
>     )
>     EXCEPT
>     (
>       SELECT
>         c
>     )
>   )
>   EXCEPT
>   (
>     SELECT
>       d
>   )
>   EXCEPT
>   (
>     SELECT
>       e
>   )
> )
`,
			"(SELECT a) EXCEPT (SELECT b) EXCEPT (SELECT c) EXCEPT (SELECT d) EXCEPT (SELECT e)",
			`> (
>   SELECT
>     a
> )
> EXCEPT
> (
>   SELECT
>     b
> )
> EXCEPT
> (
>   SELECT
>     c
> )
> EXCEPT
> (
>   SELECT
>     d
> )
> EXCEPT
> (
>   SELECT
>     e
> )
`,
			[]interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("%d Builder", i), func(t *testing.T) {
			sql, values := builder.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := builderIndent.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildForMySQL4", i), func(t *testing.T) {
			sql, values := builderMySQL4.Build(c.statement)
			if sql != c.sqlForMySQL {
				t.Error(diff.SQL(sql, c.sqlForMySQL))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndentForMySQL4", i), func(t *testing.T) {
			sql, values := builderIndentMySQL4.Build(c.statement)
			if sql != c.sqlIndentForMySQL {
				t.Error(diff.SQL(sql, c.sqlIndentForMySQL))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

func TestExceptAllSQL(t *testing.T) {
	builder := sqlabble.NewBuilder("", "")
	builderIndent := sqlabble.NewBuilder("> ", "  ")
	builderMySQL4 := sqlabble.NewBuilderForMySQL4("", "")
	builderIndentMySQL4 := sqlabble.NewBuilderForMySQL4("> ", "  ")
	for i, c := range []struct {
		statement         grammar.Statement
		sql               string
		sqlIndent         string
		sqlForMySQL       string
		sqlIndentForMySQL string
		values            []interface{}
	}{
		{
			chunk.NewExceptAll(
				chunk.NewSelect(chunk.NewColumn("a")),
				chunk.NewExceptAll(
					chunk.NewExceptAll(
						chunk.NewSelect(chunk.NewColumn("b")),
						chunk.NewSelect(chunk.NewColumn("c")),
					),
					chunk.NewSelect(chunk.NewColumn("d")),
					chunk.NewSelect(chunk.NewColumn("e")),
				),
			),
			"(SELECT a) EXCEPT ALL (((SELECT b) EXCEPT ALL (SELECT c)) EXCEPT ALL (SELECT d) EXCEPT ALL (SELECT e))",
			`> (
>   SELECT
>     a
> )
> EXCEPT ALL
> (
>   (
>     (
>       SELECT
>         b
>     )
>     EXCEPT ALL
>     (
>       SELECT
>         c
>     )
>   )
>   EXCEPT ALL
>   (
>     SELECT
>       d
>   )
>   EXCEPT ALL
>   (
>     SELECT
>       e
>   )
> )
`,
			"(SELECT a) EXCEPT ALL (SELECT b) EXCEPT ALL (SELECT c) EXCEPT ALL (SELECT d) EXCEPT ALL (SELECT e)",
			`> (
>   SELECT
>     a
> )
> EXCEPT ALL
> (
>   SELECT
>     b
> )
> EXCEPT ALL
> (
>   SELECT
>     c
> )
> EXCEPT ALL
> (
>   SELECT
>     d
> )
> EXCEPT ALL
> (
>   SELECT
>     e
> )
`,
			[]interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("%d Builder", i), func(t *testing.T) {
			sql, values := builder.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := builderIndent.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildForMySQL4", i), func(t *testing.T) {
			sql, values := builderMySQL4.Build(c.statement)
			if sql != c.sqlForMySQL {
				t.Error(diff.SQL(sql, c.sqlForMySQL))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndentForMySQL4", i), func(t *testing.T) {
			sql, values := builderIndentMySQL4.Build(c.statement)
			if sql != c.sqlIndentForMySQL {
				t.Error(diff.SQL(sql, c.sqlIndentForMySQL))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}
