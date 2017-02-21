package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
)

type CreateTable struct {
	ifNotExists bool
	joiner      Joiner
}

func NewCreateTable(joiner Joiner) CreateTable {
	return CreateTable{
		ifNotExists: false,
		joiner:      joiner,
	}
}

func NewCreateTableIfNotExists(joiner Joiner) CreateTable {
	return CreateTable{
		ifNotExists: true,
		joiner:      joiner,
	}
}

func (c CreateTable) nodeize() (token.Tokenizer, []interface{}) {
	return c.container()
}

func (c CreateTable) container() (token.Container, []interface{}) {
	line := token.NewLine(token.Word(keyword.CreateTable))
	if c.ifNotExists {
		line = line.A(
			token.Space,
			token.Word(keyword.IfNotExists),
		)
	}

	middle, values := c.joiner.nodeize()
	return token.NewContainer(
		line,
	).SetMiddle(
		middle,
	), values
}

func (c CreateTable) Definitions(defs ...Definition) Definitions {
	ds := NewDefinitions(defs...)
	ds.createTable = c
	return ds
}
