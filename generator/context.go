package generator

import "strings"

type Context struct {
	breaking             bool
	prefix, indent, head string
	depth, bracketDepth  int
	flatSetOperation     bool
}

func NewContext(prefix, indent string) Context {
	return Context{
		breaking:         prefix != "" || indent != "",
		prefix:           prefix,
		indent:           indent,
		head:             "",
		depth:            0,
		bracketDepth:     0,
		flatSetOperation: false,
	}
}

func (f Context) Head() string {
	return f.head
}

func (f Context) ClearHead() Context {
	f.head = ""
	return f
}

func (f Context) SetHead(head string) Context {
	f.head = head
	return f
}

func (f Context) Breaking() bool {
	return f.breaking
}

func (f Context) Prefix() string {
	return f.prefix + strings.Repeat(f.indent, f.depth)
}

func (f Context) IncDepth() Context {
	f.depth++
	return f
}

func (f Context) ClearParenthesesDepth() Context {
	f.bracketDepth = 0
	return f
}

func (f Context) IncParenthesesDepth() Context {
	f.bracketDepth++
	return f
}

func (f Context) TopParentheses() bool {
	return f.bracketDepth == 0
}

func (c Context) SetFlatSetOperation(flat bool) Context {
	c.flatSetOperation = flat
	return c
}

func (f Context) Join(sqls ...string) string {
	ss := []string{}
	for _, sql := range sqls {
		if sql != "" {
			ss = append(ss, sql)
		}
	}

	if f.Breaking() {
		return strings.Join(ss, "")
	}
	return strings.Join(ss, " ")
}
