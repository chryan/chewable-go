package cbl

type IUpdatable interface {
	PreUpdate(gt GameTime, done chan bool)
	Update(gt GameTime, done chan bool)
	PostUpdate(gt GameTime, done chan bool)
}

type IDrawable interface {
	Draw(gt GameTime)
}

type IGameComponent interface {
	Initialise()
	Shutdown()
}
