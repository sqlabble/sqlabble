package generator

type Container struct {
	self     Node
	children ParallelNodes
}

func NewContainer(self Node, children ...Node) Container {
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
