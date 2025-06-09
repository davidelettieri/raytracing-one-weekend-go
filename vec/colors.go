package vec

import (
	"fmt"
	"math"
	"os"

	"github.com/davidelettieri/raytracing-one-weekend-go/utils"
)

type Color = Vec3

func NewColor(x, y, z float64) Color {
	return Color{
		e: [3]float64{x, y, z},
	}
}

func linearToGamma(linearComponent float64) float64 {
	if linearComponent > 0 {
		return math.Sqrt(linearComponent)
	}
	return 0
}

func WriteColor(out os.File, c Color) {
	r := c.X()
	g := c.Y()
	b := c.Z()

	// Apply a linear to gamma transform for gamma 2
	r = linearToGamma(r)
	g = linearToGamma(g)
	b = linearToGamma(b)

	intensity := utils.NewInterval(0.000, 0.999)

	rbyte := int(256 * intensity.Clamp(r))
	gbyte := int(256 * intensity.Clamp(g))
	bbyte := int(256 * intensity.Clamp(b))

	out.WriteString(fmt.Sprint(rbyte, gbyte, bbyte, "\n"))
}
