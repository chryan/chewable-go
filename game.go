package cbl

import (
	"time"
)

type Game struct {
	updatables [] IUpdatable
	drawables [] IDrawable
	components [] IGameComponent	
	running bool

	FixedStep time.Duration
	lock <- chan int
}

// Allocation
func NewGame() *Game {
	g := new(Game)
	g.updatables = make([]IUpdatable, 8)[0:0]
	g.drawables = make([]IDrawable, 8)[0:0]
	g.components = make([]IGameComponent, 8)[0:0]
	
	g.lock = make(chan int)
	return g
}

// Public
func (g *Game) AddComponent(gc IGameComponent) {
	g.components = append(g.components, gc)
	
	if up, ok := gc.(IUpdatable); ok {
		g.updatables = append(g.updatables, up)
	}
	if dr, ok := gc.(IDrawable); ok {
		g.drawables = append(g.drawables, dr)
	}
	
	//if g.running {
	//	gc.Initialise()
	//}
}

func waitForUpdate(done chan bool, numobjs int) {
	for i := 0; i < numobjs; i++ {
		<- done
	}
}

func (g *Game) Update(gt GameTime) {
	_loop := len(g.updatables)
	
	
	done := make(chan bool, _loop)
	for _, val := range g.updatables {
		go val.PreUpdate(gt, done)
	}
	waitForUpdate(done, _loop)

	done = make(chan bool, _loop)
	for _, val := range g.updatables {
		go val.Update(gt, done)
	}

	waitForUpdate(done, _loop)
	
	done = make(chan bool, _loop)
	for _, val := range g.updatables {
		go val.PostUpdate(gt, done)
	}
	waitForUpdate(done, _loop)
}

func (g *Game) Draw(gt GameTime) {
	for _, val := range g.drawables {
		val.Draw(gt)
	}
}

func (g *Game) Run() {
	g.initialise()

	ticker := time.Tick(time.Second / 60.0)
	
	var total time.Duration
	prev := time.Now()

	g.running = true
	for now := range ticker {
		elapsed := now.Sub(prev)
		total += elapsed
		prev = now

		t := GameTime{elapsed,total}
		g.Update(t)
		g.Draw(t)
		
		if !g.running {
			break
		}
	}
	
	g.shutdown()
}

func (g *Game) Exit() {
	g.running = false
}

func (g *Game) initialise() {
	for _, gc := range g.components {
		gc.Initialise()
	}
}

func (g *Game) shutdown() {
	for _, gc := range g.components {
		gc.Shutdown()
	}
}