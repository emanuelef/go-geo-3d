package position

import (
	"math"
)

type Point3D struct {
	X, Y, Z float64
}

func (a Point3D) Add(b Point3D) Point3D {
	return Point3D{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

func (a Point3D) Sub(b Point3D) Point3D {
	return Point3D{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

func (a Point3D) MultiplyByScalar(s float64) Point3D {
	return Point3D{
		X: a.X * s,
		Y: a.Y * s,
		Z: a.Z * s,
	}
}

func (a Point3D) Dot(b Point3D) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (p Point3D) Distance(a Point3D) float64 {
	return math.Sqrt(math.Pow((p.X-a.X), 2) + math.Pow((p.Y-a.Y), 2) + math.Pow((p.Z-a.Z), 2))
}
