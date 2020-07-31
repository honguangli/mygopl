package coloredpoint

import (
	"image/color"
	"mygopl/ch06/geometry"
)

type ColoredPoint struct {
	geometry.Point
	Color color.RGBA
}
