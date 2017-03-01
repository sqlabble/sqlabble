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
			`package tables
type Person struct {
	PersonID int
	FamilyName string ` + "`" + `db:"fname"` + "`" + `
	LastName string ` + "`" + `db:"lname"` + "`" + `
	BirthDate time.Time
	SocialSecurityNumber string ` + "`" + `db:"-"` + "`" + `
}`,
			`package tables

import (
	"github.com/minodisk/sqlabble/statement"
)

func (p Person) Table() statement.Table {
	name = "person"
	if tn, ok := p.(TableNamer); ok {
		name = tn.TableName()
	}
	return statement.NewTable(name)
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
			got, err := converter.Generate(c.input)
			if err != nil {
				t.Fatal(err)
			}
			if got != c.want {
				t.Error(diff.SQL(got, c.want))
			}
		})
	}
}
