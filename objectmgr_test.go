// Copyright (C) 2013 Ryan Chew. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

package cbl

import (
	"testing"
)

type TestOC int

func (t *TestOC) Initialise() {
}

func (t *TestOC) Shutdown() {
}

func TestObjectMgr(t *testing.T) {
	om := NewObjectMgr()
	obj := om.New("NewObject")

	// Test that the object is valid.
	if obj == nil || om.Get("NewObject") == nil {
		t.Fail()
	}

	// Delete and purge objects.
	om.Del("NewObject")
	if om.Get("NewObject") == nil {
		t.Fail()
	}
	om.Purge()

	// Ensure object is gone.
	if om.Get("NewObject") != nil {
		t.Fail()
	}
}

func TestObjectComponent(t *testing.T) {
	om := NewObjectMgr()
	obj := om.New("NewObject")

	obj.Components.Add("TestOC", new(TestOC))

	if testoc := obj.Components.Get("TestOC"); testoc == nil {
		t.Fail()
	} else {
		if val, ok := testoc.(*TestOC); !ok {
			t.Fail()
		} else {
			*val = 10
		}
	}

	if testoc := obj.Components.Get("TestOC"); testoc == nil {
		t.Fail()
	} else {
		if val, ok := testoc.(*TestOC); !ok || *val != 10 {
			t.Fail()
		}
	}
}
