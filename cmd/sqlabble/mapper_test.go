package sqlabble_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	q "github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/builder"
	"github.com/minodisk/sqlabble/cmd/sqlabble/fixtures/foo"
)

var (
	db *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open("mysql", "root@tcp(mysql:3306)/sqlabble")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; ; i++ {
		fmt.Println("ping")
		if err := db.Ping(); err == nil {
			break
		}
		if i == 60 {
			fmt.Println("timeout")
			os.Exit(1)
		}
		time.Sleep(time.Second)
	}

	code := m.Run()
	if err := db.Close(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	db.Exec("DROP TABLE IF EXISTS user, comment, post")
	os.Exit(code)
}

func TestMapper(t *testing.T) {
	ut := foo.NewUserTable()
	{
		sql, values := builder.Standard.Build(
			q.CreateTableIfNotExists(
				ut.Table(),
			).Definitions(
				ut.ColumnUserID().Define("int"),
				ut.ColumnName().Define("varchar(20)"),
				ut.ColumnAvatar().Define("varchar(255)"),
			),
		)
		_, err := db.Exec(sql, values...)
		if err != nil {
			t.Fatal(err)
		}
	}

	pt := foo.NewProfileTable()
	{
		sql, values := builder.Standard.Build(
			q.CreateTableIfNotExists(
				pt.Table(),
			).Definitions(
				pt.ColumnProfileID().Define("int"),
				pt.ColumnBody().Define("varchar(255)"),
				pt.ColumnUserID().Define("int"),
			),
		)
		_, err := db.Exec(sql, values...)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		sql, values := builder.Standard.Build(
			q.InsertInto(
				ut.Table(),
				ut.Columns()...,
			).Values(
				q.Params(1, "foo", "http://example.com/foo.jpg"),
				q.Params(2, "bar", "http://example.com/bar.jpg"),
				q.Params(3, "baz", "http://example.com/baz.jpg"),
			),
		)
		_, err := db.Exec(sql, values...)
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		sql, values := builder.Standard.Build(
			q.InsertInto(
				pt.Table(),
				pt.Columns()...,
			).Values(
				q.Params(1, "foo's profile", 1),
				q.Params(2, "bar's profile", 3),
				q.Params(3, "baz's profile", 2),
			),
		)
		_, err := db.Exec(sql, values...)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		mapper := foo.NewUserMapper()
		query, values := builder.MySQL.Build(
			q.Select(mapper.Columns()...).
				From(
					mapper.UserTable.
						LeftJoin(mapper.ProfileTable).
						On(mapper.UserColumns.UserID, mapper.ProfileColumns.UserID),
				),
		)
		fmt.Println(query)
		rows, err := db.Query(query, values...)
		if err != nil {
			t.Fatal(err)
		}
		users, err := mapUsers(rows)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(users)
	}
}
