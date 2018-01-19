package stmt

import (
	"github.com/sqlabble/sqlabble/keyword"
	"github.com/sqlabble/sqlabble/token"
	"github.com/sqlabble/sqlabble/tokenizer"
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

func (c CreateTable) Definitions(defs ...Definition) Definitions {
	ds := NewDefinitions(defs...)
	ds.createTable = c
	return ds
}

func (c CreateTable) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return c.container()
}

func (c CreateTable) container() (tokenizer.Container, []interface{}) {
	line := tokenizer.NewLine(token.Word(keyword.CreateTable))
	if c.ifNotExists {
		line = line.A(
			token.Word(keyword.IfNotExists),
		)
	}

	middle, values := c.joiner.nodeize()
	return tokenizer.NewContainer(
		line,
	).SetMiddle(
		middle,
	), values
}
