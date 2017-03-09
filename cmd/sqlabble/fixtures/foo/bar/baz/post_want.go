package baz

import (
	"database/sql"
	"strings"

	"github.com/minodisk/sqlabble/stmt"
)

type PostDB struct {
	Table                stmt.Table
	TableAlias           stmt.TableAlias
	PostIDColumn         stmt.Column
	PostIDColumnAlias    stmt.ColumnAlias
	ArticleIDColumn      stmt.Column
	ArticleIDColumnAlias stmt.ColumnAlias
	Article              ArticleDB
}

func NewPostDB(aliases ...string) PostDB {
	alias := strings.Join(aliases, ".")
	if alias == "" {
		alias = "posts"
	}
	return PostDB{
		Table:                stmt.NewTable("posts"),
		TableAlias:           stmt.NewTable("posts").As(alias),
		PostIDColumn:         stmt.NewTableAlias(alias).Column("post_id"),
		PostIDColumnAlias:    stmt.NewTableAlias(alias).Column("post_id").As(strings.Join(append(aliases, "PostID"), ".")),
		ArticleIDColumn:      stmt.NewTableAlias(alias).Column("article_id"),
		ArticleIDColumnAlias: stmt.NewTableAlias(alias).Column("article_id").As(strings.Join(append(aliases, "ArticleID"), ".")),
		Article:              NewArticleDB(append(aliases, "Article")...),
	}
}

func (p PostDB) Register(mapper map[string]interface{}, dist *Post, aliases ...string) {
	mapper[strings.Join(append(aliases, "PostID"), ".")] = &dist.PostID
	mapper[strings.Join(append(aliases, "ArticleID"), ".")] = &dist.ArticleID
	p.Article.Register(mapper, &dist.Article, append(aliases, "Article")...)
}

func (p PostDB) Columns() []stmt.Column {
	return []stmt.Column{
		p.PostIDColumn,
		p.ArticleIDColumn,
	}
}

func (p PostDB) ColumnAliases() []stmt.ColumnAlias {
	aliases := []stmt.ColumnAlias{
		p.PostIDColumnAlias,
		p.ArticleIDColumnAlias,
	}
	aliases = append(aliases, p.Article.ColumnAliases()...)
	return aliases
}

func (p PostDB) Selectors() []stmt.ColOrAliasOrFuncOrSub {
	as := p.ColumnAliases()
	is := make([]stmt.ColOrAliasOrFuncOrSub, len(as))
	for i, a := range as {
		is[i] = a
	}
	return is
}

func (p PostDB) Map(rows *sql.Rows) ([]Post, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	dist := []Post{}
	for rows.Next() {
		mapper := make(map[string]interface{})
		di := Post{}
		p.Register(mapper, &di)
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
