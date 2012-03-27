package main

import (
	"image/png"
	"net/http"
	"os"
	"strings"
)

// FilterImageHandler applies requested filters to the pictures and
// serves its modified version.
//
// GET /filters?
//
// w - A response writer.
// r - Request to be performed.
//
func FilterImageHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	path, filterName := r.FormValue("pic"), r.FormValue("f")
	pic := NewPicture(path)

	if err := pic.Load(); err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	filters := strings.Split(filterName, ",")
	result := ApplyFiltersBatch(filters, pic.Image)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, result)
}
