// Copyright (C) 2013 Ryan Chew. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

package cbl

import (
	"reflect"
	"testing"
)

type TestType struct {
	TestValue  int32
	NameValue  string
	InnerValue InnerType
}

type InnerType struct {
	Value float32
}

func TestTypeFactory(t *testing.T) {
	defer TypeFactory.Clear()
	TypeFactory.
		Add(reflect.TypeOf(new(TestType))).
		Add(reflect.TypeOf(new(InnerType)))

	if newtype := TypeFactory.New("TestType", "cbl"); newtype != nil {
		if _, ok := newtype.(TestType); !ok {
			t.Errorf("Unable to create type: TestType")
		}
	}

	if newtype := TypeFactory.New("InnerType", "cbl"); newtype != nil {
		if _, ok := newtype.(InnerType); !ok {
			t.Errorf("Unable to create type: InnerType")
		}
	}

	if gettype := TypeFactory.Get("TestType", "cbl"); gettype == nil {
		t.Errorf("Unable to get type: TestType")
	}

	if gettype := TypeFactory.Get("InnerType", "cbl"); gettype == nil {
		t.Errorf("Unable to get type: InnerType")
	}
}
