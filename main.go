package main

import (
	"github.com/davidelettieri/raytracing-one-weekend-go/camera"
	"github.com/davidelettieri/raytracing-one-weekend-go/hittable"
	"github.com/davidelettieri/raytracing-one-weekend-go/utils"
	"github.com/davidelettieri/raytracing-one-weekend-go/vec"
)

func main() {
	world := hittable.NewHittableList()

	refPoint := vec.NewPoint3(4, 0.2, 0)
	groundMaterial := hittable.NewLambertian(vec.NewColor(0.5, 0.5, 0.5))
	world.Add(hittable.NewSphere(vec.NewPoint3(0, -1000, 0), 1000, groundMaterial))

	for a := -11.0; a < 11; a++ {
		for b := -11.0; b < 11; b++ {
			chooseMat := utils.RandomFloat64()
			center := vec.NewPoint3(a+0.9*utils.RandomFloat64(), 0.2, b+0.9*utils.RandomFloat64())

			if center.Subtract(refPoint).Length() > 0.9 {
				if chooseMat < 0.8 {
					// diffuse
					albedo := vec.ComponentsMultiply(vec.Random(), vec.Random())
					sphereMaterial := hittable.NewLambertian(albedo)
					world.Add(hittable.NewSphere(center, 0.2, sphereMaterial))
				} else if chooseMat < 0.95 {
					albedo := vec.RandomInInterval(0.5, 1)
					fuzz := utils.RandomFloat64InInterval(0, 0.5)
					sphereMaterial := hittable.NewMetal(albedo, fuzz)
					world.Add(hittable.NewSphere(center, 0.2, sphereMaterial))
				} else {
					sphereMaterial := hittable.NewDielectric(1.5)
					world.Add(hittable.NewSphere(center, 0.2, sphereMaterial))
				}
			}
		}
	}

	material1 := hittable.NewDielectric(1.5)
	world.Add(hittable.NewSphere(vec.NewPoint3(0, 1, 0), 1.0, material1))

	material2 := hittable.NewLambertian(vec.NewColor(0.4, 0.2, 0.1))
	world.Add(hittable.NewSphere(vec.NewPoint3(-4, 1, 0), 1.0, material2))

	material3 := hittable.NewMetal(vec.NewColor(0.7, 0.6, 0.5), 0.0)
	world.Add(hittable.NewSphere(vec.NewPoint3(4, 1, 0), 1.0, material3))

	lookFrom := vec.NewPoint3(13, 2, 2)
	lookTo := vec.NewPoint3(0, 0, 0)
	upDirection := vec.NewVec3(0, 1, 0)
	cam := camera.NewCamera(16.0/9.0, 1200, 10, 50, 20, lookFrom, lookTo, upDirection, 0.6, 10.0)
	cam.Render(world)
}
