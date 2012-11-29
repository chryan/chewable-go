package cbl_test

import (
	"cbl"
	"math/rand"
	"testing"
	"time"
)

type TestComponent struct {
	game  *cbl.Game
	timer int

	init     bool
	shutdown bool
	str string
}

func (t *TestComponent) PreUpdate(time cbl.GameTime, done chan bool) {
	t.timer++
	switch {
	case t.timer >= 60:
		t.str = "a"
		t.game.Exit()
	}
	
	done <- true
}

func (t *TestComponent) Update(gt cbl.GameTime, done chan bool) {
	time.Sleep(time.Duration(rand.Int31()%10) * time.Millisecond)

	switch {
	case t.timer >= 60:
		t.str += "b"
		t.game.Exit()
	}
	done <- true
}

func (t *TestComponent) PostUpdate(time cbl.GameTime, done chan bool) {
	switch {
	case t.timer >= 60:
		t.str += "c"
	}
	done <- true
}

func (t *TestComponent) Draw(gt cbl.GameTime) {
}

func (t *TestComponent) Initialise() {
	t.init = true
}

func (t *TestComponent) Shutdown() {
	t.shutdown = true
}

func TestGameComponents(t *testing.T) {
	g := cbl.NewGame()

	var tc []*TestComponent
	for i := 0; i < 100; i++ {
		newtc := &TestComponent{g, 0, false, false, ""}
		g.AddComponent(newtc)
		tc = append(tc, newtc)
	}

	g.Run()

	for i := 0; i < 100; i++ {
		if tc[i].timer != 60 || !tc[i].init || !tc[i].shutdown || tc[i].str != "abc" {
			t.Logf( "%s", tc[i].str )
			t.FailNow()
		}
	}
}
