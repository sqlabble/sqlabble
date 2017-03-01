package converter_test

import (
	"fmt"
	"testing"

	"github.com/minodisk/sqlabble/cmd/gensqlabble/converter"
	"github.com/minodisk/sqlabble/internal/diff"
)

func TestSqlabble(t *testing.T) {
	for i, c := range []struct {
		input string
		want  string
	}{
		{
			`package testcase
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
			`package testcase
// db:""
type Person struct {
	PersonID int
	FamilyName string ` + "`" + `db:"fname"` + "`" + `
	LastName string ` + "`" + `db:"lname"` + "`" + `
	BirthDate time.Time
	SocialSecurityNumber string ` + "`" + `db:"-"` + "`" + `
	password string
}`,
			`package testcase

import "github.com/minodisk/sqlabble/statement"

func (p Person) Table() statement.Table {
	return statement.NewTable("person")
}

func (p Person) Columns() []statement.Column {
	return []statement.Column{
		p.ColumnPersonID(),
		p.ColumnFamilyName(),
		p.ColumnLastName(),
		p.ColumnBirthDate(),
	}
}

func (p Person) ColumnPersonID() statement.Column {
	return statement.NewColumn("person_id")
}

func (p Person) ColumnFamilyName() statement.Column {
	return statement.NewColumn("fname")
}

func (p Person) ColumnLastName() statement.Column {
	return statement.NewColumn("lname")
}

func (p Person) ColumnBirthDate() statement.Column {
	return statement.NewColumn("birth_date")
}
`,
		},
		{
			`package testcase
// db:"persons"
type Person struct {
	PersonID int
	FamilyName string ` + "`" + `db:"fname"` + "`" + `
	LastName string ` + "`" + `db:"lname"` + "`" + `
	BirthDate time.Time
	SocialSecurityNumber string ` + "`" + `db:"-"` + "`" + `
	password string
}`,
			`package testcase

import "github.com/minodisk/sqlabble/statement"

func (p Person) Table() statement.Table {
	return statement.NewTable("persons")
}

func (p Person) Columns() []statement.Column {
	return []statement.Column{
		p.ColumnPersonID(),
		p.ColumnFamilyName(),
		p.ColumnLastName(),
		p.ColumnBirthDate(),
	}
}

func (p Person) ColumnPersonID() statement.Column {
	return statement.NewColumn("person_id")
}

func (p Person) ColumnFamilyName() statement.Column {
	return statement.NewColumn("fname")
}

func (p Person) ColumnLastName() statement.Column {
	return statement.NewColumn("lname")
}

func (p Person) ColumnBirthDate() statement.Column {
	return statement.NewColumn("birth_date")
}
`,
		},
		{
			`package testcase
type Person struct { // db:"persons"
	PersonID int
	FamilyName string ` + "`" + `db:"fname"` + "`" + `
	LastName string ` + "`" + `db:"lname"` + "`" + `
	BirthDate time.Time
	SocialSecurityNumber string ` + "`" + `db:"-"` + "`" + `
	password string
}`,
			`package testcase

import "github.com/minodisk/sqlabble/statement"

func (p Person) Table() statement.Table {
	return statement.NewTable("persons")
}

func (p Person) Columns() []statement.Column {
	return []statement.Column{
		p.ColumnPersonID(),
		p.ColumnFamilyName(),
		p.ColumnLastName(),
		p.ColumnBirthDate(),
	}
}

func (p Person) ColumnPersonID() statement.Column {
	return statement.NewColumn("person_id")
}

func (p Person) ColumnFamilyName() statement.Column {
	return statement.NewColumn("fname")
}

func (p Person) ColumnLastName() statement.Column {
	return statement.NewColumn("lname")
}

func (p Person) ColumnBirthDate() statement.Column {
	return statement.NewColumn("birth_date")
}
`,
		},
		{
			`package testcase
type Person struct {
	PersonID int
	FamilyName string ` + "`" + `db:"fname"` + "`" + `
	LastName string ` + "`" + `db:"lname"` + "`" + `
	BirthDate time.Time
	SocialSecurityNumber string ` + "`" + `db:"-"` + "`" + `
	password string
} // db:"persons"`,
			`package testcase

import "github.com/minodisk/sqlabble/statement"

func (p Person) Table() statement.Table {
	return statement.NewTable("persons")
}

func (p Person) Columns() []statement.Column {
	return []statement.Column{
		p.ColumnPersonID(),
		p.ColumnFamilyName(),
		p.ColumnLastName(),
		p.ColumnBirthDate(),
	}
}

func (p Person) ColumnPersonID() statement.Column {
	return statement.NewColumn("person_id")
}

func (p Person) ColumnFamilyName() statement.Column {
	return statement.NewColumn("fname")
}

func (p Person) ColumnLastName() statement.Column {
	return statement.NewColumn("lname")
}

func (p Person) ColumnBirthDate() statement.Column {
	return statement.NewColumn("birth_date")
}
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			g, err := converter.Generate([]byte(c.input))
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
