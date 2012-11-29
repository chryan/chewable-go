package cbl

import (
	"time"
)

type Game struct {
	Objects        *ObjectMgr
	preupdatables  []IPreUpdatable
	updatables     []IUpdatable
	postupdatables []IPostUpdatable
	drawables      []IDrawable
	components     []IComponent
	running        bool

	FixedStep time.Duration
}

// Allocation
func NewGame() *Game {
	g := new(Game)
	g.preupdatables = make([]IPreUpdatable, 8)[0:0]
	g.updatables = make([]IUpdatable, 8)[0:0]
	g.postupdatables = make([]IPostUpdatable, 8)[0:0]
	g.drawables = make([]IDrawable, 8)[0:0]
	g.components = make([]IComponent, 8)[0:0]
	g.Objects = NewObjectMgr()
	return g
}

// Public
func (g *Game) AddComponent(gc IComponent) {
	g.components = append(g.components, gc)

	if u, ok := gc.(IPreUpdatable); ok {
		g.preupdatables = append(g.preupdatables, u)
	}

	if u, ok := gc.(IUpdatable); ok {
		g.updatables = append(g.updatables, u)
	}

	if u, ok := gc.(IPostUpdatable); ok {
		g.postupdatables = append(g.postupdatables, u)
	}

	if dr, ok := gc.(IDrawable); ok {
		g.drawables = append(g.drawables, dr)
	}
}

func (g *Game) Tick(gt GameTime) {
	g.Objects.Purge()
	g.preUpdate(gt)
	g.updateAndDraw(gt)
	g.postUpdate(gt)
}

func (g *Game) preUpdate(gt GameTime) {
	if l := len(g.preupdatables); l > 0 {
		done := make(chan bool)
		for _, val := range g.preupdatables {
			go val.PreUpdate(gt, done)
		}

		defer waitForUpdate(done, l)
	}
}

func (g *Game) updateAndDraw(gt GameTime) {
	if l := len(g.updatables); l > 0 {
		done := make(chan bool)
		for _, val := range g.updatables {
			go val.Update(gt, done)
		}
		defer waitForUpdate(done, l)
	}

	// Draw will happen while updates are happening.
	for _, val := range g.drawables {
		val.Draw(gt)
	}
}

func (g *Game) postUpdate(gt GameTime) {
	if l := len(g.postupdatables); l > 0 {
		done := make(chan bool)
		for _, val := range g.postupdatables {
			go val.PostUpdate(gt, done)
		}
		defer waitForUpdate(done, l)
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

		t := GameTime{elapsed, total}
		g.Tick(t)

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

func waitForUpdate(done chan bool, numobjs int) {
	for i := 0; i < numobjs; i++ {
		<-done
	}
}
