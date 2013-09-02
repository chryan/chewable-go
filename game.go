// Copyright (C) 2013 Ryan Chew. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

package cbl

import (
	"time"
)

type Game struct {
	Objects        *ObjectMgr
	preupdatables  []IPreUpdatable
	updatables     []IUpdatable
	postupdatables []IPostUpdatable
	predrawables   []IPreDrawable
	drawables      []IDrawable
	postdrawables  []IPostDrawable
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
	if l := len(g.preupdatables); l > 0 {
		done := make(chan bool)
		for _, val := range g.preupdatables {
			go val.PreUpdate(gt, done)
		}
		defer waitForUpdate(done, l)
	}

	// Draw will happen while updates are happening.
	for _, val := range g.predrawables {
		val.PreDraw(gt)
	}
}

func (g *Game) runStep(gt GameTime) {
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

func (g *Game) postStep(gt GameTime) {
	if l := len(g.postupdatables); l > 0 {
		done := make(chan bool)
		for _, val := range g.postupdatables {
			go val.PostUpdate(gt, done)
		}
		defer waitForUpdate(done, l)
	}

	// Draw will happen while updates are happening.
	for _, val := range g.postdrawables {
		val.PostDraw(gt)
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
		g.tick(t)

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
