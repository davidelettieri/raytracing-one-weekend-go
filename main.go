package main

import (
	"github.com/davidelettieri/raytracing-one-weekend-go/camera"
	"github.com/davidelettieri/raytracing-one-weekend-go/hittable"
	"github.com/davidelettieri/raytracing-one-weekend-go/vec"
)

func main() {
	world := hittable.NewHittableList(hittable.NewSphere(vec.NewPoint3(0, 0, -1), 0.5))
	world.Add(hittable.NewSphere(vec.NewPoint3(0, -100.5, -1), 100))
	cam := camera.NewCamera(16.0/9.0, 400, 100)
	cam.Render(world)
}
