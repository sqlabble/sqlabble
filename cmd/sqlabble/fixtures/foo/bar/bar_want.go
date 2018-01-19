package bar

import (
	"database/sql"
	"strings"

	"github.com/sqlabble/sqlabble/cmd/sqlabble/fixtures/foo"
	"github.com/sqlabble/sqlabble/stmt"
)

type ArticleDB struct {
	Table               stmt.Table
	TableAlias          stmt.TableAlias
	PostIDColumn        stmt.Column
	PostIDColumnAlias   stmt.ColumnAlias
	BodyColumn          stmt.Column
	BodyColumnAlias     stmt.ColumnAlias
	AuthorIDColumn      stmt.Column
	AuthorIDColumnAlias stmt.ColumnAlias
	Author              foo.UserDB
}

func NewArticleDB(aliases ...string) ArticleDB {
	alias := strings.Join(aliases, ".")
	if alias == "" {
		alias = "articles"
	}
	return ArticleDB{
		Table:               stmt.NewTable("articles"),
		TableAlias:          stmt.NewTable("articles").As(alias),
		PostIDColumn:        stmt.NewTableAlias(alias).Column("post_id"),
		PostIDColumnAlias:   stmt.NewTableAlias(alias).Column("post_id").As(strings.Join(append(aliases, "PostID"), ".")),
		BodyColumn:          stmt.NewTableAlias(alias).Column("body"),
		BodyColumnAlias:     stmt.NewTableAlias(alias).Column("body").As(strings.Join(append(aliases, "Body"), ".")),
		AuthorIDColumn:      stmt.NewTableAlias(alias).Column("author_id"),
		AuthorIDColumnAlias: stmt.NewTableAlias(alias).Column("author_id").As(strings.Join(append(aliases, "AuthorID"), ".")),
		Author:              foo.NewUserDB(append(aliases, "Author")...),
	}
}

func (a ArticleDB) Register(mapper map[string]interface{}, dist *Article, aliases ...string) {
	mapper[strings.Join(append(aliases, "PostID"), ".")] = &dist.PostID
	mapper[strings.Join(append(aliases, "Body"), ".")] = &dist.Body
	mapper[strings.Join(append(aliases, "AuthorID"), ".")] = &dist.AuthorID
	a.Author.Register(mapper, &dist.Author, append(aliases, "Author")...)
}

func (a ArticleDB) Columns() []stmt.Column {
	return []stmt.Column{
		a.PostIDColumn,
		a.BodyColumn,
		a.AuthorIDColumn,
	}
}

func (a ArticleDB) ColumnAliases() []stmt.ColumnAlias {
	aliases := []stmt.ColumnAlias{
		a.PostIDColumnAlias,
		a.BodyColumnAlias,
		a.AuthorIDColumnAlias,
	}
	aliases = append(aliases, a.Author.ColumnAliases()...)
	return aliases
}

func (a ArticleDB) Selectors() []stmt.ValOrColOrAliasOrFuncOrSubOrFormula {
	as := a.ColumnAliases()
	is := make([]stmt.ValOrColOrAliasOrFuncOrSubOrFormula, len(as))
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
