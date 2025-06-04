package vec

import "math"

type Vec3 struct {
	e [3]float64
}

type Point3 = Vec3

func NewPoint3(x, y, z float64) Point3 {
	return Vec3{
		e: [3]float64{x, y, z},
	}
}

func NewVec3(x, y, z float64) Vec3 {
	return Vec3{
		e: [3]float64{x, y, z},
	}
}

func (v Vec3) X() float64 { return v.e[0] }
func (v Vec3) Y() float64 { return v.e[1] }
func (v Vec3) Z() float64 { return v.e[2] }

func (v Vec3) Negate() Vec3 {
	return Vec3{
		e: [3]float64{-v.e[0], -v.e[1], -v.e[2]},
	}
}

func (v Vec3) Add(w Vec3) Vec3 {
	u := Vec3{}
	u.e[0] = v.e[0] + w.e[0]
	u.e[1] = v.e[1] + w.e[1]
	u.e[2] = v.e[2] + w.e[2]
	return u
}

func (v Vec3) Subtract(w Vec3) Vec3 {
	u := Vec3{}
	u.e[0] = v.e[0] - w.e[0]
	u.e[1] = v.e[1] - w.e[1]
	u.e[2] = v.e[2] - w.e[2]
	return u
}

func (v Vec3) Multiply(t float64) Vec3 {
	u := Vec3{}
	u.e[0] = v.e[0] * t
	u.e[1] = v.e[1] * t
	u.e[2] = v.e[2] * t
	return u
}

func (v Vec3) Divide(t float64) Vec3 {
	return v.Multiply(1 / t)
}

func (v Vec3) LengthSquared() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) Unit() Vec3 {
	return v.Divide(v.Length())
}
