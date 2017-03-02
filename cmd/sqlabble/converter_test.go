package sqlabble_test

import (
	"fmt"
	"testing"

	"github.com/minodisk/sqlabble/cmd/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

func TestGenerate(t *testing.T) {
	for i, c := range []struct {
		input string
		want  string
	}{
		{
			`package tables
type Person struct {
	PersonID int
	FamilyName string ` + "`" + `db:"fname"` + "`" + `
	LastName string ` + "`" + `db:"lname"` + "`" + `
	BirthDate time.Time
	SocialSecurityNumber string ` + "`" + `db:"-"` + "`" + `
	password string
}`,
			``,
		},
		{
			`package tables
// db:""
type Person struct {
	PersonID int
	FamilyName string ` + "`" + `db:"fname"` + "`" + `
	LastName string ` + "`" + `db:"lname"` + "`" + `
	BirthDate time.Time
	SocialSecurityNumber string ` + "`" + `db:"-"` + "`" + `
	password string
}`,
			``,
		},
		{
			`package tables
// +db:
type Person struct {
	PersonID int
	FamilyName string ` + "`" + `db:"fname"` + "`" + `
	LastName string ` + "`" + `db:"lname"` + "`" + `
	BirthDate time.Time
	SocialSecurityNumber string ` + "`" + `db:"-"` + "`" + `
	password string
}`,
			``,
		},
		{
			`package tables
// +db:""
type Person struct {
	PersonID int
	FamilyName string ` + "`" + `db:"fname"` + "`" + `
	LastName string ` + "`" + `db:"lname"` + "`" + `
	BirthDate time.Time
	SocialSecurityNumber string ` + "`" + `db:"-"` + "`" + `
	password string
}`,
			`package tables

import "github.com/minodisk/sqlabble/stmt"

func (p Person) Table() stmt.Table {
	return stmt.NewTable("person")
}

func (p Person) Columns() []stmt.Column {
	return []stmt.Column{
		p.ColumnPersonID(),
		p.ColumnFamilyName(),
		p.ColumnLastName(),
		p.ColumnBirthDate(),
	}
}

func (p Person) ColumnPersonID() stmt.Column {
	return stmt.NewColumn("person_id")
}

func (p Person) ColumnFamilyName() stmt.Column {
	return stmt.NewColumn("fname")
}

func (p Person) ColumnLastName() stmt.Column {
	return stmt.NewColumn("lname")
}

func (p Person) ColumnBirthDate() stmt.Column {
	return stmt.NewColumn("birth_date")
}
`,
		},
		{
			`package tables
// +db:"persons"
type Person struct {
	PersonID int
	FamilyName string ` + "`" + `db:"fname"` + "`" + `
	LastName string ` + "`" + `db:"lname"` + "`" + `
	BirthDate time.Time
	SocialSecurityNumber string ` + "`" + `db:"-"` + "`" + `
	password string
}`,
			`package tables

import "github.com/minodisk/sqlabble/stmt"

func (p Person) Table() stmt.Table {
	return stmt.NewTable("persons")
}

func (p Person) Columns() []stmt.Column {
	return []stmt.Column{
		p.ColumnPersonID(),
		p.ColumnFamilyName(),
		p.ColumnLastName(),
		p.ColumnBirthDate(),
	}
}

func (p Person) ColumnPersonID() stmt.Column {
	return stmt.NewColumn("person_id")
}

func (p Person) ColumnFamilyName() stmt.Column {
	return stmt.NewColumn("fname")
}

func (p Person) ColumnLastName() stmt.Column {
	return stmt.NewColumn("lname")
}

func (p Person) ColumnBirthDate() stmt.Column {
	return stmt.NewColumn("birth_date")
}
`,
		},
		{
			`package tables
type Person struct { // +db:"persons"
	PersonID int
	FamilyName string ` + "`" + `db:"fname"` + "`" + `
	LastName string ` + "`" + `db:"lname"` + "`" + `
	BirthDate time.Time
	SocialSecurityNumber string ` + "`" + `db:"-"` + "`" + `
	password string
}`,
			`package tables

import "github.com/minodisk/sqlabble/stmt"

func (p Person) Table() stmt.Table {
	return stmt.NewTable("persons")
}

func (p Person) Columns() []stmt.Column {
	return []stmt.Column{
		p.ColumnPersonID(),
		p.ColumnFamilyName(),
		p.ColumnLastName(),
		p.ColumnBirthDate(),
	}
}

func (p Person) ColumnPersonID() stmt.Column {
	return stmt.NewColumn("person_id")
}

func (p Person) ColumnFamilyName() stmt.Column {
	return stmt.NewColumn("fname")
}

func (p Person) ColumnLastName() stmt.Column {
	return stmt.NewColumn("lname")
}

func (p Person) ColumnBirthDate() stmt.Column {
	return stmt.NewColumn("birth_date")
}
`,
		},
		{
			`package tables
type Person struct {
	PersonID int
	FamilyName string ` + "`" + `db:"fname"` + "`" + `
	LastName string ` + "`" + `db:"lname"` + "`" + `
	BirthDate time.Time
	SocialSecurityNumber string ` + "`" + `db:"-"` + "`" + `
	password string
} // +db:"persons"`,
			`package tables

import "github.com/minodisk/sqlabble/stmt"

func (p Person) Table() stmt.Table {
	return stmt.NewTable("persons")
}

func (p Person) Columns() []stmt.Column {
	return []stmt.Column{
		p.ColumnPersonID(),
		p.ColumnFamilyName(),
		p.ColumnLastName(),
		p.ColumnBirthDate(),
	}
}

func (p Person) ColumnPersonID() stmt.Column {
	return stmt.NewColumn("person_id")
}

func (p Person) ColumnFamilyName() stmt.Column {
	return stmt.NewColumn("fname")
}

func (p Person) ColumnLastName() stmt.Column {
	return stmt.NewColumn("lname")
}

func (p Person) ColumnBirthDate() stmt.Column {
	return stmt.NewColumn("birth_date")
}
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			g, err := sqlabble.Convert([]byte(c.input))
			if err != nil {
				t.Fatal(err)
			}
			got := string(g)
			if got != c.want {
				t.Error(diff.SQL(got, c.want))
			}
		})
	}
}
