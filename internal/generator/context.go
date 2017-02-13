package generator

import "strings"

type Context struct {
	driverName           string
	breaking             bool
	prefix, indent, head string
	depth, bracketDepth  int
}

func NewContext(driverName, prefix, indent string) Context {
	return Context{
		driverName:   strings.ToLower(driverName),
		breaking:     prefix != "" || indent != "",
		prefix:       prefix,
		indent:       indent,
		head:         "",
		depth:        0,
		bracketDepth: 0,
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

func (f Context) ClearBracketDepth() Context {
	f.bracketDepth = 0
	return f
}

func (f Context) IncBracketDepth() Context {
	f.bracketDepth++
	return f
}

func (f Context) TopBracket() bool {
	return f.bracketDepth == 0
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
