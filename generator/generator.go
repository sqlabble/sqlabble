package generator

import (
	"fmt"
	"strings"

	"github.com/minodisk/sqlabble/operator"
)

type Unions struct {
	separator  Expression
	generators []Node
}

func NewUnions(separator Expression, generators ...Node) Unions {
	return Unions{
		separator:  separator,
		generators: generators,
	}
}

func (us Unions) ToSQL(ctx Context) (string, []interface{}) {
	res := []Node{}
	for i, g := range us.generators {
		if needsBracket(ctx, g) {
			g = NewParentheses(g)
		}
		if i == 0 {
			res = append(res, g)
			continue
		}
		res = append(res, us.separator, g)
	}
	return NewNodes(res...).ToSQL(ctx)
}

func needsBracket(ctx Context, generator Node) bool {
	if !ctx.flatSets {
		return true
	}

	gs, ok := generator.(Nodes)
	if !ok {
		return true
	}

	for _, g := range gs {
		if _, ok := g.(Unions); !ok {
			return true
		}
	}

	return false
}

type Node interface {
	ToSQL(Context) (string, []interface{})
}

type Nodes []Node

func NewNodes(generators ...Node) Nodes {
	return Nodes(generators)
}

func (ns Nodes) ToSQL(ctx Context) (string, []interface{}) {
	sqls := make([]string, len(ns))
	values := []interface{}{}
	for i, e := range ns {
		var vs []interface{}
		if i == 0 {
			sqls[i], vs = e.ToSQL(ctx)
			values = append(values, vs...)
		} else {
			sqls[i], vs = e.ToSQL(ctx.ClearHead())
			values = append(values, vs...)
		}
	}
	return ctx.Join(sqls...), values
}

type Expressions []Expression

func NewExpressions(es ...Expression) Expressions {
	return es
}

func (es Expressions) ToSQL(ctx Context) (string, []interface{}) {
	if len(es) == 0 {
		return "", nil
	}
	exp := es[0]
	for i := 1; i < len(es); i++ {
		e := es[i]
		exp = exp.Append(e)
	}
	return exp.ToSQL(ctx)
}

type Expression struct {
	sql    string
	values []interface{}
}

func NewExpression(sql string, values ...interface{}) Expression {
	if len(values) == 0 {
		values = []interface{}{}
	}
	return Expression{
		sql:    sql,
		values: values,
	}
}

func (e Expression) ToSQL(ctx Context) (string, []interface{}) {
	h := ctx.Head()
	ctx = ctx.ClearHead()
	if ctx.Breaking() {
		p := ctx.Prefix()
		return fmt.Sprintf("%s%s%s\n", p, h, e.sql), e.values
	}
	return fmt.Sprintf("%s%s", h, e.sql), e.values
}

func (e Expression) Prepend(exp Expression) Expression {
	e.sql = exp.sql + " " + e.sql
	e.values = append(e.values, exp.values...)
	return e
}

func (e Expression) Append(exp Expression) Expression {
	e.sql = e.sql + " " + exp.sql
	e.values = append(e.values, exp.values...)
	return e
}

func NewArray(es ...Expression) Expression {
	sqls := make([]string, len(es))
	values := []interface{}{}
	for i, e := range es {
		sqls[i] = e.sql
		values = append(values, e.values...)
	}
	return NewExpression(
		fmt.Sprintf("(%s)", strings.Join(sqls, ", ")),
		values...,
	)
}

func NewPlaceholders(values ...interface{}) Expression {
	return NewExpression(
		placeholders(len(values)),
		values...,
	)
}

func placeholders(i int) string {
	s := ""
	for ; i > 0; i-- {
		if i > 1 {
			s += "?, "
			continue
		}
		s += "?"
	}
	return s
}

type Container struct {
	self     Expression
	children Nodes
}

func NewContainer(self Expression, children ...Node) Container {
	return Container{
		self:     self,
		children: children,
	}
}

func (c Container) ToSQL(ctx Context) (string, []interface{}) {
	ps, pvs := c.self.ToSQL(ctx)
	ctx = ctx.ClearHead()
	cs, cvs := c.children.ToSQL(ctx.IncDepth())
	return ctx.Join(ps, cs), append(pvs, cvs...)
}

func (c Container) AddChild(children ...Node) Container {
	c.children = append(c.children, children...)
	return c
}

type Join struct {
	sep        string
	generators []Node
}

// func NewJoin(sep string, generators ...Node) Join {
// 	return Join{
// 		sep:        sep,
// 		generators: generators,
// 	}
// }
//
// func (j Join) ToSQL(ctx Context) (string, []interface{}) {
// 	sqls := make([]string, len(j.generators))
// 	values := []interface{}{}
// 	for i, g := range j.generators {
// 		var vs []interface{}
// 		sqls[i], vs = g.ToSQL(ctx)
// 		values = append(values, vs...)
// 	}
// 	return NewExpression(strings.Join(sqls, j.sep), values...).ToSQL(ctx)
// }

type Comma struct {
	generators []Node
}

func NewComma(generators ...Node) Comma {
	return Comma{
		generators: generators,
	}
}

func (c Comma) ToSQL(ctx Context) (string, []interface{}) {
	sqls := make([]string, len(c.generators))
	values := []interface{}{}
	for i, t := range c.generators {
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

type Not struct {
	generator Node
}

func NewNot(generator Node) Not {
	return Not{
		generator: generator,
	}
}

func (n Not) ToSQL(ctx Context) (string, []interface{}) {
	// This is a code to delete duplicate NOTs,
	// but optimizing is not a task of SQL builder,
	// so comment it out.
	// if t, ok := n.operator.(Not); ok {
	// 	return t.operator.Format(c)
	// }

	op := "NOT "

	head := ctx.Head()
	ctx = ctx.ClearParenthesesDepth()

	sql, values := n.generator.ToSQL(ctx.IncDepth().ClearHead())

	if ctx.Breaking() {
		p := ctx.Prefix()
		return fmt.Sprintf("%s%s%s(\n%s%s)\n", p, head, op, sql, p), values
	}

	return fmt.Sprintf("%s%s(%s)", head, op, sql), values
}

type Parentheses struct {
	generator Node
}

func NewParentheses(g Node) Parentheses {
	return Parentheses{
		generator: g,
	}
}

func (b Parentheses) ToSQL(ctx Context) (string, []interface{}) {
	head := ctx.Head()
	ctx = ctx.ClearParenthesesDepth()

	sql, values := b.generator.ToSQL(ctx.IncDepth().ClearHead())

	if ctx.Breaking() {
		p := ctx.Prefix()
		return fmt.Sprintf("%s%s(\n%s%s)\n", p, head, sql, p), values
	}

	return fmt.Sprintf("%s(%s)", head, sql), values
}
