package main

import (
	"image"
	"image/color"
)

type SepiaFilter struct{}

func (f *SepiaFilter) Apply(orig image.Image) image.Image {
	res := image.NewRGBA(orig.Bounds())
	EachPixel(orig, func(x, y int, r, g, b, a uint8) {
		rt := (float32(r) * 0.393) + (float32(g) * 0.769) + (float32(b) * 0.189)
		gt := (float32(r) * 0.349) + (float32(g) * 0.686) + (float32(b) * 0.168)
		bt := (float32(r) * 0.272) + (float32(g) * 0.534) + (float32(b) * 0.131)
		if rt > 255 {
			rt = 255
		}
		if gt > 255 {
			gt = 255
		}
		if bt > 255 {
			bt = 255
		}
		c := color.RGBA{uint8(rt), uint8(gt), uint8(bt), a}
		res.Set(x, y, c)
	})
	return res
}

func init() {
	Filters["sepia"] = &SepiaFilter{}
}
