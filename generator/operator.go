package generator

import (
	"fmt"

	"github.com/minodisk/sqlabble/operator"
)

type Operator struct {
	operator   operator.Operator
	generators []Node
}

func NewOperator(operator operator.Operator, generators ...Node) Operator {
	return Operator{
		operator:   operator,
		generators: generators,
	}
}

func (o Operator) ToSQL(ctx Context) (string, []interface{}) {
	for {
		if len(o.generators) != 1 {
			break
		}
		f := o.generators[0]
		t, ok := f.(Operator)
		if !ok {
			break
		}
		o = t
	}

	head := ctx.Head()
	hasParentheses := head != "" || !ctx.TopParentheses()
	ctx = ctx.ClearHead()
	c1 := ctx.IncParenthesesDepth()
	if hasParentheses {
		c1 = c1.IncDepth()
	}

	sqls := make([]string, len(o.generators))
	values := []interface{}{}
	for i, f := range o.generators {
		var vs []interface{}
		if i == 0 {
			sqls[i], vs = f.ToSQL(c1.ClearHead())
			values = append(values, vs...)
			continue
		}
		sqls[i], vs = f.ToSQL(c1.SetHead(fmt.Sprintf("%s ", o.operator)))
		values = append(values, vs...)
	}
	sql := ctx.Join(sqls...)

	if !hasParentheses {
		return sql, values
	}
	if ctx.Breaking() {
		p := ctx.Prefix()
		return fmt.Sprintf("%s%s(\n%s%s)\n", p, head, sql, p), values
	}
	return fmt.Sprintf("%s(%s)", head, sql), values
}
