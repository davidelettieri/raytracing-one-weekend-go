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
	hit_record, hit := world.Hit(ray, utils.NewInterval(0, math.Inf(1)))
	if hit {
		return hit_record.GetNormal().Add(vec.NewColor(1, 1, 1)).Multiply(0.5)
	}

	unit_direction := ray.GetDirection().Unit()
	a := 0.5 * (unit_direction.Y() + 1)
	return vec.NewColor(1.0, 1.0, 1.0).Multiply(1.0 - a).Add(vec.NewColor(0.5, 0.7, 1.0).Multiply(a))
}

func main() {
	// Image
	aspect_ratio := 16.0 / 9.0
	image_width := 400

	// Calculate the image height, and ensure that it's at least 1.
	image_height := int(float64(image_width) / aspect_ratio)
	if image_height < 1 {
		image_height = 1
	}

	world := hittable.NewHittableList(hittable.NewSphere(vec.NewPoint3(0, 0, -1), 0.5))
	world.Add(hittable.NewSphere(vec.NewPoint3(0, -100.5, -1), 100))

	// Camera

	focal_length := 1.0
	viewport_height := 2.0
	viewport_width := viewport_height * (float64(image_width) / float64(image_height))
	camera_center := vec.NewPoint3(0, 0, 0)

	// Calculate the vectors across the horizontal and down the vertical viewport edges.
	viewport_u := vec.NewVec3(viewport_width, 0, 0)
	viewport_v := vec.NewVec3(0, -viewport_height, 0)

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.
	pixel_delta_u := viewport_u.Divide(float64(image_width))
	pixel_delta_v := viewport_v.Divide(float64(image_height))

	// Calculate the location of the upper left pixel.
	viewport_upper_left := camera_center.Subtract(vec.NewVec3(0, 0, focal_length)).Subtract(viewport_u.Divide(2)).Subtract(viewport_v.Divide(2))
	pixel00_loc := viewport_upper_left.Add(pixel_delta_u.Add(pixel_delta_v).Multiply(0.5))

	// Render

	fmt.Print("P3\n", image_width, " ", image_height, "\n255\n")

	for j := range image_height {
		println("\nScanlines remaining: ", image_height-j, " ")
		for i := range image_width {
			pixel_center := pixel00_loc.Add(pixel_delta_u.Multiply(float64(i))).Add(pixel_delta_v.Multiply(float64(j)))
			ray_direction := pixel_center.Subtract(camera_center)
			ray := ray.NewRay(camera_center, ray_direction)
			pixel_color := rayColor(ray, world)
			vec.WriteColor(*os.Stdout, pixel_color)
		}
	}

	println("\rDone.			\n")
}
