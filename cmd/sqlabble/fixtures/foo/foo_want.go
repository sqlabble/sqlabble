package foo

import (
	"database/sql"
	"fmt"

	"github.com/minodisk/sqlabble/stmt"
)

type UserTable struct {
	Prof *ProfileTable
}

func NewUserTable() *UserTable {
	return &UserTable{}
}

func (u *UserTable) Table() *stmt.Table {
	return stmt.NewTable("users")
}

func (u *UserTable) ColumnUserID() *stmt.Column {
	return u.Table().Column("user_id")
}

func (u *UserTable) ColumnName() *stmt.Column {
	return u.Table().Column("name")
}

func (u *UserTable) ColumnAvatar() *stmt.Column {
	return u.Table().Column("avatar")
}

func (u *UserTable) Columns() []*stmt.Column {
	return []*stmt.Column{
		u.ColumnUserID(),
		u.ColumnName(),
		u.ColumnAvatar(),
	}
}

type UserMapper struct {
	UserTable      *stmt.TableAlias
	ProfileTable   *stmt.TableAlias
	UserColumns    UserMapperUserColumns
	ProfileColumns UserMapperProfileColumns
}

type UserMapperUserColumns struct {
	UserID *stmt.ColumnAlias
	Name   *stmt.ColumnAlias
	Avatar *stmt.ColumnAlias
}

type UserMapperProfileColumns struct {
	ProfileID *stmt.ColumnAlias
	Body      *stmt.ColumnAlias
	UserID    *stmt.ColumnAlias
}

func NewUserMapper() *UserMapper {
	u := NewUserTable()
	return &UserMapper{
		UserTable:    u.Table().As("users"),
		ProfileTable: u.Prof.Table().As("profiles"),
		UserColumns: UserMapperUserColumns{
			UserID: u.ColumnUserID().As("users.user_id"),
			Name:   u.ColumnName().As("users.name"),
			Avatar: u.ColumnAvatar().As("users.avatar"),
		},
		ProfileColumns: UserMapperProfileColumns{
			ProfileID: u.Prof.ColumnProfileID().As("profiles.profile_id"),
			Body:      u.Prof.ColumnBody().As("profiles.body"),
			UserID:    u.Prof.ColumnUserID().As("profiles.user_id"),
		},
	}
}

func (u *UserMapper) Columns() []*stmt.ColumnAlias {
	return []*stmt.ColumnAlias{
		u.UserColumns.UserID,
		u.UserColumns.Name,
		u.UserColumns.Avatar,
		u.ProfileColumns.ProfileID,
		u.ProfileColumns.Body,
		u.ProfileColumns.UserID,
	}
}

func (u *UserMapper) Map(rows *sql.Rows) ([]User, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	fmt.Println(cols)
	dist := []User{}
	for rows.Next() {
		d := User{}
		ar := map[string]interface{}{
			"users.user_id":       &d.UserID,
			"users.name":          &d.Name,
			"users.avatar":        &d.Avatar,
			"profiles.profile_id": &d.Prof.ProfileID,
			"profiles.body":       &d.Prof.Body,
			"profiles.user_id":    &d.Prof.UserID,
		}
		refs := make([]interface{}, len(cols))
		for i, c := range cols {
			refs[i] = ar[c]
		}
		if err := rows.Scan(refs...); err != nil {
			return nil, err
		}
		dist = append(dist, d)
	}
	return dist, nil
}

type ProfileTable struct{}

func NewProfileTable() *ProfileTable {
	return &ProfileTable{}
}

func (p *ProfileTable) Table() *stmt.Table {
	return stmt.NewTable("profiles")
}

func (p *ProfileTable) ColumnProfileID() *stmt.Column {
	return p.Table().Column("profile_id")
}

func (p *ProfileTable) ColumnBody() *stmt.Column {
	return p.Table().Column("body")
}

func (p *ProfileTable) ColumnUserID() *stmt.Column {
	return p.Table().Column("user_id")
}

func (p *ProfileTable) Columns() []*stmt.Column {
	return []*stmt.Column{
		p.ColumnProfileID(),
		p.ColumnBody(),
		p.ColumnUserID(),
	}
}
