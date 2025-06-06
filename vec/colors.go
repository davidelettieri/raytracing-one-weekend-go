package vec

import (
	"fmt"
	"os"

	"github.com/davidelettieri/raytracing-one-weekend-go/utils"
)

type Color = Vec3

func NewColor(x, y, z float64) Color {
	return Color{
		e: [3]float64{x, y, z},
	}
}

func WriteColor(out os.File, c Color) {
	r := c.X()
	g := c.Y()
	b := c.Z()

	intensity := utils.NewInterval(0.0, 0.999)

	rbyte := int(256 * intensity.Clamp(r))
	gbyte := int(256 * intensity.Clamp(g))
	bbyte := int(256 * intensity.Clamp(b))

	out.WriteString(fmt.Sprint(rbyte, gbyte, bbyte, "\n"))
}
