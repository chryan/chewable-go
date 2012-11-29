package cbl

type IPreUpdatable interface {
	PreUpdate(gt GameTime, done chan bool)
}

type IUpdatable interface {
	Update(gt GameTime, done chan bool)
}

type IPostUpdatable interface {
	PostUpdate(gt GameTime, done chan bool)
}

type IDrawable interface {
	Draw(gt GameTime)
}

type IComponent interface {
	Initialise()
	Shutdown()
}
