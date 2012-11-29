package cbl_test

import (
	"cbl"
	"math"
	"testing"
)

const (
	VECX = 1
	VECY = 2
)

func TestVec2(t *testing.T) {
	vec := cbl.Vec2f{VECX, VECY}

	if vec.Dot(vec) != 5 {
		t.FailNow()
	}

	if l := math.Sqrt(VECX*VECX + VECY*VECY); vec.Len() != l {
		t.FailNow()
	} else if norm := vec.Norm(); norm.X != VECX/l || norm.Y != VECY/l {
		t.FailNow()
	}

	test1 := vec.Add(vec).Sub(vec).Add(vec).Add(vec).Mul(100).Div(10)

	if test1.X != 30 || test1.Y != 60 {
		t.FailNow()
	}

	test2 := cbl.Vec2f{6, 6}

	test1 = test2
	test2.X, test2.Y = 10, 20

	test1, test2 = test2, test1

	if test2.X != 6 || test2.Y != 6 {
		t.FailNow()
	}
	if test1.X != 10 || test1.Y != 20 {
		t.FailNow()
	}
	if test1.Dist(test2) != math.Sqrt(4*4+14*14) {
		t.FailNow()
	}
}
