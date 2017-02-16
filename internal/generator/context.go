package generator

import "strings"

type Context struct {
	prefix, indent      string
	breaking, flatSets  bool
	head                string
	depth, bracketDepth int
	disablePrefix       bool
}

func newContext(o Options) Context {
	return Context{
		prefix:        o.Prefix,
		indent:        o.Indent,
		breaking:      o.Prefix != "" || o.Indent != "",
		flatSets:      o.FlatSets,
		head:          "",
		depth:         0,
		bracketDepth:  0,
		disablePrefix: false,
	}
}

func (c Context) Head() string {
	return c.head
}

func (c Context) ClearHead() Context {
	c.head = ""
	return c
}

func (c Context) SetHead(head string) Context {
	c.head = head
	return c
}

func (c Context) Breaking() bool {
	return c.breaking
}

func (c Context) Prefix() string {
	if c.disablePrefix {
		c.disablePrefix = false
		return ""
	}
	return c.prefix + strings.Repeat(c.indent, c.depth)
}

func (c Context) DisablePrefix(b bool) Context {
	c.disablePrefix = b
	return c
}

func (c Context) IncDepth() Context {
	c.depth++
	return c
}

func (c Context) ClearParenthesesDepth() Context {
	c.bracketDepth = 0
	return c
}

func (c Context) IncParenthesesDepth() Context {
	c.bracketDepth++
	return c
}

func (c Context) TopParentheses() bool {
	return c.bracketDepth == 0
}

func (c Context) SetFlatSetOperation(flat bool) Context {
	c.flatSets = flat
	return c
}

func (c Context) Join(sqls ...string) string {
	ss := []string{}
	for _, sql := range sqls {
		if sql != "" {
			ss = append(ss, sql)
		}
	}

	if c.Breaking() {
		return strings.Join(ss, "")
	}
	return strings.Join(ss, " ")
}
