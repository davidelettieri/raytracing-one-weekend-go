package hittable

import (
	"errors"

	"github.com/davidelettieri/raytracing-one-weekend-go/ray"
	"github.com/davidelettieri/raytracing-one-weekend-go/utils"
	"github.com/davidelettieri/raytracing-one-weekend-go/vec"
)

type HitRecord struct {
	p         vec.Point3
	normal    vec.Vec3
	t         float64
	frontFace bool
}

func (h HitRecord) Normal() vec.Vec3 {
	return h.normal
}

func (h HitRecord) Point() vec.Point3 {
	return h.p
}

func (h *HitRecord) SetFaceNormal(ray ray.Ray, outwardNormal vec.Vec3) error {
	if outwardNormal.LengthSquared() != 1 {
		return errors.New("length of outward normal must be 1")
	}

	h.frontFace = vec.Dot(ray.Direction(), outwardNormal) < 0
	if h.frontFace {
		h.normal = outwardNormal
	} else {
		h.normal = outwardNormal.Negate()
	}

	return nil
}

type Hittable interface {
	Hit(ray ray.Ray, interval utils.Interval) (HitRecord, bool)
}
