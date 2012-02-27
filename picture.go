package main

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// Picture is a wrapper for the raw image, providing i.a reading image
// from the file.
type Picture struct {
	// Path to the image location.
	Path string
	// Raw image object.
	Image image.Image
	// The image format.
	Format string
}

// NewPicture allocates memory for the picture object.
//
// path - The image location path.
//
// Returns new picture.
func NewPicture(path string) *Picture {
	return &Picture{Path: path}
}

// Load reads image from the file.
//
// Returns an error if something went wrong.
func (pic *Picture) Load() (err error) {
	var file *os.File
	if file, err = os.Open(pic.Path); err != nil {
		return
	}
	pic.Image, pic.Format, err = image.Decode(file)
	return
}

// EachPixel is a helper for iterating over all the pixels of given image.
//
// img - The image to be iterated over.
// fn  - A callback executed on each position, takes coordinates and color
//       values as arguments.
//
// Example:
//
//     EachPixel(img, func(x, y int, r, g, b, a uint8) {
//         // do something with colors...
//     })
//
func EachPixel(img image.Image, fn func(x, y int, r, g, b, a uint8)) {
	rect := img.Bounds()
	for i := 0; i < rect.Dx()*rect.Dy(); i++ {
		x, y := i%rect.Dx()+rect.Min.X, i/rect.Dx()+rect.Min.Y
		pixel := img.At(x, y)
		r32, g32, b32, a32 := pixel.RGBA()
		r8, g8, b8, a8 := uint8(r32), uint8(g32), uint8(b32), uint8(a32)
		fn(x, y, r8, g8, b8, a8)
	}
}
