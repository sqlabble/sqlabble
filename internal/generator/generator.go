package generator

import (
	"fmt"
	"strings"

	"github.com/minodisk/sqlabble/internal/grammar/operator"
)

type Generator interface {
	Generate(Context) (string, []interface{})
}

type Generators []Generator

func NewGenerators(generators ...Generator) Generators {
	return Generators(generators)
}

func (fs Generators) Generate(ctx Context) (string, []interface{}) {
	sqls := make([]string, len(fs))
	values := []interface{}{}
	for i, e := range fs {
		var vs []interface{}
		if i == 0 {
			sqls[i], vs = e.Generate(ctx)
			values = append(values, vs...)
		} else {
			sqls[i], vs = e.Generate(ctx.ClearHead())
			values = append(values, vs...)
		}
	}
	return ctx.Join(sqls...), values
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

func (e Expression) Generate(ctx Context) (string, []interface{}) {
	h := ctx.Head()
	ctx = ctx.ClearHead()
	if ctx.Breaking() {
		p := ctx.Prefix()
		return fmt.Sprintf("%s%s%s\n", p, h, e.sql), e.values
	}
	return fmt.Sprintf("%s%s", h, e.sql), e.values
}

func (e Expression) Prepend(sql string, values ...interface{}) Expression {
	if len(values) == 0 {
		values = []interface{}{}
	}
	e.sql = sql + e.sql
	e.values = append(values, e.values...)
	return e
}

func (e Expression) Append(sql string, values ...interface{}) Expression {
	if len(values) == 0 {
		values = []interface{}{}
	}
	e.sql = e.sql + sql
	e.values = append(e.values, values...)
	return e
}

type Container struct {
	self     Expression
	children Generators
}

func NewContainer(self Expression, children ...Generator) Container {
	return Container{
		self:     self,
		children: children,
	}
}

func (c Container) Generate(ctx Context) (string, []interface{}) {
	ps, pvs := c.self.Generate(ctx)
	ctx = ctx.ClearHead()
	cs, cvs := c.children.Generate(ctx.IncDepth())
	return ctx.Join(ps, cs), append(pvs, cvs...)
}

func (c Container) AddChild(children ...Generator) Container {
	c.children = append(c.children, children...)
	return c
}

type Join struct {
	sep        string
	generators []Generator
}

func NewJoin(sep string, generators ...Generator) Join {
	return Join{
		sep:        sep,
		generators: generators,
	}
}

func (j Join) Generate(ctx Context) (string, []interface{}) {
	sqls := make([]string, len(j.generators))
	values := []interface{}{}
	for i, g := range j.generators {
		var vs []interface{}
		sqls[i], vs = g.Generate(NonBreakingContext)
		values = append(values, vs...)
	}
	return NewExpression(strings.Join(sqls, j.sep), values...).Generate(ctx)
}

type Comma struct {
	generators []Generator
}

func NewComma(generators ...Generator) Comma {
	return Comma{
		generators: generators,
	}
}

func (c Comma) Generate(ctx Context) (string, []interface{}) {
	sqls := make([]string, len(c.generators))
	values := []interface{}{}
	for i, t := range c.generators {
		var vs []interface{}
		if i == 0 {
			sqls[i], vs = t.Generate(ctx)
			values = append(values, vs...)
			continue
		}
		sqls[i], vs = t.Generate(ctx.SetHead(", "))
		values = append(values, vs...)
	}
	return strings.Join(sqls, ""), values
}

type Operator struct {
	operator   operator.Operator
	generators []Generator
}

func NewOperator(operator operator.Operator, generators ...Generator) Operator {
	return Operator{
		operator:   operator,
		generators: generators,
	}
}

func (o Operator) Generate(ctx Context) (string, []interface{}) {
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
	hasBracket := head != "" || !ctx.TopBracket()
	ctx = ctx.ClearHead()
	c1 := ctx.IncBracketDepth()
	if hasBracket {
		c1 = c1.IncDepth()
	}

	sqls := make([]string, len(o.generators))
	values := []interface{}{}
	for i, f := range o.generators {
		var vs []interface{}
		if i == 0 {
			sqls[i], vs = f.Generate(c1.ClearHead())
			values = append(values, vs...)
			continue
		}
		sqls[i], vs = f.Generate(c1.SetHead(fmt.Sprintf("%s ", o.operator)))
		values = append(values, vs...)
	}
	sql := ctx.Join(sqls...)

	if !hasBracket {
		return sql, values
	}
	if ctx.Breaking() {
		p := ctx.Prefix()
		return fmt.Sprintf("%s%s(\n%s%s)\n", p, head, sql, p), values
	}
	return fmt.Sprintf("%s(%s)", head, sql), values
}

type Not struct {
	generator Generator
}

func NewNot(generator Generator) Not {
	return Not{
		generator: generator,
	}
}

func (n Not) Generate(ctx Context) (string, []interface{}) {
	// This is a code to delete duplicate NOTs,
	// but optimizing is not a task of SQL builder,
	// so comment it out.
	// if t, ok := n.operator.(Not); ok {
	// 	return t.operator.Format(c)
	// }

	op := "NOT "

	head := ctx.Head()
	ctx = ctx.ClearBracketDepth()

	sql, values := n.generator.Generate(ctx.IncDepth().ClearHead())

	if ctx.Breaking() {
		p := ctx.Prefix()
		return fmt.Sprintf("%s%s%s(\n%s%s)\n", p, head, op, sql, p), values
	}

	return fmt.Sprintf("%s%s(%s)", head, op, sql), values
}

type Bracket struct {
	generator Generator
}

func NewBracket(g Generator) Bracket {
	return Bracket{
		generator: g,
	}
}

func (b Bracket) Generate(ctx Context) (string, []interface{}) {
	head := ctx.Head()
	ctx = ctx.ClearBracketDepth()

	sql, values := b.generator.Generate(ctx.IncDepth().ClearHead())

	if ctx.Breaking() {
		p := ctx.Prefix()
		return fmt.Sprintf("%s%s(\n%s%s)\n", p, head, sql, p), values
	}

	return fmt.Sprintf("%s(%s)", head, sql), values
}
