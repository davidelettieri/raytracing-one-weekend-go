package hittable

import (
	"errors"

	"github.com/davidelettieri/raytracing-one-weekend-go/ray"
	"github.com/davidelettieri/raytracing-one-weekend-go/utils"
	"github.com/davidelettieri/raytracing-one-weekend-go/vec"
)

type HitRecord struct {
	p          vec.Point3
	normal     vec.Vec3
	t          float64
	front_face bool
}

func (h HitRecord) GetNormal() vec.Vec3 {
	return h.normal
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
	Hit(ray ray.Ray, interval utils.Interval) (HitRecord, bool)
}
