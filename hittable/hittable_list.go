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
	closestSoFar := interval.Max()
	hitRecord := HitRecord{}
	for _, obj := range hl.objects {
		objHitRecord, hit := obj.Hit(ray, utils.NewInterval(interval.Min(), closestSoFar))

		if hit {
			hitAnything = true
			closestSoFar = objHitRecord.t
			hitRecord = objHitRecord
		}
	}

	return hitRecord, hitAnything
}
