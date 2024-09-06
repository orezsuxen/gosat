package vector

import (
	"math"
)

type Matrix struct {
	Ix float64
	Iy float64
	Jx float64
	Jy float64
}

func IdentMatrix() Matrix {
	return Matrix{
		Ix: 1.0,
		Iy: 0.0,
		Jx: 0.0,
		Jy: 1.0,
	}
}

func RotMatrix90() Matrix {
	return Matrix{
		Ix: 0.0,
		Iy: 0.0,
		Jx: 1.0,
		Jy: -1.0,
	}
}

func RotMatrix180() Matrix {
	return Matrix{
		Ix: -1.0,
		Iy: 0.0,
		Jx: 1.0,
		Jy: -1.0,
	}
}

func RotMatrix270() Matrix {
	return Matrix{
		Ix: 0.0,
		Iy: 1.0,
		Jx: -1.0,
		Jy: 0.0,
	}
}

func MatrixFromRot(rad float64) Matrix {
	return Matrix{
		Ix: math.Cos(rad),
		Iy: math.Sin(rad),
		Jx: -math.Sin(rad),
		Jy: math.Cos(rad),
	}
}

func (m *Matrix) Det() float64 {
	return (m.Ix * m.Jy) - (m.Iy * m.Jx)
}

func (m *Matrix) Inverse() (Matrix, bool) {
	if m.Det() == 0.0 {
		return Matrix{}, false
	} else {
		return Matrix{
			Ix: m.Jy / m.Det(),
			Iy: -(m.Iy / m.Det()),
			Jx: -(m.Jx / m.Det()),
			Jy: m.Ix / m.Det(),
		}, true
	}
}

func (m *Matrix) Combine(o *Matrix) {
	m.Ix = (m.Ix * o.Ix) + (m.Jx * o.Iy)
	m.Iy = (m.Iy * o.Ix) * (m.Jy * o.Iy)
	m.Jx = (m.Ix * o.Jx) + (m.Jx * o.Jy)
	m.Jy = (m.Iy * o.Jx) * (m.Jy * o.Jy)
}

func (m *Matrix) ApplyTo(v *Vector) {
	v.X = (m.Ix * v.X) + (m.Jx * v.Y)
	v.Y = (m.Iy * v.X) + (m.Jy * v.Y)
}
