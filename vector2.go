package cbl

import (
	"math"
	"fmt"
)

type Vector2f struct {
	X float64
	Y float64
}

type Vector2i struct {
	X int
	Y int
}

func (v *Vector2f) Add(rhs *Vector2f) *Vector2f {
	return &Vector2f{v.X + rhs.X, v.Y + rhs.Y}
}

func (v *Vector2f) Sub(rhs *Vector2f) *Vector2f {
	return &Vector2f{v.X - rhs.X, v.Y - rhs.Y}
}

func (v *Vector2f) Mul(rhs *Vector2f) *Vector2f {
	return &Vector2f{v.X * rhs.X, v.Y * rhs.Y}
}

func (v *Vector2f) Div(rhs *Vector2f) *Vector2f {
	return &Vector2f{v.X / rhs.X, v.Y / rhs.Y}
}

func (v *Vector2f) Normalised() *Vector2f {
	_len := v.Length()
	return &Vector2f{v.X / _len, v.Y / _len}
}

func (v *Vector2f) Length() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v *Vector2f) String() string {
	return fmt.Sprintf("{%v, %v}", v.X, v.Y)
}