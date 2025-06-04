package main

import (
	"fmt"
	"os"

	"davidelettieri.it/raytracing/vec"
)

func main() {
	image_width := 256
	image_height := 256

	fmt.Print("P3\n", image_width, " ", image_height, "\n255\n")

	for j := range image_height {
		println("\nScanlines remaining: ", image_height-j, " ")
		for i := range image_width {
			r := float64(i) / (float64(image_width) - 1)
			g := float64(j) / (float64(image_height) - 1)
			b := 0.0

			pixel_color := vec.NewColor(r, g, b)

			vec.WriteColor(*os.Stdout, pixel_color)
		}
	}

	println("\rDone.			\n")
}
