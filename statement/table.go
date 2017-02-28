package statement

import (
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Table struct {
	Prev
	Next
	name string
}

func NewTable(name string) *Table {
	return &Table{
		name: name,
	}
}

func (t *Table) As(alias string) *TableAs {
	a := NewTableAs(alias)
	a.table = t
	return a
}

func (t *Table) Join(table TableOrAlias) *Join {
	j := NewJoin(table)
	Link(t, j)
	return j
}

func (t *Table) InnerJoin(table TableOrAlias) *Join {
	j := NewInnerJoin(table)
	Link(t, j)
	return j
}

func (t *Table) LeftJoin(table TableOrAlias) *Join {
	j := NewLeftJoin(table)
	Link(t, j)
	return j
}

func (t *Table) RightJoin(table TableOrAlias) *Join {
	j := NewRightJoin(table)
	Link(t, j)
	return j
}

func (t *Table) nodeize() (tokenizer.Tokenizer, []interface{}) {
	if t.name == "" {
		return nil, nil
	}
	return tokenizer.NewLine(token.Word(t.name)), nil
}

func (t *Table) previous() Joiner {
	return nil
}

// isTableOrAlias always returns true.
// This method exists only to implement the interface TableOrAlias.
// This is a shit of duck typing, but anyway it works.
func (t *Table) isTableOrAlias() bool {
	return true
}

// isTableOrAliasOrJoiner always returns true.
// This method exists only to implement the interface TableOrAliasOrJoiner.
// This is a shit of duck typing, but anyway it works.
func (t *Table) isTableOrAliasOrJoiner() bool {
	return true
}
