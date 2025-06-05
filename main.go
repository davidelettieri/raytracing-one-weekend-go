package main

import (
	"fmt"
	"os"

	"davidelettieri.it/raytracing/ray"
	"davidelettieri.it/raytracing/vec"
)

func rayColor(ray *ray.Ray) vec.Color {
	unit_direction := ray.GetDirection().Unit()
	a := 0.5 * (unit_direction.Y() + 1)
	return vec.NewColor(1.0, 1.0, 1.0).Multiply(1.0 - a).Add(vec.NewColor(0.5, 0.7, 1.0).Multiply(a))
}

func hitSphere(center *vec.Point3, radius float64, ray *ray.Ray) bool {
	origin := ray.GetOrigin()
	oc := center.Subtract(*origin)
	a := vec.Dot(*ray.GetDirection(), *ray.GetDirection())
	b := -2.0 * vec.Dot(*ray.GetDirection(), oc)
	c := vec.Dot(oc, oc) - radius*radius
	discriminant := b*b - 4*a*c
	return discriminant >= 0
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
			ray := ray.NewRay(&camera_center, &ray_direction)
			pixel_color := rayColor(&ray)
			vec.WriteColor(*os.Stdout, pixel_color)
		}
	}

	println("\rDone.			\n")
}
