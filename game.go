// Copyright (C) 2013 Ryan Chew. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

package cbl

import (
	"time"
)

type Game struct {
	Objects    *ObjectMgr
	Services   *Services
	FixedStep  time.Duration
	DropFrames uint

	preupdatables  []IPreUpdatable
	updatables     []IUpdatable
	postupdatables []IPostUpdatable
	predrawables   []IPreDrawable
	drawables      []IDrawable
	postdrawables  []IPostDrawable
	components     []IComponent
	running        bool
}

// Allocation
func NewGame() *Game {
	return &Game{
		Objects:    NewObjectMgr(),
		Services:   NewServices(),
		FixedStep:  time.Second / 60.0,
		DropFrames: 5,

		preupdatables:  make([]IPreUpdatable, 8)[0:0],
		updatables:     make([]IUpdatable, 8)[0:0],
		postupdatables: make([]IPostUpdatable, 8)[0:0],

		predrawables:  make([]IPreDrawable, 8)[0:0],
		drawables:     make([]IDrawable, 8)[0:0],
		postdrawables: make([]IPostDrawable, 8)[0:0],

		components: make([]IComponent, 8)[0:0],
	}
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
	if dr, ok := gc.(IPreDrawable); ok {
		g.predrawables = append(g.predrawables, dr)
	}
	if dr, ok := gc.(IDrawable); ok {
		g.drawables = append(g.drawables, dr)
	}
	if dr, ok := gc.(IPostDrawable); ok {
		g.postdrawables = append(g.postdrawables, dr)
	}
}

func (g *Game) preStep(gt GameTime) {
	for _, u := range g.preupdatables {
		u.PreUpdate(gt)
	}
	for _, d := range g.predrawables {
		d.PreDraw(gt)
	}
}

func (g *Game) runStep(gt GameTime) {
	for _, u := range g.updatables {
		u.Update(gt)
	}
	for _, val := range g.drawables {
		val.Draw(gt)
	}
}

func (g *Game) postStep(gt GameTime) {
	for _, u := range g.postupdatables {
		u.PostUpdate(gt)
	}
	for _, d := range g.postdrawables {
		d.PostDraw(gt)
	}
}

// Start the game loop.
func (g *Game) Run() {
	g.initialise()

	//tick := time.Nanosecond * 10

	var total, elapsed time.Duration
	prev := time.Now()

	g.running = true

	for g.running {
		now := time.Now()
		since := now.Sub(prev)

		if since > 0 {
			elapsed += since
			prev = now

			if elapsed > g.FixedStep {
				totalreal := total + elapsed
				elapsedreal := elapsed
				for slow := uint(0); elapsed > g.FixedStep; slow++ {
					total += g.FixedStep
					elapsed -= g.FixedStep

					t := GameTime{g.FixedStep, total, elapsedreal, totalreal, slow > 0}
					g.tick(t)

					if g.DropFrames > slow {
						elapsed = 0
						break
					}
				}
			}
		}
	}

	g.shutdown()
}

// Trigger the game to exit.
func (g *Game) Exit() {
	g.running = false
}

func (g *Game) initialise() {
	for _, gc := range g.components {
		gc.Initialise()
	}
}

func (g *Game) shutdown() {
	for i := len(g.components); i > 0; i-- {
		g.components[i-1].Shutdown()
	}
}

func (g *Game) tick(gt GameTime) {
	g.Objects.Purge()
	g.preStep(gt)
	g.runStep(gt)
	g.postStep(gt)
}

func waitForUpdate(done chan bool, numobjs int) {
	for i := 0; i < numobjs; i++ {
		<-done
	}
}
