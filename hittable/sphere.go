package hittable

import (
	"math"

	"github.com/davidelettieri/raytracing-one-weekend-go/ray"
	"github.com/davidelettieri/raytracing-one-weekend-go/utils"
	"github.com/davidelettieri/raytracing-one-weekend-go/vec"
)

type Sphere struct {
	center   vec.Point3
	radius   float64
	material Material
}

func (s Sphere) Hit(ray ray.Ray, interval utils.Interval) (HitRecord, bool) {
	oc := s.center.Subtract(ray.Origin())
	rayDirection := ray.Direction()
	a := rayDirection.LengthSquared()
	h := vec.Dot(rayDirection, oc)
	c := oc.LengthSquared() - s.radius*s.radius

	discriminant := h*h - a*c
	if discriminant < 0 {
		return HitRecord{}, false
	}

	sqrtd := math.Sqrt(discriminant)

	root := (h - sqrtd) / a

	if !interval.Surrounds(root) {
		root = (h + sqrtd) / a
		if !interval.Surrounds(root) {
			return HitRecord{}, false
		}
	}

	p := ray.At(root)

	rec := HitRecord{
		t:        root,
		p:        p,
		normal:   p.Subtract(s.center).Divide(s.radius),
		material: s.material,
	}

	outwardNormal := rec.p.Subtract(s.center).Divide(s.radius)
	rec.SetFaceNormal(ray, outwardNormal)

	return rec, true
}

func NewSphere(center vec.Point3, radius float64, material Material) Sphere {
	return Sphere{
		center:   center,
		radius:   math.Max(0, radius),
		material: material,
	}
}
