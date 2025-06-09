package main

import (
	"github.com/davidelettieri/raytracing-one-weekend-go/camera"
	"github.com/davidelettieri/raytracing-one-weekend-go/hittable"
	"github.com/davidelettieri/raytracing-one-weekend-go/vec"
)

func main() {
	world := hittable.NewHittableList()

	materialGround := hittable.NewLambertian(vec.NewColor(0.8, 0.8, 0.0))
	materialCenter := hittable.NewLambertian(vec.NewColor(0.1, 0.2, 0.5))
	materialLeft := hittable.NewDielectric(1.5)
	materialBubble := hittable.NewDielectric(1.0 / 1.5)
	materialRight := hittable.NewMetal(vec.NewColor(0.8, 0.6, 0.2), 1)

	world.Add(hittable.NewSphere(vec.NewPoint3(0, -100.5, -1), 100, materialGround))
	world.Add(hittable.NewSphere(vec.NewPoint3(0, 0, -1.2), 0.5, materialCenter))
	world.Add(hittable.NewSphere(vec.NewPoint3(-1, 0, -1), 0.5, materialLeft))
	world.Add(hittable.NewSphere(vec.NewPoint3(-1, 0, -1), 0.4, materialBubble))
	world.Add(hittable.NewSphere(vec.NewPoint3(1, 0, -1), 0.5, materialRight))
	cam := camera.NewCamera(16.0/9.0, 400, 100, 50)
	cam.Render(world)
}
