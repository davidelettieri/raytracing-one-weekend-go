package hittable

import (
	"math"

	"github.com/davidelettieri/raytracing-one-weekend-go/ray"
	"github.com/davidelettieri/raytracing-one-weekend-go/vec"
)

type Material interface {
	Scatter(rIn ray.Ray, rec HitRecord) (ray.Ray, vec.Color, bool)
}

type Lambertian struct {
	albedo vec.Color
}

func NewLambertian(color vec.Color) Material {
	return Lambertian{
		albedo: color,
	}
}

func (l Lambertian) Scatter(rIn ray.Ray, rec HitRecord) (ray.Ray, vec.Color, bool) {
	scatterDirection := rec.Normal().Add(vec.RandomUnitVector())

	if scatterDirection.NearZero() {
		scatterDirection = rec.Normal()
	}

	scattered := ray.NewRay(rec.Point(), scatterDirection)
	return scattered, l.albedo, true
}

type Metal struct {
	albedo vec.Color
	fuzz   float64
}

func NewMetal(color vec.Color, fuzz float64) Material {
	return Metal{
		albedo: color,
		fuzz:   math.Min(fuzz, 1),
	}
}

func (m Metal) Scatter(rIn ray.Ray, rec HitRecord) (ray.Ray, vec.Color, bool) {
	reflected := vec.Reflect(rIn.Direction(), rec.Normal())
	reflected = reflected.Unit().Add(vec.RandomUnitVector().Multiply(m.fuzz))
	scattered := ray.NewRay(rec.Point(), reflected)
	return scattered, m.albedo, vec.Dot(scattered.Direction(), rec.Normal()) > 0
}
