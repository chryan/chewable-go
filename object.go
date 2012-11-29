package cbl

type objectcomponents map[string]IComponent

type Object struct {
	Components objectcomponents
}

func NewObject() *Object {
	return &Object{Components: make(objectcomponents)}
}

func (o *Object) Initialise() {
	for _, comp := range o.Components {
		comp.Initialise()
	}
}

func (o *Object) Shutdown() {
	for _, comp := range o.Components {
		comp.Shutdown()
	}
}

func (oc *objectcomponents) Add(name string, comp IComponent) IComponent {
	if c, ok := (*oc)[name]; ok {
		// Log here.
		return c
	}

	(*oc)[name] = comp
	return comp
}

func (oc *objectcomponents) Get(name string) IComponent {
	return (*oc)[name]
}
