package main_test

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	q "github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/builder"
	"github.com/minodisk/sqlabble/cmd/sqlabble/fixtures/foo"
	"github.com/minodisk/sqlabble/internal/diff"
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
	db.Exec("DROP TABLE IF EXISTS users, profiles")
	os.Exit(code)
}

func TestMapper(t *testing.T) {
	u := foo.NewUserDB()
	{
		query, values := builder.Standard.Build(
			q.CreateTableIfNotExists(
				u.Table,
			).Definitions(
				u.UserIDColumn.Define("int"),
				u.NameColumn.Define("varchar(20)"),
				u.AvatarColumn.Define("varchar(255)"),
			),
		)
		_, err := db.Exec(query, values...)
		if err != nil {
			t.Fatal(err)
		}
	}

	pt := foo.NewProfileDB()
	{
		query, values := builder.Standard.Build(
			q.CreateTableIfNotExists(
				pt.Table,
			).Definitions(
				pt.ProfileIDColumn.Define("int"),
				pt.BodyColumn.Define("varchar(255)"),
				pt.UserIDColumn.Define("int"),
			),
		)
		_, err := db.Exec(query, values...)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		query, values := builder.MySQLIndented.Build(
			q.InsertInto(
				u.Table,
				u.Columns()...,
			).Values(
				q.Params(1, "foo", "http://example.com/foo.jpg"),
				q.Params(2, "bar", "http://example.com/bar.jpg"),
				q.Params(3, "baz", "http://example.com/baz.jpg"),
			),
		)
		_, err := db.Exec(query, values...)
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		query, values := builder.MySQLIndented.Build(
			q.InsertInto(
				pt.Table,
				pt.Columns()...,
			).Values(
				q.Params(1, "foo's profile", 1),
				q.Params(2, "baz's profile", 3),
				q.Params(3, "bar's profile", 2),
			),
		)
		_, err := db.Exec(query, values...)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		u := foo.NewUserDB()
		query, values := builder.MySQLIndented.Build(
			q.Select(u.Selectors()...).
				From(
					u.TableAlias.
						LeftJoin(u.Prof.TableAlias).
						On(
							u.UserIDColumn,
							u.Prof.UserIDColumn,
						),
				).
				OrderBy(
					u.UserIDColumn.Asc(),
				),
		)

		rows, err := db.Query(query, values...)
		if err != nil {
			t.Fatal(err)
		}
		got, err := u.Map(rows)
		if err != nil {
			t.Fatal(err)
		}

		want := []foo.User{
			{
				UserID: 1,
				Name:   "foo",
				Avatar: "http://example.com/foo.jpg",
				Prof: foo.Profile{
					ProfileID: 1,
					Body:      "foo's profile",
					UserID:    1,
				},
			},
			{
				UserID: 2,
				Name:   "bar",
				Avatar: "http://example.com/bar.jpg",
				Prof: foo.Profile{
					ProfileID: 3,
					Body:      "bar's profile",
					UserID:    2,
				},
			},
			{
				UserID: 3,
				Name:   "baz",
				Avatar: "http://example.com/baz.jpg",
				Prof: foo.Profile{
					ProfileID: 2,
					Body:      "baz's profile",
					UserID:    3,
				},
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Error(diff.Values(got, want))
		}
	}

	{
		p := foo.NewProfileDB()
		query, values := builder.MySQLIndented.Build(
			q.Select(p.Selectors()...).
				From(
					p.TableAlias,
				).
				OrderBy(
					p.ProfileIDColumn.Asc(),
				),
		)

		rows, err := db.Query(query, values...)
		if err != nil {
			t.Fatal(err)
		}
		got, err := p.Map(rows)
		if err != nil {
			t.Fatal(err)
		}

		want := []foo.Profile{
			{
				ProfileID: 1,
				Body:      "foo's profile",
				UserID:    1,
			},
			{
				ProfileID: 2,
				Body:      "baz's profile",
				UserID:    3,
			},
			{
				ProfileID: 3,
				Body:      "bar's profile",
				UserID:    2,
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Error(diff.Values(got, want))
		}
	}
}
