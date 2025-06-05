package hittable

import (
	"github.com/davidelettieri/raytracing-one-weekend-go/ray"
	"github.com/davidelettieri/raytracing-one-weekend-go/utils"
)

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

func (hl HittableList) Hit(ray ray.Ray, interval utils.Interval) (HitRecord, bool) {
	hit_anything := false
	closest_so_far := interval.GetMax()
	hit_record := HitRecord{}
	for _, obj := range hl.objects {
		obj_hit_record, hit := obj.Hit(ray, utils.NewInterval(interval.GetMin(), closest_so_far))

		if hit {
			hit_anything = true
			closest_so_far = obj_hit_record.t
			hit_record = obj_hit_record
		}
	}

	return hit_record, hit_anything
}
