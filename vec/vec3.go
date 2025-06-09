package vec

import (
	"math"

	"github.com/davidelettieri/raytracing-one-weekend-go/utils"
)

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

func (v Vec3) NearZero() bool {
	s := 1e-8
	return math.Abs(v.e[0]) < s && math.Abs(v.e[1]) < s && math.Abs(v.e[2]) < s
}

func ComponentsMultiply(u, v Vec3) Vec3 {
	return NewVec3(v.e[0]*u.e[0], v.e[1]*u.e[1], v.e[2]*u.e[2])
}

func Reflect(v Vec3, n Vec3) Vec3 {
	return v.Subtract(n.Multiply(Dot(v, n) * 2))
}

func Refract(uv, n Vec3, ETaiOverEtat float64) Vec3 {
	cosTheta := math.Min(Dot(uv.Negate(), n), 1.0)
	rOutPerp := uv.Add(n.Multiply(cosTheta)).Multiply(ETaiOverEtat)
	rOUtParallel := n.Multiply(-math.Sqrt(math.Abs(1.0 - rOutPerp.LengthSquared())))
	return rOUtParallel.Add(rOutPerp)
}

func Dot(u, v Vec3) float64 {
	return u.e[0]*v.e[0] + u.e[1]*v.e[1] + u.e[2]*v.e[2]
}

func Random() Vec3 {
	return NewVec3(utils.RandomFloat64(), utils.RandomFloat64(), utils.RandomFloat64())
}

func RandomInInterval(min, max float64) Vec3 {
	return NewVec3(utils.RandomFloat64InInterval(min, max), utils.RandomFloat64InInterval(min, max), utils.RandomFloat64InInterval(min, max))
}

func RandomUnitVector() Vec3 {
	for {
		p := RandomInInterval(-1, 1)
		lenqs := p.LengthSquared()
		if lenqs > 1e-160 && lenqs <= 1 {
			return p.Divide(math.Sqrt(lenqs))
		}
	}
}

func RandomOnHemisphere(normal Vec3) Vec3 {
	onUnitSphere := RandomUnitVector()
	if Dot(onUnitSphere, normal) > 0.0 {
		return onUnitSphere
	}
	return onUnitSphere.Negate()
}
