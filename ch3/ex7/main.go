package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 20

	var roots = [4]complex128{complex(1, 0), complex(-1, 0), complex(0, 1), complex(0, -1)}
	var rootColor = [4](func(uint8) color.Color){
		func(n uint8) color.Color {
			return color.RGBA{R: 255, A: 255 - contrast*n}
		},
		func(n uint8) color.Color {
			return color.RGBA{G: 255, A: 255 - contrast*n}
		},
		func(n uint8) color.Color {
			return color.RGBA{B: 255, A: 255 - contrast*n}
		},
		func(n uint8) color.Color {
			return color.RGBA{R: 255, G: 255, A: 255 - contrast*n}
		},
	}

	for n := uint8(0); n < iterations; n++ {
		for i, root := range roots {
			if cmplx.Abs(z-root) < 1e-8 {
				return rootColor[i](n)
			}
		}
		z = z - f(z)/fPrime(z)
	}
	return color.Black
}

func f(z complex128) complex128 {
	return z*z*z*z - 1
}

func fPrime(z complex128) complex128 {
	return 4 * z * z * z
}
