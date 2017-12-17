package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6
)

type params struct {
	height int
	width  int
	color  string
}

func (p *params) set(values url.Values) {
	w, _ := strconv.Atoi(values.Get("width"))
	p.width = w
	h, _ := strconv.Atoi(values.Get("height"))
	p.height = h
	c := values.Get("color")
	p.color = c
}

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		p := new(params)
		p.set(r.URL.Query())
		w.Header().Set("Content-Type", "image/svg+xml")
		plot(w, p)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

func plot(out io.Writer, p *params) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", p.width, p.height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, p)
			bx, by := corner(i, j, p)
			cx, cy := corner(i, j+1, p)
			dx, dy := corner(i+1, j+1, p)
			if isValid([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, p.color)
			}
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func isValid(vals []float64) bool {
	for _, v := range vals {
		if math.IsInf(v, 0) {
			return false
		}
	}
	return true
}

func corner(i, j int, p *params) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	xyscale := float64(p.width) / 2 / xyrange
	zscale := float64(p.height) * 0.4

	sx := float64(p.width)/2 + (x-y)*cos30*xyscale
	sy := float64(p.height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
