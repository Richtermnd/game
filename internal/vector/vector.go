package vector

import "math"

// Vector
//
// All methods use value (not pointer(!)) and returns a new Vector
// I think it's more safety and easy to use.
// You can change vector with assigment to X and Y directly
type Vector struct {
	X, Y float64
}

func New[T int | float64](x, y T) Vector {
	return Vector{
		X: float64(x),
		Y: float64(y),
	}
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v Vector) Sub(v2 Vector) Vector {
	return Vector{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

func (v Vector) Scale(scale float64) Vector {
	return Vector{
		X: v.X * scale,
		Y: v.Y * scale,
	}
}

func (v *Vector) Len() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector) Normalize() Vector {
	l := v.Len()
	if l == 0 {
		return Vector{}
	}
	return Vector{
		X: v.X / l,
		Y: v.Y / l,
	}
}

func (v *Vector) Dot(v2 Vector) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func Add(v1, v2 Vector) Vector {
	return v1.Add(v2)
}

func Sub(v1, v2 Vector) Vector {
	return v1.Sub(v2)
}

func Scale(v Vector, scale float64) Vector {
	return v.Scale(scale)
}
