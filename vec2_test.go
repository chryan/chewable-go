// Copyright (C) 2013 Ryan Chew. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

package cbl

import (
	"testing"

	"math"
)

func TestVec2(t *testing.T) {
	var vec1 = Vec2f{1, 2}
	var vec2 = Vec2f{6, 6}

	if dot := vec1.Dot(vec2); dot != vec1.X*vec2.X+vec1.Y*vec2.Y {
		t.Fatalf("Dot product check failed: %v", dot)
	}

	if l := float32(math.Sqrt(float64(vec1.X*vec1.X + vec1.Y*vec1.Y))); failNear(vec1.Len(), l) {
		t.Fatalf("Length check failed: %v", l)
	} else if norm := vec1.Norm(); failNear(norm.X, vec1.X/l) || failNear(norm.Y, vec1.Y/l) {
		t.Fatalf("Normalize check failed: %v", norm)
	}

	test1 := vec1.Add(vec1).Sub(vec1).Add(vec1).Add(vec1).Mul(100).Div(10)

	if failNear(test1.X, 30) || failNear(test1.Y, 60.0) {
		t.Fatalf("Add/Sub/Mul/Div check failed: %v", test1)
	}

	test2 := vec2
	test1 = test2
	test2.X, test2.Y = 10, 20
	test1, test2 = test2, test1

	if failNear(test2.X, 6) || failNear(test2.Y, 6) {
		t.Fatalf("test2 swap failed: %v", test2)
	}
	if failNear(test1.X, 10) || failNear(test1.Y, 20) {
		t.Fatalf("test1 swap failed: %v", test1)
	}

	if dist1, dist2 := test1.Dist(test2), float32(math.Sqrt(4*4+14*14)); failNear(dist1, dist2) {
		t.Fatalf("Distance check failed: %v : %v", dist1, dist2)
	}
}
