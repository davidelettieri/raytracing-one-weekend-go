package hittable

import (
	"errors"
	"math"

	"github.com/davidelettieri/raytracing-one-weekend-go/ray"
	"github.com/davidelettieri/raytracing-one-weekend-go/vec"
)

type HitRecord struct {
	p          vec.Point3
	normal     vec.Vec3
	t          float64
	front_face bool
}

func (h *HitRecord) SetFaceNormal(ray ray.Ray, outward_normal vec.Vec3) error {
	if outward_normal.LengthSquared() != 1 {
		return errors.New("length of outward normal must be 1")
	}

	h.front_face = vec.Dot(ray.GetDirection(), outward_normal) < 0
	if h.front_face {
		h.normal = outward_normal
	} else {
		h.normal = outward_normal.Negate()
	}

	return nil
}

type Hittable interface {
	Hit(ray ray.Ray, ray_tmin, ray_tmax float64) bool
}

type Sphere struct {
	center vec.Point3
	radius float64
}

func (s Sphere) Hit(ray ray.Ray, ray_tmin, ray_tmax float64) (HitRecord, bool) {
	oc := s.center.Subtract(ray.GetOrigin())
	rayDirection := ray.GetDirection()
	a := rayDirection.LengthSquared()
	h := vec.Dot(rayDirection, oc)
	c := oc.LengthSquared() - s.radius*s.radius

	discriminant := h*h - a*c
	sqrtd := math.Sqrt(discriminant)

	root := (h - sqrtd) / a

	if root <= ray_tmin || root >= ray_tmax {
		root = (h + sqrtd) / a
		if root <= ray_tmin || root >= ray_tmax {
			return HitRecord{}, false
		}
	}

	p := ray.At(root)

	rec := HitRecord{
		t:      root,
		p:      p,
		normal: p.Subtract(s.center).Divide(s.radius),
	}

	outward_normal := rec.p.Subtract(s.center).Divide(s.radius)
	rec.SetFaceNormal(ray, outward_normal)

	return rec, true
}

func NewSphere(center vec.Point3, radius float64) Sphere {
	return Sphere{
		center: center,
		radius: math.Max(0, radius),
	}
}
