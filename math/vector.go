package math

import m "math"

const EPSILON = 0.0001

type Vector struct {
	x, y float64
}

func NewVector(x, y float64) Vector {
	return Vector{
		x: x,
		y: y,
	}
}

func NewVectorFromInt(x, y int) Vector {
	return Vector{
		x: float64(x),
		y: float64(y),
	}
}

func (v Vector) SquareLength() float64 {
	return v.x*v.x + v.y*v.y
}

func (v Vector) Length() float64 {
	return m.Sqrt(v.SquareLength())
}

func (v Vector) Normalized() Vector {
	l := v.Length()
	return Vector{
		x: v.x / l,
		y: v.y / l,
	}
}

func (v Vector) Times(coefficient float64) Vector {
	return Vector{
		x: v.x * coefficient,
		y: v.y * coefficient,
	}
}

func (v Vector) Add(o Vector) Vector {
	return Vector{
		x: v.x + o.x,
		y: v.y + o.y,
	}
}

func (v Vector) Sub(o Vector) Vector {
	return Vector{
		x: v.x - o.x,
		y: v.y - o.y,
	}
}

func (v Vector) EqRough(o Vector) bool {
	return v.Sub(o).SquareLength() < EPSILON
}

func (v Vector) X() float64 {
	return v.x
}

func (v Vector) Y() float64 {
	return v.y
}
