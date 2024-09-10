package vector

import (
	"math"
)

type Vector struct {
	X float64
	Y float64
}

func (v *Vector) Dist(o *Vector) float64 {
	return math.Sqrt(math.Pow((v.X-o.X), 2) + math.Pow((v.Y-o.Y), 2))
}

func (v *Vector) Magn() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func (v *Vector) Dot(o *Vector) float64 {
	return (v.X * o.X) + (v.Y * o.Y)
}

func (v *Vector) Cross(o *Vector) float64 {
	return (v.X * o.Y) - (v.Y * o.X)
}

func (v *Vector) Unit() Vector {
	return Vector{
		X: v.X / v.Magn(),
		Y: v.Y / v.Magn(),
	}
}

func (v *Vector) AsRot() float64 {
	return math.Atan2(v.Unit().Y, v.Unit().X)
}

func VectorFromRot(rad float64) Vector {
	return Vector{
		X: math.Cos(rad),
		Y: math.Sin(rad),
	}
}

func (v *Vector) TransformBy(m *Matrix) {
	m.ApplyTo(v) //reusing Matrix apply
}

// =================================================
// Add

func Add(v *Vector, o *Vector) Vector {
	return Vector{
		X: v.X + o.X,
		Y: v.Y + o.Y,
	}
}
func AddScalar(v *Vector, s float64) Vector {
	return Vector{
		X: v.X + s,
		Y: v.Y + s,
	}
}

func (v *Vector) Add(o *Vector) {
	v.X += o.X
	v.Y += o.Y
}

func (v *Vector) AddScalar(s float64) {
	v.X += s
	v.Y += s
}

//Sub

func Sub(v *Vector, o *Vector) Vector {
	return Vector{
		X: v.X - o.X,
		Y: v.Y - o.Y,
	}
}
func SubScalar(v *Vector, s float64) Vector {
	return Vector{
		X: v.X - s,
		Y: v.Y - s,
	}
}

func (v *Vector) Sub(o *Vector) {
	v.X -= o.X
	v.Y -= o.Y
}

func (v *Vector) SubScalar(s float64) {
	v.X -= s
	v.Y -= s
}

// Mul

func Mul(v *Vector, o *Vector) Vector {
	return Vector{
		X: v.X * o.X,
		Y: v.Y * o.Y,
	}
}

func MulScalar(v *Vector, s float64) Vector {
	return Vector{
		X: v.X * s,
		Y: v.Y * s,
	}
}

func (v *Vector) Mul(o *Vector) {
	v.X *= o.X
	v.Y *= o.Y
}

func (v *Vector) MulScalar(s float64) {
	v.X *= s
	v.Y *= s
}

//Div

func Div(v *Vector, o *Vector) Vector {
	return Vector{
		X: v.X / o.X,
		Y: v.Y / o.Y,
	}
}

func DivScalar(v *Vector, s float64) Vector {
	return Vector{
		X: v.X / s,
		Y: v.Y / s,
	}
}

func (v *Vector) Div(o *Vector) {
	v.X /= o.X
	v.Y /= o.Y
}

func (v *Vector) DivScalar(s float64) {
	v.X /= s
	v.Y /= s
}
