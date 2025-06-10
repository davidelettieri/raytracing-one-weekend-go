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

var hitInterval = utils.NewInterval(0.001, math.Inf(1))

type Camera struct {
	aspectRatio         float64
	imageWidth          int
	samplesPerPixel     int
	maxDepth            int
	verticalFieldOfView float64
	lookFrom            vec.Point3
	lookAt              vec.Point3
	upDirection         vec.Vec3
	defocusAngle        float64
	focusDistance       float64
	imageHeight         int
	pixelSamplesScale   float64
	center              vec.Point3
	pixel00Loc          vec.Point3
	pixelDeltaU         vec.Vec3
	pixelDeltaV         vec.Vec3
	u                   vec.Vec3
	v                   vec.Vec3
	w                   vec.Vec3
	defocusDiskU        vec.Vec3
	defocusDiskV        vec.Vec3
}

func NewCamera(
	aspectRatio float64,
	imageWidth, samplesPerPixel, maxDepth int,
	verticalFieldOfView float64,
	lookFrom, lookAt vec.Point3,
	upDirection vec.Vec3,
	defocusAngle, focusDistance float64) Camera {
	return Camera{
		aspectRatio:         aspectRatio,
		imageWidth:          imageWidth,
		samplesPerPixel:     samplesPerPixel,
		maxDepth:            maxDepth,
		verticalFieldOfView: verticalFieldOfView,
		lookFrom:            lookFrom,
		lookAt:              lookAt,
		upDirection:         upDirection,
		defocusAngle:        defocusAngle,
		focusDistance:       focusDistance,
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

	c.center = c.lookFrom

	theta := utils.DegreesToRadians(c.verticalFieldOfView)
	h := math.Tan(theta / 2)
	viewportHeight := 2.0 * h * c.focusDistance
	viewportWidth := viewportHeight * (float64(c.imageWidth) / float64(c.imageHeight))

	c.w = c.lookFrom.Subtract(c.lookAt).Unit()
	c.u = vec.Cross(c.upDirection, c.w).Unit()
	c.v = vec.Cross(c.w, c.u)

	// Calculate the vectors across the horizontal and down the vertical viewport edges.
	viewportU := c.u.Multiply(viewportWidth)
	viewportV := c.v.Multiply(-viewportHeight)

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.
	c.pixelDeltaU = viewportU.Divide(float64(c.imageWidth))
	c.pixelDeltaV = viewportV.Divide(float64(c.imageHeight))

	// Calculate the location of the upper left pixel.
	viewportUpperLeft := c.center.Subtract(c.w.Multiply(c.focusDistance)).Subtract(viewportU.Divide(2)).Subtract(viewportV.Divide(2))
	c.pixel00Loc = viewportUpperLeft.Add(c.pixelDeltaU.Add(c.pixelDeltaV).Multiply(0.5))

	// Calculate the camera defocus disk basis vectors.
	defocusRadius := c.focusDistance * math.Tan(utils.DegreesToRadians(c.defocusAngle/2))
	c.defocusDiskU = c.u.Multiply(defocusRadius)
	c.defocusDiskV = c.v.Multiply(defocusRadius)
}

func (c Camera) getRay(i, j int) ray.Ray {
	// Construct a camera ray originating from the defocus disk and directed at a randomly
	// sampled point around the pixel location i, j.

	offset := sampleSquare()
	pixelSample := c.pixel00Loc.Add(c.pixelDeltaU.Multiply(float64(i) + offset.X())).Add(c.pixelDeltaV.Multiply(float64(j) + offset.Y()))
	var rayOrigin vec.Point3

	if c.defocusAngle <= 0 {
		rayOrigin = c.center
	} else {
		rayOrigin = c.defocusDiskSample()
	}
	rayDirection := pixelSample.Subtract(rayOrigin)

	return ray.NewRay(rayOrigin, rayDirection)
}

func (c Camera) defocusDiskSample() vec.Point3 {
	p := vec.RandomInUnitDisk()
	return c.center.Add(c.defocusDiskU.Multiply(p.X())).Add(c.defocusDiskV.Multiply(p.Y()))
}

func sampleSquare() vec.Vec3 {
	return vec.NewVec3(utils.RandomFloat64()-0.5, utils.RandomFloat64()-0.5, 0)
}

func rayColor(r ray.Ray, depth int, world hittable.Hittable) vec.Color {
	if depth <= 0 {
		return vec.NewColor(0, 0, 0)
	}

	rec, hit := world.Hit(r, hitInterval)
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
