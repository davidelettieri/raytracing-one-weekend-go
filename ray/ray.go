package ray

import "github.com/davidelettieri/raytracing-one-weekend-go/vec"

type Ray struct {
	origin    *vec.Point3
	direction *vec.Vec3
}

func NewRay(origin *vec.Point3, direction *vec.Vec3) Ray {
	return Ray{
		origin:    origin,
		direction: direction,
	}
}

func (r Ray) GetOrigin() *vec.Point3 {
	return r.origin
}

func (r Ray) GetDirection() *vec.Vec3 {
	return r.direction
}

func (r Ray) At(t float64) vec.Point3 {
	return r.origin.Add(r.direction.Multiply(t))
}
