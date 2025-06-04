package vec

import (
	"fmt"
	"os"
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

	rbyte := int(255.999 * r)
	gbyte := int(255.999 * g)
	bbyte := int(255.999 * b)

	out.WriteString(fmt.Sprint(rbyte, gbyte, bbyte, "\n"))
}
