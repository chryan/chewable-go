package cbl

type ComponentList struct {
	components     []IComponent
	preupdatables  []IPreUpdatable
	updatables     []IUpdatable
	postupdatables []IPostUpdatable
	predrawables   []IPreDrawable
	drawables      []IDrawable
	postdrawables  []IPostDrawable
}

func NewComponentList() *ComponentList {
	return &ComponentList{
		components: make([]IComponent, 8)[0:0],

		preupdatables:  make([]IPreUpdatable, 8)[0:0],
		updatables:     make([]IUpdatable, 8)[0:0],
		postupdatables: make([]IPostUpdatable, 8)[0:0],

		predrawables:  make([]IPreDrawable, 8)[0:0],
		drawables:     make([]IDrawable, 8)[0:0],
		postdrawables: make([]IPostDrawable, 8)[0:0],
	}
}

func (c *ComponentList) Add(component IComponent) IComponent {
	c.components = append(c.components, component)

	if u, ok := component.(IPreUpdatable); ok {
		c.preupdatables = append(c.preupdatables, u)
	}
	if u, ok := component.(IUpdatable); ok {
		c.updatables = append(c.updatables, u)
	}
	if u, ok := component.(IPostUpdatable); ok {
		c.postupdatables = append(c.postupdatables, u)
	}
	if dr, ok := component.(IPreDrawable); ok {
		c.predrawables = append(c.predrawables, dr)
	}
	if dr, ok := component.(IDrawable); ok {
		c.drawables = append(c.drawables, dr)
	}
	if dr, ok := component.(IPostDrawable); ok {
		c.postdrawables = append(c.postdrawables, dr)
	}

	return component
}

func (c *ComponentList) Initialise() {
	for _, component := range c.components {
		component.Initialise()
	}
}

func (c *ComponentList) Shutdown() {
	for i := len(c.components); i > 0; i-- {
		c.components[i-1].Shutdown()
	}
}
