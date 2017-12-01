package baz

import (
	"database/sql"
	"strings"

	"github.com/minodisk/sqlabble/stmt"
)

type ArticleDB struct {
	Table                stmt.Table
	TableAlias           stmt.TableAlias
	ArticleIDColumn      stmt.Column
	ArticleIDColumnAlias stmt.ColumnAlias
	SubjectColumn        stmt.Column
	SubjectColumnAlias   stmt.ColumnAlias
	BodyColumn           stmt.Column
	BodyColumnAlias      stmt.ColumnAlias
}

func NewArticleDB(aliases ...string) ArticleDB {
	alias := strings.Join(aliases, ".")
	if alias == "" {
		alias = "articles"
	}
	return ArticleDB{
		Table:                stmt.NewTable("articles"),
		TableAlias:           stmt.NewTable("articles").As(alias),
		ArticleIDColumn:      stmt.NewTableAlias(alias).Column("article_id"),
		ArticleIDColumnAlias: stmt.NewTableAlias(alias).Column("article_id").As(strings.Join(append(aliases, "ArticleID"), ".")),
		SubjectColumn:        stmt.NewTableAlias(alias).Column("subject"),
		SubjectColumnAlias:   stmt.NewTableAlias(alias).Column("subject").As(strings.Join(append(aliases, "Subject"), ".")),
		BodyColumn:           stmt.NewTableAlias(alias).Column("body"),
		BodyColumnAlias:      stmt.NewTableAlias(alias).Column("body").As(strings.Join(append(aliases, "Body"), ".")),
	}
}

func (a ArticleDB) Register(mapper map[string]interface{}, dist *Article, aliases ...string) {
	mapper[strings.Join(append(aliases, "ArticleID"), ".")] = &dist.ArticleID
	mapper[strings.Join(append(aliases, "Subject"), ".")] = &dist.Subject
	mapper[strings.Join(append(aliases, "Body"), ".")] = &dist.Body
}

func (a ArticleDB) Columns() []stmt.Column {
	return []stmt.Column{
		a.ArticleIDColumn,
		a.SubjectColumn,
		a.BodyColumn,
	}
}

func (a ArticleDB) ColumnAliases() []stmt.ColumnAlias {
	aliases := []stmt.ColumnAlias{
		a.ArticleIDColumnAlias,
		a.SubjectColumnAlias,
		a.BodyColumnAlias,
	}
	return aliases
}

func (a ArticleDB) Selectors() []stmt.ValOrColOrAliasOrFuncOrSub {
	as := a.ColumnAliases()
	is := make([]stmt.ValOrColOrAliasOrFuncOrSub, len(as))
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
