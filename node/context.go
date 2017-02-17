package node

import "strings"

var (
	Standard         = NewContext("", "", `"`, false)
	IndentedStandard = NewContext("", "  ", `"`, false)
	MySQL            = NewContext("", "", "`", false)
	MySQL4           = NewContext("", "", "`", true)
)

// Context is a container for storing the state to be output
// in the process of building a query.
type Context struct {
	prefix, indent, Quote string
	breaking, flatSets    bool
	head                  string
	depth, bracketDepth   int
}

func NewContext(prefix, indent, quote string, flatSets bool) Context {
	if quote == "" {
		quote = `"`
	}
	return Context{
		prefix:       prefix,
		indent:       indent,
		breaking:     prefix != "" || indent != "",
		Quote:        quote,
		flatSets:     flatSets,
		head:         "",
		depth:        0,
		bracketDepth: 0,
	}
}

func (c Context) CurrentHead() string {
	return c.head
}

func (c Context) ClearHead() Context {
	c.head = ""
	return c
}

func (c Context) setHead(head string) Context {
	c.head = head
	return c
}

func (c Context) IsBreaking() bool {
	return c.breaking
}

func (c Context) Prefix() string {
	return c.prefix + strings.Repeat(c.indent, c.depth)
}

func (c Context) incDepth() Context {
	c.depth++
	return c
}

func (c Context) clearParenthesesDepth() Context {
	c.bracketDepth = 0
	return c
}

func (c Context) incParenthesesDepth() Context {
	c.bracketDepth++
	return c
}

func (c Context) isTopParentheses() bool {
	return c.bracketDepth == 0
}

func (c Context) setFlatSet(flat bool) Context {
	c.flatSets = flat
	return c
}

func (c Context) join(sqls ...string) string {
	ss := []string{}
	for _, sql := range sqls {
		if sql != "" {
			ss = append(ss, sql)
		}
	}

	if c.IsBreaking() {
		return strings.Join(ss, "")
	}
	return strings.Join(ss, " ")
}
