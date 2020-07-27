package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
	dx                     = float64(xmax-xmin) / width
	dy                     = float64(ymax-ymin) / height
	iterations             = 200
)

func main() {

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)*dy + ymin
		for px := 0; px < width; px++ {
			x := float64(px)*dx + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func getIterations(z complex128) uint8 {
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return n
		}
	}
	return iterations
}

func average(a, b, c, d, e uint8) uint8 {
	var sum float64 = 0
	sum = float64(a) + float64(b) + float64(c) + float64(d) + float64(e)
	return uint8(math.RoundToEven(sum / 5.0))
}

func mandelbrot(z complex128) color.Color {
	const contrast = 15

	n0 := getIterations(z)
	n1 := getIterations(z + complex(0.5*dx, 0.5*dy))
	n2 := getIterations(z + complex(0.5*dx, -0.5*dy))
	n3 := getIterations(z + complex(-0.5*dx, 0.5*dy))
	n4 := getIterations(z + complex(-0.5*dx, -0.5*dy))

	n := average(n0, n1, n2, n3, n4)
	if n < iterations {
		return color.RGBA{
			B: 255 - contrast*n,
			G: n*contrast - 255,
			R: 0,
			A: 255,
		}
	}
	return color.Black
}
