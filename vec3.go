// Copyright (C) 2013 Ryan Chew. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

package cbl

import (
	"fmt"
	"math"
)

type Vec3f struct {
	X float32
	Y float32
	Z float32
}

func (v Vec3f) Add(rhs Vec3f) Vec3f {
	return Vec3f{v.X + rhs.X, v.Y + rhs.Y, v.Z + rhs.Z}
}

func (v Vec3f) Sub(rhs Vec3f) Vec3f {
	return Vec3f{v.X - rhs.X, v.Y - rhs.Y, v.Z - rhs.Z}
}

func (v Vec3f) Mul(rhs float32) Vec3f {
	return Vec3f{v.X * rhs, v.Y * rhs, v.Z * rhs}
}

func (v Vec3f) Div(rhs float32) Vec3f {
	return Vec3f{v.X / rhs, v.Y / rhs, v.Z / rhs}
}

func (v Vec3f) Dot(rhs Vec3f) float32 {
	return v.X*rhs.X + v.Y*rhs.Y + v.Z*rhs.Z
}

func (v Vec3f) Len() float32 {
	return float32(math.Sqrt(float64(v.Dot(v))))
}

func (v Vec3f) Dist(rhs Vec3f) float32 {
	x, y, z := v.X-rhs.X, v.Y-rhs.Y, v.Z-rhs.Z
	return float32(math.Sqrt(float64(x*x + y*y + z*z)))
}

func (v Vec3f) Norm() Vec3f {
	_len := v.Len()
	return Vec3f{v.X / _len, v.Y / _len, v.Z / _len}
}

func (v Vec3f) String() string {
	return fmt.Sprintf("{%v, %v, %v}", v.X, v.Y, v.Z)
}
