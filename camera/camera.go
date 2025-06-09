package camera

import (
	"fmt"
	"math"
	"os"

	"github.com/davidelettieri/raytracing-one-weekend-go/hittable"
	"github.com/davidelettieri/raytracing-one-weekend-go/ray"
	"github.com/davidelettieri/raytracing-one-weekend-go/utils"
	"github.com/davidelettieri/raytracing-one-weekend-go/vec"
)

type Camera struct {
	aspectRatio       float64
	imageWidth        int
	samplesPerPixel   int
	maxDepth          int
	imageHeight       int
	pixelSamplesScale float64
	center            vec.Point3
	pixel00Loc        vec.Point3
	pixelDeltaU       vec.Vec3
	pixelDeltaV       vec.Vec3
}

func NewCamera(aspectRatio float64, imageWidth, samplesPerPixel, maxDepth int) Camera {
	return Camera{
		aspectRatio:     aspectRatio,
		imageWidth:      imageWidth,
		samplesPerPixel: samplesPerPixel,
		maxDepth:        maxDepth,
	}
}

func (c Camera) Render(world hittable.Hittable) {
	c.initialize()

	fmt.Print("P3\n", c.imageWidth, " ", c.imageHeight, "\n255\n")

	for j := range c.imageHeight {
		println("\nScanlines remaining: ", c.imageHeight-j, " ")
		for i := range c.imageWidth {
			pixelColor := vec.NewColor(0, 0, 0)
			for range c.samplesPerPixel {
				ray := c.getRay(i, j)
				pixelColor = pixelColor.Add(rayColor(ray, c.maxDepth, world))
			}

			vec.WriteColor(*os.Stdout, pixelColor.Multiply(c.pixelSamplesScale))
		}
	}

	println("\rDone.			\n")
}

func (c *Camera) initialize() {
	// Calculate the image height, and ensure that it's at least 1.
	c.imageHeight = int(float64(c.imageWidth) / c.aspectRatio)
	if c.imageHeight < 1 {
		c.imageHeight = 1
	}

	c.pixelSamplesScale = 1.0 / float64(c.samplesPerPixel)

	// Camera

	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(c.imageWidth) / float64(c.imageHeight))
	c.center = vec.NewPoint3(0, 0, 0)

	// Calculate the vectors across the horizontal and down the vertical viewport edges.
	viewportU := vec.NewVec3(viewportWidth, 0, 0)
	viewportV := vec.NewVec3(0, -viewportHeight, 0)

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.
	c.pixelDeltaU = viewportU.Divide(float64(c.imageWidth))
	c.pixelDeltaV = viewportV.Divide(float64(c.imageHeight))

	// Calculate the location of the upper left pixel.
	viewportUpperLeft := c.center.Subtract(vec.NewVec3(0, 0, focalLength)).Subtract(viewportU.Divide(2)).Subtract(viewportV.Divide(2))
	c.pixel00Loc = viewportUpperLeft.Add(c.pixelDeltaU.Add(c.pixelDeltaV).Multiply(0.5))
}

func (c Camera) getRay(i, j int) ray.Ray {
	offset := sampleSquare()
	pixelSample := c.pixel00Loc.Add(c.pixelDeltaU.Multiply(float64(i) + offset.X())).Add(c.pixelDeltaV.Multiply(float64(j) + offset.Y()))
	rayOrigin := c.center
	rayDirection := pixelSample.Subtract(rayOrigin)

	return ray.NewRay(rayOrigin, rayDirection)
}

func sampleSquare() vec.Vec3 {
	return vec.NewVec3(utils.RandomFloat64()-0.5, utils.RandomFloat64()-0.5, 0)
}

func rayColor(r ray.Ray, depth int, world hittable.Hittable) vec.Color {
	if depth < 0 {
		return vec.NewColor(0, 0, 0)
	}

	rec, hit := world.Hit(r, utils.NewInterval(0.001, math.Inf(1)))
	if hit {
		scattered, attenuation, ok := rec.Material().Scatter(r, rec)
		if ok {
			return vec.ComponentsMultiply(attenuation, rayColor(scattered, depth-1, world))
		}

		return vec.NewColor(0, 0, 0)
	}

	unitDirection := r.Direction().Unit()
	a := 0.5 * (unitDirection.Y() + 1)
	return vec.NewColor(1.0, 1.0, 1.0).Multiply(1.0 - a).Add(vec.NewColor(0.5, 0.7, 1.0).Multiply(a))
}
