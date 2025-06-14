package hittable

import (
	"github.com/davidelettieri/raytracing-one-weekend-go/ray"
	"github.com/davidelettieri/raytracing-one-weekend-go/utils"
)

type HittableList struct {
	objects []Hittable
}

func NewHittableList() HittableList {
	hl := HittableList{
		objects: []Hittable{},
	}
	return hl
}

func (h *HittableList) Add(object Hittable) {
	h.objects = append(h.objects, object)
}

func (h *HittableList) Clear() {
	h.objects = []Hittable{}
}

func (hl HittableList) Hit(ray ray.Ray, interval utils.Interval) (HitRecord, bool) {
	hitAnything := false
	localInterval := utils.NewInterval(interval.Min(), interval.Max())
	hitRecord := HitRecord{}
	for _, obj := range hl.objects {
		objHitRecord, hit := obj.Hit(ray, localInterval)

		if hit {
			hitAnything = true
			localInterval.SetMax(objHitRecord.t)
			hitRecord = objHitRecord
		}
	}

	return hitRecord, hitAnything
}
