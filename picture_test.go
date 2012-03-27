package main

import (
	"net/http"
	"testing"
)

func init() {
	go func() {
		http.Handle("/", http.FileServer(http.Dir("./pics")))
		http.ListenAndServe(":9999", nil)
	}()
}

func TestNewPicture(t *testing.T) {
	pic := NewPicture("./test.jpg")
	if pic.Path != "./test.jpg" {
		t.Errorf("Expected path to be ./test.jpg, got %s", pic.Path)
	}
}

func TestPictureLoadValidFile(t *testing.T) {
	pic := NewPicture("./pics/bubble.jpg")
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

func TestPictureLoadValidURL(t *testing.T) {
	pic := NewPicture("http://localhost:9999/bubble.jpg")
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

func TestPictureLoadInvalidURL(t *testing.T) {
	pic := NewPicture("http://notexistingpage.com/invalid.jpg")
	if err := pic.Load(); err == nil {
		t.Errorf("Expected to get an error while loading invalid file")
	}
}
