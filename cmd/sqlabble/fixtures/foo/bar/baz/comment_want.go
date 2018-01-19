package baz

import (
	"database/sql"
	"strings"

	"github.com/sqlabble/sqlabble/stmt"
)

type CommentDB struct {
	Table                stmt.Table
	TableAlias           stmt.TableAlias
	CommentIDColumn      stmt.Column
	CommentIDColumnAlias stmt.ColumnAlias
	ArticleIDColumn      stmt.Column
	ArticleIDColumnAlias stmt.ColumnAlias
	Article              ArticleDB
}

func NewCommentDB(aliases ...string) CommentDB {
	alias := strings.Join(aliases, ".")
	if alias == "" {
		alias = "comments"
	}
	return CommentDB{
		Table:                stmt.NewTable("comments"),
		TableAlias:           stmt.NewTable("comments").As(alias),
		CommentIDColumn:      stmt.NewTableAlias(alias).Column("comment_id"),
		CommentIDColumnAlias: stmt.NewTableAlias(alias).Column("comment_id").As(strings.Join(append(aliases, "CommentID"), ".")),
		ArticleIDColumn:      stmt.NewTableAlias(alias).Column("article_id"),
		ArticleIDColumnAlias: stmt.NewTableAlias(alias).Column("article_id").As(strings.Join(append(aliases, "ArticleID"), ".")),
		Article:              NewArticleDB(append(aliases, "Article")...),
	}
}

func (c CommentDB) Register(mapper map[string]interface{}, dist *Comment, aliases ...string) {
	mapper[strings.Join(append(aliases, "CommentID"), ".")] = &dist.CommentID
	mapper[strings.Join(append(aliases, "ArticleID"), ".")] = &dist.ArticleID
	c.Article.Register(mapper, &dist.Article, append(aliases, "Article")...)
}

func (c CommentDB) Columns() []stmt.Column {
	return []stmt.Column{
		c.CommentIDColumn,
		c.ArticleIDColumn,
	}
}

func (c CommentDB) ColumnAliases() []stmt.ColumnAlias {
	aliases := []stmt.ColumnAlias{
		c.CommentIDColumnAlias,
		c.ArticleIDColumnAlias,
	}
	aliases = append(aliases, c.Article.ColumnAliases()...)
	return aliases
}

func (c CommentDB) Selectors() []stmt.ValOrColOrAliasOrFuncOrSubOrFormula {
	as := c.ColumnAliases()
	is := make([]stmt.ValOrColOrAliasOrFuncOrSubOrFormula, len(as))
	for i, a := range as {
		is[i] = a
	}
	return is
}

func (c CommentDB) Map(rows *sql.Rows) ([]Comment, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	dist := []Comment{}
	for rows.Next() {
		mapper := make(map[string]interface{})
		di := Comment{}
		c.Register(mapper, &di)
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
