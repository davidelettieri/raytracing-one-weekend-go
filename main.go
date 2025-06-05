package main

import (
	"fmt"
	"math"
	"os"

	"github.com/davidelettieri/raytracing-one-weekend-go/hittable"
	"github.com/davidelettieri/raytracing-one-weekend-go/ray"
	"github.com/davidelettieri/raytracing-one-weekend-go/utils"
	"github.com/davidelettieri/raytracing-one-weekend-go/vec"
)

func rayColor(ray ray.Ray, world hittable.Hittable) vec.Color {
	hitRecord, hit := world.Hit(ray, utils.NewInterval(0, math.Inf(1)))
	if hit {
		return hitRecord.Normal().Add(vec.NewColor(1, 1, 1)).Multiply(0.5)
	}

	unitDirection := ray.Direction().Unit()
	a := 0.5 * (unitDirection.Y() + 1)
	return vec.NewColor(1.0, 1.0, 1.0).Multiply(1.0 - a).Add(vec.NewColor(0.5, 0.7, 1.0).Multiply(a))
}

func main() {
	// Image
	aspectRatio := 16.0 / 9.0
	imageWidth := 400

	// Calculate the image height, and ensure that it's at least 1.
	imageHeight := int(float64(imageWidth) / aspectRatio)
	if imageHeight < 1 {
		imageHeight = 1
	}

	world := hittable.NewHittableList(hittable.NewSphere(vec.NewPoint3(0, 0, -1), 0.5))
	world.Add(hittable.NewSphere(vec.NewPoint3(0, -100.5, -1), 100))

	// Camera

	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(imageWidth) / float64(imageHeight))
	cameraCenter := vec.NewPoint3(0, 0, 0)

	// Calculate the vectors across the horizontal and down the vertical viewport edges.
	viewportU := vec.NewVec3(viewportWidth, 0, 0)
	viewportV := vec.NewVec3(0, -viewportHeight, 0)

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.
	pixelDeltaU := viewportU.Divide(float64(imageWidth))
	pixelDeltaV := viewportV.Divide(float64(imageHeight))

	// Calculate the location of the upper left pixel.
	viewportUpperLeft := cameraCenter.Subtract(vec.NewVec3(0, 0, focalLength)).Subtract(viewportU.Divide(2)).Subtract(viewportV.Divide(2))
	pixel00Loc := viewportUpperLeft.Add(pixelDeltaU.Add(pixelDeltaV).Multiply(0.5))

	// Render

	fmt.Print("P3\n", imageWidth, " ", imageHeight, "\n255\n")

	for j := range imageHeight {
		println("\nScanlines remaining: ", imageHeight-j, " ")
		for i := range imageWidth {
			pixelCenter := pixel00Loc.Add(pixelDeltaU.Multiply(float64(i))).Add(pixelDeltaV.Multiply(float64(j)))
			rayDirection := pixelCenter.Subtract(cameraCenter)
			ray := ray.NewRay(cameraCenter, rayDirection)
			pixelColor := rayColor(ray, world)
			vec.WriteColor(*os.Stdout, pixelColor)
		}
	}

	println("\rDone.			\n")
}
