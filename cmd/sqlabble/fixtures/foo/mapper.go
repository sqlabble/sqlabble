package foo

import (
	"database/sql"

	"github.com/minodisk/sqlabble/stmt"
)

func (u UserTable) Mapper() (stmt.From, func(sql.Rows) ([]User, error)) {
	return stmt.
			NewSelect(
				u.ColumnUserID().As("users.user_id"),
				u.ColumnName().As("users.name"),
				u.ColumnAvatar().As("users.avatar"),
				u.Profile.ColumnProfileID().As("profiles.profile_id"),
				u.Profile.ColumnBody().As("profiles.body"),
			).
			From(
				u.As("users").
					InnerJoin(u.Profile.As("profiles")).
					On(stmt.NewColumn("users.profile_id"), stmt.NewColumn("profiles.profile_id")),
			),
		func(rows sql.Rows) ([]User, error) {
			aliases, err := rows.Columns()
			if err != nil {
				return nil, err
			}
			dist := []User{}
			for rows.Next() {
				d := User{}
				aref := map[string]interface{}{
					"users.user_id":       &d.UserID,
					"users.name":          &d.Name,
					"users.avatar":        &d.Avatar,
					"profiles.profile_id": &d.Profile.ProfileID,
					"profiles.body":       &d.Profile.Body,
				}
				refs := make([]interface{}, len(aliases))
				for i, alias := range aliases {
					refs[i] = aref[alias]
				}
				if err := rows.Scan(refs...); err != nil {
					return nil, err
				}
				dist = append(dist, d)
			}
			return dist, nil
		}
}
