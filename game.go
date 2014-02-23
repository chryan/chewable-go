// Copyright (C) 2013 Ryan Chew. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

package cbl

import (
	"time"
)

type Game struct {
	Objects    *ObjectMgr
	Components *ComponentList
	Services   *Services
	FixedStep  time.Duration
	DropFrames uint
	running    bool
}

// Allocation
func NewGame() *Game {
	return NewGameWithServices(NewServices())
}

func NewGameWithServices(services *Services) *Game {
	return &Game{
		Objects:    NewObjectMgr(),
		Components: NewComponentList(),
		Services:   services,
		FixedStep:  time.Second / 60.0,
		DropFrames: 5,
	}
}

func (g *Game) preStep(gt GameTime) {
	for _, u := range g.Components.preupdatables {
		u.PreUpdate(gt)
	}
	for _, d := range g.Components.predrawables {
		d.PreDraw(gt)
	}
}

func (g *Game) runStep(gt GameTime) {
	for _, u := range g.Components.updatables {
		u.Update(gt)
	}
	for _, val := range g.Components.drawables {
		val.Draw(gt)
	}
}

func (g *Game) postStep(gt GameTime) {
	for _, u := range g.Components.postupdatables {
		u.PostUpdate(gt)
	}
	for _, d := range g.Components.postdrawables {
		d.PostDraw(gt)
	}
}

// Start the game loop.
func (g *Game) Run() {
	g.initialise()

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

					if slow >= g.DropFrames {
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
	g.Components.Initialise()
}

func (g *Game) shutdown() {
	g.Components.Shutdown()
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
