package cbl

import (
	"fmt"
	"math"
)

type Vec2f struct {
	X float64
	Y float64
}

type Vector2i struct {
	X int
	Y int
}

func (v Vec2f) Add(rhs Vec2f) Vec2f {
	return Vec2f{v.X + rhs.X, v.Y + rhs.Y}
}

func (v Vec2f) Sub(rhs Vec2f) Vec2f {
	return Vec2f{v.X - rhs.X, v.Y - rhs.Y}
}

func (v Vec2f) Mul(rhs float64) Vec2f {
	return Vec2f{v.X * rhs, v.Y * rhs}
}

func (v Vec2f) Div(rhs float64) Vec2f {
	return Vec2f{v.X / rhs, v.Y / rhs}
}

func (v Vec2f) Dot(rhs Vec2f) float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vec2f) Len() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vec2f) Dist(rhs Vec2f) float64 {
	x, y := v.X-rhs.X, v.Y-rhs.Y
	return math.Sqrt(x*x + y*y)
}

func (v Vec2f) Norm() Vec2f {
	_len := v.Len()
	return Vec2f{v.X / _len, v.Y / _len}
}

func (v Vec2f) String() string {
	return fmt.Sprintf("{%v, %v}", v.X, v.Y)
}
