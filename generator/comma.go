package generator

import "strings"

type Comma struct {
	nodes []Node
}

func NewComma(nodes ...Node) Comma {
	return Comma{
		nodes: nodes,
	}
}

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
		sqls[i], vs = t.ToSQL(ctx.SetHead(", "))
		values = append(values, vs...)
	}
	return strings.Join(sqls, ""), values
}
