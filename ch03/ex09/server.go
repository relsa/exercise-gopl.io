package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"net/url"
	"strconv"
)

type params struct {
	centerX int
	centerY int
	scale   float64
}

func (p *params) set(values url.Values) {
	x, _ := strconv.Atoi(values.Get("x"))
	p.centerX = x
	y, _ := strconv.Atoi(values.Get("y"))
	p.centerY = y
	s, _ := strconv.ParseFloat(values.Get("scale"), 64)
	p.scale = s
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		p := new(params)
		p.set(r.URL.Query())
		plot(w, p)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

func plot(out io.Writer, p *params) {
	const (
		xmin_, ymin_, xmax_, ymax_ = -2, -2, +2, +2
		width, height              = 1024, 1024
	)

	xmin := xmin_ / p.scale
	ymin := ymin_ / p.scale
	xmax := xmax_ / p.scale
	ymax := ymax_ / p.scale

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py-p.centerY)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px-p.centerX)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
