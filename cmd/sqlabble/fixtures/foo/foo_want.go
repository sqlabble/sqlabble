package foo

import (
	"github.com/minodisk/sqlabble/stmt"
)

type UserTable struct {
	stmt.Table
	Profile ProfileTable
}

func (u UserTable) NewUserTable() UserTable {
	return UserTable{
		Table: stmt.NewTable("users"),
	}
}

func (u UserTable) ColumnUserID() stmt.Column {
	return u.Table.Column("user_id")
}

func (u UserTable) ColumnName() stmt.Column {
	return u.Table.Column("name")
}

func (u UserTable) ColumnAvatar() stmt.Column {
	return u.Table.Column("avatar")
}

func (u UserTable) Columns() []stmt.Column {
	return []stmt.Column{
		u.ColumnUserID(),
		u.ColumnName(),
		u.ColumnAvatar(),
	}
}

type ProfileTable struct {
	stmt.Table
}

func (p ProfileTable) NewProfileTable() ProfileTable {
	return ProfileTable{
		Table: stmt.NewTable("profiles"),
	}
}

func (p ProfileTable) ColumnProfileID() stmt.Column {
	return p.Table.Column("profile_id")
}

func (p ProfileTable) ColumnBody() stmt.Column {
	return p.Table.Column("body")
}

func (p ProfileTable) Columns() []stmt.Column {
	return []stmt.Column{
		p.ColumnProfileID(),
		p.ColumnBody(),
	}
}
