package sqlabble_test

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	q "github.com/sqlabble/sqlabble"
	"github.com/sqlabble/sqlabble/builder"
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
		sql, values := builder.MySQL.Build(
			q.CreateTableIfNotExists(
				q.Table("user"),
			).Definitions(
				q.Column("id").Define("int"),
				q.Column("name").Define("varchar(20)"),
				q.Column("avatar").Define("varchar(255)"),
			),
		)
		_, err := db.Exec(sql, values...)
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		sql, values := builder.MySQL.Build(
			q.CreateTableIfNotExists(
				q.Table("comment"),
			).Definitions(
				q.Column("id").Define("int"),
				q.Column("body").Define("varchar(255)"),
				q.Column("author_id").Define("int"),
			),
		)
		_, err := db.Exec(sql, values...)
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		sql, values := builder.MySQL.Build(
			q.CreateTableIfNotExists(
				q.Table("post"),
			).Definitions(
				q.Column("id").Define("int"),
				q.Column("title").Define("varchar(20)"),
				q.Column("body").Define("varchar(255)"),
				q.Column("author_id").Define("int"),
				q.Column("comment_id").Define("int"),
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
		sql, values := builder.MySQL.Build(
			q.InsertInto(
				q.Table("user"),
				q.Column("id"), q.Column("name"), q.Column("avatar"),
			).Values(
				q.Vals(1, "foo", "http://example.com/foo.jpg"),
				q.Vals(2, "bar", "http://example.com/bar.jpg"),
				q.Vals(3, "baz", "http://example.com/baz.jpg"),
			),
		)
		_, err := db.Exec(sql, values...)
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		sql, values := builder.MySQL.Build(
			q.InsertInto(
				q.Table("comment"),
				q.Column("id"), q.Column("body"), q.Column("author_id"),
			).Values(
				q.Vals(1, "abcdefg", 3),
			),
		)
		_, err := db.Exec(sql, values...)
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		sql, values := builder.MySQL.Build(
			q.InsertInto(
				q.Table("post"),
				q.Column("id"), q.Column("title"), q.Column("body"), q.Column("author_id"), q.Column("comment_id"),
			).Values(
				q.Vals(1, "this is title", "this is body", 2, 1),
			),
		)
		_, err := db.Exec(sql, values...)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestSelect(t *testing.T) {
	sql, values := builder.MySQL.Build(
		q.Select(
			q.Column("id").As("User.ID"),
			q.Column("name").As("User.Name"),
			q.Column("avatar").As("User.Avatar"),
		).From(
			q.Table("user"),
		).Where(
			q.Column("id").Eq(q.Val(3)),
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

func TestNewSession(t *testing.T) {
	type args struct {
		builderName string
	}
	tests := []struct {
		name    string
		args    args
		want    *q.Session
		wantErr error
	}{
		{
			name: "check `standard`",
			args: args{
				builderName: "standard",
			},
			want: &q.Session{
				DB:      nil,
				Builder: &builder.Standard,
			},
			wantErr: nil,
		},
		{
			name: "check `standard_indented`",
			args: args{
				builderName: "standard_indented",
			},
			want: &q.Session{
				DB:      nil,
				Builder: &builder.StandardIndented,
			},
			wantErr: nil,
		},
		{
			name: "check `mysql`",
			args: args{
				builderName: "mysql",
			},
			want: &q.Session{
				DB:      nil,
				Builder: &builder.MySQL,
			},
			wantErr: nil,
		},
		{
			name: "check `mysql_indented`",
			args: args{
				builderName: "mysql_indented",
			},
			want: &q.Session{
				DB:      nil,
				Builder: &builder.MySQLIndented,
			},
			wantErr: nil,
		},
		{
			name: "check `hoge_fuga`",
			args: args{
				builderName: "hoge_fuga",
			},
			want:    nil,
			wantErr: builder.NewErrBuilderNotSupported("hoge_fuga"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := q.NewSession(tt.args.builderName, nil)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("NewSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSession() = %v, want %v", got, tt.want)
			}
		})
	}
}
