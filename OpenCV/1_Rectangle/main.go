package main

// xhost +
// docker run -it --net=host --ipc=host -e DISPLAY=$DISPLAY -v /tmp/.X11-unix:/tmp/.X11-unix gocv /bin/bash

import (
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

func main() {
	window := gocv.NewWindow("test")
	img := gocv.NewMatWithSize(200, 200, gocv.MatTypeCV8UC3)

	gocv.Rectangle(&img, image.Rectangle{Min: image.Point{X: 60, Y: 60}, Max: image.Point{X: 140, Y: 140}}, color.RGBA{R: 0, G: 0, B: 255}, 1)
	gocv.Rectangle(&img, image.Rectangle{Min: image.Point{X: 50, Y: 50}, Max: image.Point{X: 150, Y: 150}}, color.RGBA{R: 0, G: 255, B: 0}, 1)
	gocv.Rectangle(&img, image.Rectangle{Min: image.Point{X: 40, Y: 40}, Max: image.Point{X: 160, Y: 160}}, color.RGBA{R: 255, G: 0, B: 0}, 1)

	window.IMShow(img)
	window.WaitKey(0)
}
