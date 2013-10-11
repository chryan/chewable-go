// Copyright (C) 2013 Ryan Chew. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

package cbl

import (
	"testing"
)

type TestComponent struct {
	game  *Game
	timer int

	init     bool
	shutdown bool
	str string
}

func (t *TestComponent) PreUpdate(time GameTime) {
	t.timer++
	switch {
	case t.timer >= 60:
		t.str = "a"
		t.game.Exit()
	}
}

func (t *TestComponent) Update(gt GameTime) {
	switch {
	case t.timer >= 60:
		t.str += "b"
		t.game.Exit()
	}
}

func (t *TestComponent) PostUpdate(time GameTime) {
	switch {
	case t.timer >= 60:
		t.str += "c"
	}
}

func (t *TestComponent) Draw(gt GameTime) {
}

func (t *TestComponent) Initialise() {
	t.init = true
}

func (t *TestComponent) Shutdown() {
	t.shutdown = true
}

func TestGameComponents(t *testing.T) {
	g := NewGame()

	var tc []*TestComponent
	for i := 0; i < 100; i++ {
		newtc := &TestComponent{g, 0, false, false, ""}
		g.AddComponent(newtc)
		tc = append(tc, newtc)
		if tc[i].timer != 0 || tc[i].init || tc[i].shutdown {
			t.Fatalf("Component initialise failed: %s", tc[i])
		}
	}

	g.Run()

	for i := 0; i < 100; i++ {
		if tc[i].timer != 60 || !tc[i].init || !tc[i].shutdown || tc[i].str != "abc" {
			t.Fatalf( "Component shutdown failed: %s", tc[i] )
		}
	}
}
