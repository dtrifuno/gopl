package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
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
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:9999", nil))
}

func handler(rw http.ResponseWriter, r *http.Request) {
	fmt.Print("handling a connection\n")
	funcString := r.URL.Query().Get("function")

	f := waterdrop
	switch funcString {
	case "eggbox":
		f = eggbox
	case "saddle":
		f = saddle
	}

	writeSVG(rw, f)
}

func writeSVG(rw http.ResponseWriter, f func(float64, float64) float64) {
	rw.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(rw, "<?xml version=\"1.0\" encoding=\"utf-8\"?>\r\n")
	fmt.Fprintf(rw, "<svg version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" "+
		"x=\"0px\" y=\"0px\" viewBox=\"0 0 %d %d\" style=\"enable-background:new 0 0 %[1]d %[2]d;\" xml:space=\"preserve\">\r\n", width, height)

	fmt.Fprintf(rw, "<style type=\"text/css\">\r\n"+
		"\t.st0{fill:#FFFFFF;stroke:#FF0000;stroke-width:0.7;}\r\n"+
		"\t.st1{fill:#FFFFFF;stroke:#CC0033;stroke-width:0.7;}\r\n"+
		"\t.st2{fill:#FFFFFF;stroke:#990066;stroke-width:0.7;}\r\n"+
		"\t.st3{fill:#FFFFFF;stroke:#660099;stroke-width:0.7;}\r\n"+
		"\t.st4{fill:#FFFFFF;stroke:#3300CC;stroke-width:0.7;}\r\n"+
		"\t.st5{fill:#FFFFFF;stroke:#0000FF;stroke-width:0.7;}\r\n"+
		"</style>\r\n")

	// find max and min
	zMin, zMax := findMinAndMax(f)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			addPolygonIfValid(rw, float64(i), float64(j), zMin, zMax, f)
		}
	}
	fmt.Fprint(rw, "</svg>\n")
}

func findMinAndMax(f func(float64, float64) float64) (float64, float64) {
	zMax := math.Inf(-1)
	zMin := math.Inf(1)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange * ((float64(i)+0.5)/cells - 0.5)
			y := xyrange * ((float64(j)+0.5)/cells - 0.5)
			z := f(x, y)
			if math.IsInf(z, 0) || math.IsNaN(z) {
				continue
			}
			if z > zMax {
				zMax = z
			}
			if z < zMin {
				zMin = z
			}
		}
	}
	return zMin, zMax
}

func addPolygonIfValid(rw http.ResponseWriter, i, j, zMin, zMax float64, f func(float64, float64) float64) {
	x1, y1 := corner(i+1, j, f)
	x2, y2 := corner(i, j, f)
	x3, y3 := corner(i, j+1, f)
	x4, y4 := corner(i+1, j+1, f)

	// skip polygon if it contains an invalid value
	for _, val := range [...]float64{x1, y1, x2, y2, x3, y3, x4, y4} {
		if math.IsInf(val, 0) || math.IsNaN(val) {
			return
		}
	}

	// determine color of polygon
	xMid := xyrange * ((i+0.5)/cells - 0.5)
	yMid := xyrange * ((j+0.5)/cells - 0.5)
	zMid := f(xMid, yMid)
	zMidScaled := (zMid - zMin) / (zMax - zMin)
	stClass := int(math.Round(5 * zMidScaled))

	fmt.Fprintf(rw, "<polygon class=\"st%d\" points=\"%.1f,%.1f %.1f,%.1f %.1f,%.1f %.1f,%.1f\"/>\r\n",
		stClass, x1, y1, x2, y2, x3, y3, x4, y4)
}

func corner(i, j float64, f func(x, y float64) float64) (float64, float64) {
	// Find point (x, y) at corner of cell (i, j).
	x := xyrange * (i/cells - 0.5)
	y := xyrange * (j/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x, y, z) isometrically onto 2D SVG canvas (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func waterdrop(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func eggbox(x, y float64) float64 {
	r := math.Pow(math.Sin(x), 2) + math.Pow(math.Sin(y), 2)
	return r / 6
}

func saddle(x, y float64) float64 {
	r := x*x/20.0 - y*y/40.0
	return r / 20.0
}
