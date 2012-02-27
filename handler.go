package main

import (
	"image"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Handler is an HTTP handler for serving filtered images.
type Handler struct{}

// ServeHTTP applies requested filters to the pictures and serves its
// modified version.
//
// w - A response writer.
// r - Request to be performed.
//
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(ImagesRoot, r.URL.Path)
	pic := NewPicture(path)

	if err := pic.Load(); err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	r.ParseForm()
	result := ApplyFiltersBatch(strings.Split(r.FormValue("filter"), ","), pic.Image)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, result)
}

// ApplyFiltersBatch applies list of filters to the specified image.
//
// filters - The list of filters to apply.
// img     - An image to be processed.
//
// Returns modified image.
func ApplyFiltersBatch(filters []string, img image.Image) (res image.Image) {
	res = img
	for _, filter := range filters {
		res = ApplyFilter(filter, res)
	}
	return
}

// ApplyFilter applies specified filter to the image.
//
// filter - The name of the filter to be applied.
// img    - An image to be processed.
//
// Returns modified image.
func ApplyFilter(filter string, img image.Image) (res image.Image) {
	switch filter {
	case "desaturate":
		res = DesaturateFilter(img)
	case "sepia":
		res = SepiaFilter(img)
	default:
		res = img
	}
	return
}
