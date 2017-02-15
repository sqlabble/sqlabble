package generator

type ParallelNodes []Node

func NewParallelNodes(nodes ...Node) ParallelNodes {
	ns := []Node{}
	for _, n := range nodes {
		if n == nil {
			continue
		}
		ns = append(ns, n)
	}
	return ParallelNodes(ns)
}

func (ns ParallelNodes) ToSQL(ctx Context) (string, []interface{}) {
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

// type SerialNodes []Node
//
// func NewSerialNodes(nodes ...Node) SerialNodes {
// 	return SerialNodes(nodes)
// }
//
// func (ns SerialNodes) ToSQL(ctx Context) (string, []interface{}) {
// 	sqls := make([]string, len(ns))
// 	values := []interface{}{}
// 	for i, n := range ns {
// 		var vs []interface{}
// 		if i == 0 {
// 			sqls[i], vs = n.ToSQL(ctx)
// 			values = append(values, vs...)
// 		} else {
// 			sqls[i], vs = n.ToSQL(ctx.DisablePrefix(true))
// 			values = append(values, vs...)
// 		}
// 	}
// 	return strings.Join(sqls, " "), values
// }
