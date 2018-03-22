package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

const (
	gifFormat  = "gif"
	jpegFormat = "jpeg"
	pngFormat  = "png"
)

var (
	imageFormat = flag.String("format", pngFormat, "format")
)

func main() {
	flag.Parse()

	img, err := readImage(os.Stdin)
	if err != nil {
		fmt.Errorf("fail to read: %v", err)
		os.Exit(1)
	}

	switch *imageFormat {
	case gifFormat:
		err = toGIF(img, os.Stdout)
	case jpegFormat:
		err = toJPEG(img, os.Stdout)
	case pngFormat:
		err = toPNG(img, os.Stdout)
	default:
		fmt.Fprintf(os.Stderr, "format: [gif|jpeg|png]")
	}

	if err != nil {
		fmt.Errorf("fail to write: %v", err)
		os.Exit(1)
	}
}

func readImage(in io.Reader) (image.Image, error) {
	img, kind, err := image.Decode(in)
	if err != nil {
		return nil, err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return img, nil
}

func toGIF(img image.Image, out io.Writer) error {
	return gif.Encode(out, img, &gif.Options{NumColors: 256, Quantizer: nil, Drawer: nil})
}

func toJPEG(img image.Image, out io.Writer) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toPNG(img image.Image, out io.Writer) error {
	return png.Encode(out, img)
}
