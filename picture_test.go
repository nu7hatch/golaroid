package main

import (
	"testing"
)

func TestNewPicture(t *testing.T) {
	pic := NewPicture("./test.jpg")
	if pic.Path != "./test.jpg" {
		t.Errorf("Expected path to be ./test.jpg, got %s", pic.Path)
	}
}

func TestPictureLoadValidFile(t *testing.T) {
	pic := NewPicture("./bubble.jpg")
	if err := pic.Load(); err != nil {
		t.Errorf("Expected to load file without errors, got: %v", err)
	}
	if pic.Image == nil {
		t.Errorf("Expected image to be loaded")
	}
	if pic.Format != "jpeg" {
		t.Errorf("Expected image to be in correct format, got: %s", pic.Format)
	}
}

func TestPictureLoadInvalidFile(t *testing.T) {
	pic := NewPicture("./invalid.jpg")
	if err := pic.Load(); err == nil {
		t.Errorf("Expected to get an error while loading invalid file")
	}
}
