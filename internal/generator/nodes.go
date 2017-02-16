package generator

// Nodes is list of Node.
type Nodes []Node

// NewNodes return a new Nodes excluding nil.
func NewNodes(nodes ...Node) Nodes {
	ns := []Node{}
	for _, n := range nodes {
		if n == nil {
			continue
		}
		ns = append(ns, n)
	}
	return Nodes(ns)
}

// ToSQL returns a query and a slice of values.
func (ns Nodes) ToSQL(ctx Context) (string, []interface{}) {
	sqls := make([]string, len(ns))
	values := []interface{}{}
	for i, e := range ns {
		var vs []interface{}
		if i == 0 {
			sqls[i], vs = e.ToSQL(ctx)
			values = append(values, vs...)
		} else {
			sqls[i], vs = e.ToSQL(ctx.clearHead())
			values = append(values, vs...)
		}
	}
	return ctx.join(sqls...), values
}
