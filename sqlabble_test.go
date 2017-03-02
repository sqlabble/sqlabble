package sqlabble_test

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
)

var (
	db *sql.DB
)

type User struct {
	ID     int
	Name   string
	Avatar string
}

type Comment struct {
	ID     int
	Body   string
	Author User
}

type Post struct {
	ID      int
	Title   string
	Body    string
	Author  User
	Comment Comment
}

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

func TestCreateTable(t *testing.T) {
	{
		sql, values := builder.Standard.Build(
			q.CreateTableIfNotExists(
				q.T("user"),
			).Definitions(
				q.C("id").Define("int"),
				q.C("name").Define("varchar(20)"),
				q.C("avatar").Define("varchar(255)"),
			),
		)
		_, err := db.Exec(sql, values...)
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		sql, values := builder.Standard.Build(
			q.CreateTableIfNotExists(
				q.T("comment"),
			).Definitions(
				q.C("id").Define("int"),
				q.C("body").Define("varchar(255)"),
				q.C("author_id").Define("int"),
			),
		)
		_, err := db.Exec(sql, values...)
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		sql, values := builder.Standard.Build(
			q.CreateTableIfNotExists(
				q.T("post"),
			).Definitions(
				q.C("id").Define("int"),
				q.C("title").Define("varchar(20)"),
				q.C("body").Define("varchar(255)"),
				q.C("author_id").Define("int"),
				q.C("comment_id").Define("int"),
			),
		)
		_, err := db.Exec(sql, values...)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestInsertInto(t *testing.T) {
	{
		sql, values := builder.Standard.Build(
			q.InsertInto(
				q.T("user"),
				q.C("id"), q.C("name"), q.C("avatar"),
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
				q.T("comment"),
				q.C("id"), q.C("body"), q.C("author_id"),
			).Values(
				q.Params(1, "abcdefg", 3),
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
				q.T("post"),
				q.C("id"), q.C("title"), q.C("body"), q.C("author_id"), q.C("comment_id"),
			).Values(
				q.Params(1, "this is title", "this is body", 2, 1),
			),
		)
		_, err := db.Exec(sql, values...)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestSelect(t *testing.T) {
	sql, values := builder.Standard.Build(
		q.Select(
			q.C("id").As("User.ID"),
			q.C("name").As("User.Name"),
			q.C("avatar").As("User.Avatar"),
		).From(
			q.T("user"),
		).Where(
			q.C("id").Eq(q.Param(3)),
		),
	)

	rows, err := db.Query(sql, values...)
	if err != nil {
		t.Fatal(err)
	}

	cols, err := rows.Columns()
	if err != nil {
		t.Fatal(err)
	}

	want := []string{
		"User.ID",
		"User.Name",
		"User.Avatar",
	}
	if !reflect.DeepEqual(cols, want) {
		t.Errorf("columns got %+v; want %+v", cols, want)
	}

	if !rows.Next() {
		t.Fatal("no rows")
	}

	var (
		id     int
		name   string
		avatar string
	)
	if err := rows.Scan(&id, &name, &avatar); err != nil {
		t.Fatal(err)
	}
	if id != 3 {
		t.Errorf("id got %d; want %d", id, 3)
	}
}
