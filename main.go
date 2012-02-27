package main

import (
	"net/http"
	"flag"
	"log"
)

var (
	// Server will be bound to this address.
	Addr string
	// Path to the folder with images.
	ImagesRoot string
)

func init() {
	flag.StringVar(&Addr, "addr", ":8090", "Server will be bound to this address")
	flag.StringVar(&ImagesRoot, "images-root", "./", "Path to the folder with images")
	flag.Parse()
}

func main() {
	srv := &http.Server{
		Addr:    Addr,
		Handler: &ImageHandler{},
	}
	log.Fatal(srv.ListenAndServe())
}