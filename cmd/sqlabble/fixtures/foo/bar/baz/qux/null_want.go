package qux

import (
	"database/sql"
	"strings"

	"github.com/minodisk/sqlabble/stmt"
)

type ArticleDB struct {
	Table              stmt.Table
	TableAlias         stmt.TableAlias
	BoolColumn         stmt.Column
	BoolColumnAlias    stmt.ColumnAlias
	Float64Column      stmt.Column
	Float64ColumnAlias stmt.ColumnAlias
	Int64Column        stmt.Column
	Int64ColumnAlias   stmt.ColumnAlias
	StringColumn       stmt.Column
	StringColumnAlias  stmt.ColumnAlias
}

func NewArticleDB(aliases ...string) ArticleDB {
	alias := strings.Join(aliases, ".")
	if alias == "" {
		alias = "articles"
	}
	return ArticleDB{
		Table:              stmt.NewTable("articles"),
		TableAlias:         stmt.NewTable("articles").As(alias),
		BoolColumn:         stmt.NewTableAlias(alias).Column("bool"),
		BoolColumnAlias:    stmt.NewTableAlias(alias).Column("bool").As(strings.Join(append(aliases, "Bool"), ".")),
		Float64Column:      stmt.NewTableAlias(alias).Column("float64"),
		Float64ColumnAlias: stmt.NewTableAlias(alias).Column("float64").As(strings.Join(append(aliases, "Float64"), ".")),
		Int64Column:        stmt.NewTableAlias(alias).Column("int64"),
		Int64ColumnAlias:   stmt.NewTableAlias(alias).Column("int64").As(strings.Join(append(aliases, "Int64"), ".")),
		StringColumn:       stmt.NewTableAlias(alias).Column("string"),
		StringColumnAlias:  stmt.NewTableAlias(alias).Column("string").As(strings.Join(append(aliases, "String"), ".")),
	}
}

func (a ArticleDB) Register(mapper map[string]interface{}, dist *Article, aliases ...string) {
	mapper[strings.Join(append(aliases, "Bool"), ".")] = &dist.Bool
	mapper[strings.Join(append(aliases, "Float64"), ".")] = &dist.Float64
	mapper[strings.Join(append(aliases, "Int64"), ".")] = &dist.Int64
	mapper[strings.Join(append(aliases, "String"), ".")] = &dist.String
}

func (a ArticleDB) Columns() []stmt.Column {
	return []stmt.Column{
		a.BoolColumn,
		a.Float64Column,
		a.Int64Column,
		a.StringColumn,
	}
}

func (a ArticleDB) ColumnAliases() []stmt.ColumnAlias {
	aliases := []stmt.ColumnAlias{
		a.BoolColumnAlias,
		a.Float64ColumnAlias,
		a.Int64ColumnAlias,
		a.StringColumnAlias,
	}
	return aliases
}

func (a ArticleDB) Selectors() []stmt.ColOrAliasOrFuncOrSub {
	as := a.ColumnAliases()
	is := make([]stmt.ColOrAliasOrFuncOrSub, len(as))
	for i, a := range as {
		is[i] = a
	}
	return is
}

func (a ArticleDB) Map(rows *sql.Rows) ([]Article, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	dist := []Article{}
	for rows.Next() {
		mapper := make(map[string]interface{})
		di := Article{}
		a.Register(mapper, &di)
		refs := make([]interface{}, len(cols))
		for i, c := range cols {
			refs[i] = mapper[c]
		}
		if err := rows.Scan(refs...); err != nil {
			return nil, err
		}
		dist = append(dist, di)
	}
	return dist, nil
}
