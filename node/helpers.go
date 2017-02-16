package node

import "strings"

// JoinExpressions joins expressions into a Expression.
func JoinExpressions(expressions ...Expression) Expression {
	var exp Expression
	for i, e := range expressions {
		if i == 0 {
			exp = e
			continue
		}
		exp = exp.Append(e)
	}
	return exp
}

// ArrayToExpression joins expressions into a Expression with commas,
// enclosed in parentheses and returns it.
func ArrayToExpression(es ...Expression) Expression {
	l := len(es)
	if l == 0 {
		return EmptyExpression
	}
	sqls := make([]string, l)
	values := []interface{}{}
	for i, e := range es {
		sqls[i] = e.sql
		values = append(values, e.values...)
	}
	return NewExpression(
		strings.Join(sqls, ", "),
		values...,
	).WrapSQL("(", ")")
}

// ValuesToExpression returns a Expression representing the values.
func ValuesToExpression(values ...interface{}) Expression {
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

func needsParentheses(ctx Context, node Node) bool {
	if !ctx.flatSets {
		return true
	}

	gs, ok := node.(Nodes)
	if !ok {
		return true
	}

	for _, g := range gs {
		if _, ok := g.(Set); !ok {
			return true
		}
	}

	return false
}
