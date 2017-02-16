package generator

// Container is a Node that contains a parental Node and children Nodes.
type Container struct {
	self     Node
	children Nodes
}

// NewContainer returns a new Container.
func NewContainer(self Node, children ...Node) Container {
	return Container{
		self:     self,
		children: children,
	}
}

// ToSQL returns a query and a slice of values.
func (c Container) ToSQL(ctx Context) (string, []interface{}) {
	ps, pvs := c.self.ToSQL(ctx)
	ctx = ctx.clearHead()
	cs, cvs := c.children.ToSQL(ctx.incDepth())
	return ctx.join(ps, cs), append(pvs, cvs...)
}

// AddChild add children to c.
func (c Container) AddChild(children ...Node) Container {
	c.children = append(c.children, children...)
	return c
}
