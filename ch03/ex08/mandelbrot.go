package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/big"
	"math/cmplx"
	"os"
)

type fltCmplx struct {
	r *big.Float
	i *big.Float
}

func (x *fltCmplx) add(y *fltCmplx) *fltCmplx {
	r := new(big.Float)
	r.Add(x.r, y.r)

	i := new(big.Float)
	i.Add(x.i, y.i)

	return &fltCmplx{r, i}
}

func (x *fltCmplx) mul(y *fltCmplx) *fltCmplx {
	ra := new(big.Float)
	ra.Mul(x.r, y.r)
	rb := new(big.Float)
	rb.Mul(x.i, y.i)
	r := new(big.Float)
	r.Sub(ra, rb)

	ia := new(big.Float)
	ia.Mul(x.i, y.r)
	ib := new(big.Float)
	ib.Mul(x.r, y.i)
	i := new(big.Float)
	i.Add(ia, ib)

	return &fltCmplx{r, i}
}

func (x *fltCmplx) abs() float64 {
	r, _ := x.r.Float64()
	i, _ := x.i.Float64()
	return math.Hypot(r, i)
}

type ratCmplx struct {
	r *big.Rat
	i *big.Rat
}

func (x *ratCmplx) add(y *ratCmplx) *ratCmplx {
	r := new(big.Rat)
	r.Add(x.r, y.r)

	i := new(big.Rat)
	i.Add(x.i, y.i)

	return &ratCmplx{r, i}
}

func (x *ratCmplx) mul(y *ratCmplx) *ratCmplx {
	ra := new(big.Rat)
	ra.Mul(x.r, y.r)
	rb := new(big.Rat)
	rb.Mul(x.i, y.i)
	r := new(big.Rat)
	r.Sub(ra, rb)

	ia := new(big.Rat)
	ia.Mul(x.i, y.r)
	ib := new(big.Rat)
	ib.Mul(x.r, y.i)
	i := new(big.Rat)
	i.Add(ia, ib)

	return &ratCmplx{r, i}
}

func (x *ratCmplx) abs() float64 {
	r, _ := x.r.Float64()
	i, _ := x.i.Float64()
	return math.Hypot(r, i)
}

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

			// var z complex128 = complex(x, y)
			// img.Set(px, py, mandelbrot128(z))

			// var z complex64 = complex(float32(x), float32(y))
			// img.Set(px, py, mandelbrot64(z))

			// z := fltCmplx{big.NewFloat(x), big.NewFloat(y)}
			// img.Set(px, py, mandelbrotF(&z))

			r := &big.Rat{}
			r.SetFloat64(x)
			i := &big.Rat{}
			i.SetFloat64(y)
			z := ratCmplx{r, i}
			img.Set(px, py, mandelbrotR(&z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot128(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.RGBA{0, 0, 0, 0xff}
}

func mandelbrot64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.RGBA{0, 0, 0, 0xff}
}

func mandelbrotF(z *fltCmplx) color.Color {
	const iterations = 200
	const contrast = 15

	var v *fltCmplx
	for n := uint8(0); n < iterations; n++ {
		v = (v.mul(v)).add(z)
		if v.abs() > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.RGBA{0, 0, 0, 0xff}
}

func mandelbrotR(z *ratCmplx) color.Color {
	const iterations = 200
	const contrast = 15

	var v *ratCmplx
	for n := uint8(0); n < iterations; n++ {
		v = (v.mul(v)).add(z)
		if v.abs() > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.RGBA{0, 0, 0, 0xff}
}
