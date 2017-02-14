package generator

import (
	"fmt"
	"strings"

	"github.com/minodisk/sqlabble/internal/operator"
)

type Unions struct {
	separator  Expression
	generators []Generator
}

func NewUnions(separator Expression, generators ...Generator) Unions {
	return Unions{
		separator:  separator,
		generators: generators,
	}
}

func (us Unions) Generate(ctx Context) (string, []interface{}) {
	res := []Generator{}
	for i, g := range us.generators {
		if needsBracket(ctx, g) {
			g = NewBracket(g)
		}
		if i == 0 {
			res = append(res, g)
			continue
		}
		res = append(res, us.separator, g)
	}
	return NewGenerators(res...).Generate(ctx)
}

func needsBracket(ctx Context, generator Generator) bool {
	if !ctx.flatSetOperation {
		return true
	}

	gs, ok := generator.(Generators)
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

type Expressions []Expression

func NewExpressions(es ...Expression) Expressions {
	return es
}

func (es Expressions) Generate(ctx Context) (string, []interface{}) {
	if len(es) == 0 {
		return "", nil
	}
	exp := es[0]
	for i := 1; i < len(es); i++ {
		e := es[i]
		exp = exp.Append(e)
	}
	return exp.Generate(ctx)
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
		sqls[i], vs = g.Generate(ctx)
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
