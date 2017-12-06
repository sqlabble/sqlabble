package foo

import (
	"database/sql"
	"strings"

	"github.com/minodisk/sqlabble/stmt"
)

type UserDB struct {
	Table             stmt.Table
	TableAlias        stmt.TableAlias
	UserIDColumn      stmt.Column
	UserIDColumnAlias stmt.ColumnAlias
	NameColumn        stmt.Column
	NameColumnAlias   stmt.ColumnAlias
	AvatarColumn      stmt.Column
	AvatarColumnAlias stmt.ColumnAlias
	Prof              ProfileDB
}

func NewUserDB(aliases ...string) UserDB {
	alias := strings.Join(aliases, ".")
	if alias == "" {
		alias = "users"
	}
	return UserDB{
		Table:             stmt.NewTable("users"),
		TableAlias:        stmt.NewTable("users").As(alias),
		UserIDColumn:      stmt.NewTableAlias(alias).Column("user_id"),
		UserIDColumnAlias: stmt.NewTableAlias(alias).Column("user_id").As(strings.Join(append(aliases, "UserID"), ".")),
		NameColumn:        stmt.NewTableAlias(alias).Column("name"),
		NameColumnAlias:   stmt.NewTableAlias(alias).Column("name").As(strings.Join(append(aliases, "Name"), ".")),
		AvatarColumn:      stmt.NewTableAlias(alias).Column("avatar"),
		AvatarColumnAlias: stmt.NewTableAlias(alias).Column("avatar").As(strings.Join(append(aliases, "Avatar"), ".")),
		Prof:              NewProfileDB(append(aliases, "Prof")...),
	}
}

func (u UserDB) Register(mapper map[string]interface{}, dist *User, aliases ...string) {
	mapper[strings.Join(append(aliases, "UserID"), ".")] = &dist.UserID
	mapper[strings.Join(append(aliases, "Name"), ".")] = &dist.Name
	mapper[strings.Join(append(aliases, "Avatar"), ".")] = &dist.Avatar
	u.Prof.Register(mapper, &dist.Prof, append(aliases, "Prof")...)
	mapper[strings.Join(append(aliases, "NumFriends"), ".")] = &dist.NumFriends
}

func (u UserDB) Columns() []stmt.Column {
	return []stmt.Column{
		u.UserIDColumn,
		u.NameColumn,
		u.AvatarColumn,
	}
}

func (u UserDB) ColumnAliases() []stmt.ColumnAlias {
	aliases := []stmt.ColumnAlias{
		u.UserIDColumnAlias,
		u.NameColumnAlias,
		u.AvatarColumnAlias,
	}
	aliases = append(aliases, u.Prof.ColumnAliases()...)
	return aliases
}

func (u UserDB) Selectors() []stmt.ValOrColOrAliasOrFuncOrSubOrFormula {
	as := u.ColumnAliases()
	is := make([]stmt.ValOrColOrAliasOrFuncOrSubOrFormula, len(as))
	for i, a := range as {
		is[i] = a
	}
	return is
}

