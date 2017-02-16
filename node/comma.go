package node

import "strings"

// Comma is a Node that joins multiple nodes with commas.
type Comma struct {
	nodes []Node
}

// NewComma returns a new Comma.
func NewComma(nodes ...Node) Comma {
	return Comma{
		nodes: nodes,
	}
}

// ToSQL returns a query and a slice of values.
func (c Comma) ToSQL(ctx Context) (string, []interface{}) {
	sqls := make([]string, len(c.nodes))
	values := []interface{}{}
	for i, t := range c.nodes {
		var vs []interface{}
		if i == 0 {
			sqls[i], vs = t.ToSQL(ctx)
			values = append(values, vs...)
			continue
		}
		sqls[i], vs = t.ToSQL(ctx.setHead(", "))
		values = append(values, vs...)
	}
	return strings.Join(sqls, ""), values
}
