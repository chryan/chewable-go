// Copyright (C) 2013 Ryan Chew. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

package cbl

import (
	"fmt"
	"math"
)

type Vec2f struct {
	X float32
	Y float32
}

func (v Vec2f) Add(rhs Vec2f) Vec2f {
	return Vec2f{v.X + rhs.X, v.Y + rhs.Y}
}

func (v Vec2f) Sub(rhs Vec2f) Vec2f {
	return Vec2f{v.X - rhs.X, v.Y - rhs.Y}
}

func (v Vec2f) Mul(rhs float32) Vec2f {
	return Vec2f{v.X * rhs, v.Y * rhs}
}

func (v Vec2f) Div(rhs float32) Vec2f {
	return Vec2f{v.X / rhs, v.Y / rhs}
}

func (v Vec2f) Dot(rhs Vec2f) float32 {
	return v.X*rhs.X + v.Y*rhs.Y
}

func (v Vec2f) Len() float32 {
	return float32(math.Sqrt(float64(v.Dot(v))))
}

func (v Vec2f) Dist(rhs Vec2f) float32 {
	x, y := v.X-rhs.X, v.Y-rhs.Y
	return float32(math.Sqrt(float64(x*x + y*y)))
}

func (v Vec2f) Norm() Vec2f {
	_len := v.Len()
	return Vec2f{v.X / _len, v.Y / _len}
}

func (v Vec2f) String() string {
	return fmt.Sprintf("{%v, %v}", v.X, v.Y)
}
