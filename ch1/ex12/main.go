package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func handler(rw http.ResponseWriter, r *http.Request) {
	fmt.Print("handling a connection\n")
	params := map[string]float64{
		"cycles":  5,
		"res":     0.001,
		"size":    100,
		"nframes": 64,
		"delay":   8,
	}
	for k := range params {
		valString := r.URL.Query().Get(k)
		value, err := strconv.ParseFloat(valString, 64)
		if valString != "" && err == nil {
			params[k] = value
		}
	}
	rw.Header().Set("Content-Type", "image/gif")
	lissajous(rw, params)
}

var palette = []color.Color{
	color.Black,
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0xff, 0xa5, 0x00, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0x80, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0x4b, 0x00, 0x82, 0xff},
	color.RGBA{0xee, 0x82, 0xee, 0xff},
}

func lissajous(out http.ResponseWriter, params map[string]float64) {
	var (
		cycles  = params["cycles"]
		res     = params["res"]
		size    = int(params["size"])
		nframes = int(params["nframes"])
		delay   = int(params["delay"])
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	var colorIndex uint8
	for i := 0; i < nframes; i++ {
		colorIndex = uint8((i % (len(palette) - 1)) + 1)
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(float64(size)*y+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
