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

type bigFloatComplex struct {
	re, im *big.Float
}

func newBigFloatComplex(re, im float64) *bigFloatComplex {
	return &bigFloatComplex{big.NewFloat(re), big.NewFloat(im)}
}

func (b *bigFloatComplex) add(x, y *bigFloatComplex) *bigFloatComplex {
	b.re.Add(x.re, y.re)
	b.im.Add(x.im, y.im)
	return b
}

func (b *bigFloatComplex) mul(x, y *bigFloatComplex) *bigFloatComplex {
	b.re.Sub(
		new(big.Float).Mul(x.re, y.re),
		new(big.Float).Mul(x.im, y.im),
	)
	b.im.Add(
		new(big.Float).Mul(x.re, y.im),
		new(big.Float).Mul(x.im, y.re),
	)
	return b
}

func (b *bigFloatComplex) abs() float64 {
	// 敗北感
	re, _ := b.re.Float64()
	im, _ := b.im.Float64()
	return math.Hypot(re, im)
}

type bigRatComplex struct {
	re, im *big.Rat
}

func newBigRatComplex(re, im float64) *bigRatComplex {
	return &bigRatComplex{new(big.Rat).SetFloat64(re), new(big.Rat).SetFloat64(im)}
}

func (b *bigRatComplex) add(x, y *bigRatComplex) *bigRatComplex {
	b.re.Add(x.re, y.re)
	b.im.Add(x.im, y.im)
	return b
}

func (b *bigRatComplex) mul(x, y *bigRatComplex) *bigRatComplex {
	b.re.Sub(
		new(big.Rat).Mul(x.re, y.re),
		new(big.Rat).Mul(x.im, y.im),
	)
	b.im.Add(
		new(big.Rat).Mul(x.re, y.im),
		new(big.Rat).Mul(x.im, y.re),
	)
	return b
}

func (x *bigRatComplex) abs() float64 {
	re, _ := x.re.Float64()
	im, _ := x.im.Float64()
	return math.Hypot(re, im)
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

			// var z *bigFloatComplex = newBigFloatComplex(x, y)
			// img.Set(px, py, mandelbrotF(z))

			var z *bigRatComplex = newBigRatComplex(x, y)
			img.Set(px, py, mandelbrotR(z))
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

func mandelbrotF(z *bigFloatComplex) color.Color {
	const iterations = 200
	const contrast = 15

	var v *bigFloatComplex = newBigFloatComplex(0.0, 0.0)
	for n := uint8(0); n < iterations; n++ {
		v.mul(v, v)
		v.add(v, z)
		if v.abs() > 2.0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.RGBA{0, 0, 0, 0xff}
}

func mandelbrotR(z *bigRatComplex) color.Color {
	const iterations = 200
	const contrast = 15

	var v *bigRatComplex = newBigRatComplex(0.0, 0.0)
	for n := uint8(0); n < iterations; n++ {
		v.mul(v, v)
		v.add(v, z)
		if v.abs() > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.RGBA{0, 0, 0, 0xff}
}
