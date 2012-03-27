package main

import (
	"image"
)

// Filter is an interface representing all the filters.
type Filter interface {
	// Applies filter to the specified picture.
	Apply(image.Image) image.Image
}

// List of registered filters.
var Filters = map[string]Filter{}

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
// filterName - The name of the filter to be applied.
// img        - An image to be processed.
//
// Returns modified image.
func ApplyFilter(filterName string, img image.Image) (res image.Image) {
	filter, ok := Filters[filterName]
	if !ok {
		res = img
	} else {
		res = filter.Apply(img)
	}
	return
}
