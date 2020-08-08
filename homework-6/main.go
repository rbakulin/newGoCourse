package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
)

func hello(res http.ResponseWriter, req *http.Request) {
	name := "World"
	paramName := req.FormValue("name")
	if paramName != "" {
		name = req.FormValue("name")
	}
	res.Header().Set("Content-Type", "text/html")
	io.WriteString(res,
		fmt.Sprintf("<doctype html><html><head><title>Hello %s</title>" +
			"</head><body>Hello %s!</body></html>", name, name))
}

func main() {
	green := color.RGBA{0, 255, 0, 255}
	red := color.RGBA{200, 30, 30, 255}
	rectImg := image.NewRGBA(image.Rect(0, 0, 200, 200))

	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)
	// Lines
	for x := 0; x < 200; x++ {
		y := 100
		rectImg.Set(x, y, red)
	}
	for y := 0; y < 200; y++ {
		x := 100
		rectImg.Set(x, y, red)
	}

	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, rectImg)

	// HTTP-server
	fs := http.FileServer(http.Dir("homework-6/static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":80", nil)
}
