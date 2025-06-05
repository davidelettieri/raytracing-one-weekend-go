package hittable

import "github.com/davidelettieri/raytracing-one-weekend-go/ray"

type HittableList struct {
	objects []Hittable
}

func NewHittableList(object Hittable) HittableList {
	hl := HittableList{
		objects: []Hittable{},
	}

	hl.Add(object)
	return hl
}

func (h *HittableList) Add(object Hittable) {
	h.objects = append(h.objects, object)
}

func (h *HittableList) Clear() {
	h.objects = []Hittable{}
}

func (hl HittableList) Hit(ray ray.Ray, ray_tmin, ray_tmax float64) (HitRecord, bool) {
	hit_anything := false
	closest_so_far := ray_tmax
	hit_record := HitRecord{}
	for _, obj := range hl.objects {
		obj_hit_record, hit := obj.Hit(ray, ray_tmin, closest_so_far)

		if hit {
			hit_anything = true
			closest_so_far = obj_hit_record.t
			hit_record = obj_hit_record
		}
	}

	return hit_record, hit_anything
}
