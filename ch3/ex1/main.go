package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xlmns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x1, y1 := corner(i+1, j)
			x2, y2 := corner(i, j)
			x3, y3 := corner(i, j+1)
			x4, y4 := corner(i+1, j+1)
			addPolygonIfValid(x1, y1, x2, y2, x3, y3, x4, y4)
		}
	}
	fmt.Println("</svg>")
}

func addPolygonIfValid(x1, y1, x2, y2, x3, y3, x4, y4 float64) {
	for _, val := range [...]float64{x1, y1, x2, y2, x3, y3, x4, y4} {
		if math.IsInf(val, 0) || math.IsNaN(val) {
			return
		}
	}
	fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
		x1, y1, x2, y2, x3, y3, x4, y4)
}

func corner(i, j int) (float64, float64) {
	// Find point (x, y) at corner of cell (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x, y, z) isometrically onto 2D SVG canvas (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
