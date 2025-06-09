package hittable

import (
	"math"

	"github.com/davidelettieri/raytracing-one-weekend-go/ray"
	"github.com/davidelettieri/raytracing-one-weekend-go/utils"
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

type Dielectric struct {
	refractionIndex float64
}

func NewDielectric(refractionIndex float64) Material {
	return Dielectric{
		refractionIndex: refractionIndex,
	}
}

func (d Dielectric) Scatter(rIn ray.Ray, rec HitRecord) (ray.Ray, vec.Color, bool) {
	attenuation := vec.NewColor(1.0, 1.0, 1.0)
	var ri float64

	if rec.frontFace {
		ri = 1.0 / d.refractionIndex
	} else {
		ri = d.refractionIndex
	}

	unitDirection := rIn.Direction().Unit()
	cosTheta := math.Min(vec.Dot(unitDirection.Negate(), rec.Normal()), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	cannotRefract := ri*sinTheta > 1.0
	var direction vec.Vec3

	if cannotRefract || reflectance(cosTheta, ri) > utils.RandomFloat64() {
		direction = vec.Reflect(unitDirection, rec.Normal())
	} else {
		direction = vec.Refract(unitDirection, rec.Normal(), ri)
	}

	scattered := ray.NewRay(rec.Point(), direction)
	return scattered, attenuation, true
}

func reflectance(cosine, refractionIndex float64) float64 {
	r0 := (1 - refractionIndex) * (1 + refractionIndex)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}