func (u UserDB) Map(rows *sql.Rows) ([]User, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	dist := []User{}
	for rows.Next() {
		mapper := make(map[string]interface{})
		di := User{}
		u.Register(mapper, &di)
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

type ProfileDB struct {
	Table                stmt.Table
	TableAlias           stmt.TableAlias
	ProfileIDColumn      stmt.Column
	ProfileIDColumnAlias stmt.ColumnAlias
	BodyColumn           stmt.Column
	BodyColumnAlias      stmt.ColumnAlias
	UserIDColumn         stmt.Column
	UserIDColumnAlias    stmt.ColumnAlias
}

func NewProfileDB(aliases ...string) ProfileDB {
	alias := strings.Join(aliases, ".")
	if alias == "" {
		alias = "profiles"
	}
	return ProfileDB{
		Table:                stmt.NewTable("profiles"),
		TableAlias:           stmt.NewTable("profiles").As(alias),
		ProfileIDColumn:      stmt.NewTableAlias(alias).Column("profile_id"),
		ProfileIDColumnAlias: stmt.NewTableAlias(alias).Column("profile_id").As(strings.Join(append(aliases, "ProfileID"), ".")),
		BodyColumn:           stmt.NewTableAlias(alias).Column("body"),
		BodyColumnAlias:      stmt.NewTableAlias(alias).Column("body").As(strings.Join(append(aliases, "Body"), ".")),
		UserIDColumn:         stmt.NewTableAlias(alias).Column("user_id"),
		UserIDColumnAlias:    stmt.NewTableAlias(alias).Column("user_id").As(strings.Join(append(aliases, "UserID"), ".")),
	}
}

func (p ProfileDB) Register(mapper map[string]interface{}, dist *Profile, aliases ...string) {
	mapper[strings.Join(append(aliases, "ProfileID"), ".")] = &dist.ProfileID
	mapper[strings.Join(append(aliases, "Body"), ".")] = &dist.Body
	mapper[strings.Join(append(aliases, "UserID"), ".")] = &dist.UserID
}

func (p ProfileDB) Columns() []stmt.Column {
	return []stmt.Column{
		p.ProfileIDColumn,
		p.BodyColumn,
		p.UserIDColumn,
	}
}

func (p ProfileDB) ColumnAliases() []stmt.ColumnAlias {
	aliases := []stmt.ColumnAlias{
		p.ProfileIDColumnAlias,
		p.BodyColumnAlias,
		p.UserIDColumnAlias,
	}
	return aliases
}

func (p ProfileDB) Selectors() []stmt.ValOrColOrAliasOrFuncOrSubOrFormula {
	as := p.ColumnAliases()
	is := make([]stmt.ValOrColOrAliasOrFuncOrSubOrFormula, len(as))
	for i, a := range as {
		is[i] = a
	}
	return is
}

func (p ProfileDB) Map(rows *sql.Rows) ([]Profile, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	dist := []Profile{}
	for rows.Next() {
		mapper := make(map[string]interface{})
		di := Profile{}
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

type FriendDB struct {
	Table              stmt.Table
	TableAlias         stmt.TableAlias
	UserID1Column      stmt.Column
	UserID1ColumnAlias stmt.ColumnAlias
	UserID2Column      stmt.Column
	UserID2ColumnAlias stmt.ColumnAlias
}

func NewFriendDB(aliases ...string) FriendDB {
	alias := strings.Join(aliases, ".")
	if alias == "" {
		alias = "friends"
	}
	return FriendDB{
		Table:              stmt.NewTable("friends"),
		TableAlias:         stmt.NewTable("friends").As(alias),
		UserID1Column:      stmt.NewTableAlias(alias).Column("user_id_1"),
		UserID1ColumnAlias: stmt.NewTableAlias(alias).Column("user_id_1").As(strings.Join(append(aliases, "UserID1"), ".")),
		UserID2Column:      stmt.NewTableAlias(alias).Column("user_id_2"),
		UserID2ColumnAlias: stmt.NewTableAlias(alias).Column("user_id_2").As(strings.Join(append(aliases, "UserID2"), ".")),
	}
}

func (f FriendDB) Register(mapper map[string]interface{}, dist *Friend, aliases ...string) {
	mapper[strings.Join(append(aliases, "UserID1"), ".")] = &dist.UserID1
	mapper[strings.Join(append(aliases, "UserID2"), ".")] = &dist.UserID2
}

func (f FriendDB) Columns() []stmt.Column {
	return []stmt.Column{
		f.UserID1Column,
		f.UserID2Column,
	}
}

func (f FriendDB) ColumnAliases() []stmt.ColumnAlias {
	aliases := []stmt.ColumnAlias{
		f.UserID1ColumnAlias,
		f.UserID2ColumnAlias,
	}
	return aliases
}

func (f FriendDB) Selectors() []stmt.ValOrColOrAliasOrFuncOrSubOrFormula {
	as := f.ColumnAliases()
	is := make([]stmt.ValOrColOrAliasOrFuncOrSubOrFormula, len(as))
	for i, a := range as {
		is[i] = a
	}
	return is
}

func (f FriendDB) Map(rows *sql.Rows) ([]Friend, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	dist := []Friend{}
	for rows.Next() {
		mapper := make(map[string]interface{})
		di := Friend{}
		f.Register(mapper, &di)
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
