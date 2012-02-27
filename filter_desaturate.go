package main

import (
	"image"
	"image/color"
)

func DesaturateFilter(orig image.Image) image.Image {
	res := image.NewGray(orig.Bounds())
	EachPixel(orig, func(x, y int, r, g, b, a uint8) {
		gray := (0.2126 * float32(r)) + (0.7152 * float32(g)) + (0.0722 * float32(b))
		res.Set(x, y, color.Gray{uint8(gray)})
	})
	return res
}
