package main

// https://www.devdungeon.com/content/working-images-go

import (
	"bytes"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func main() {
	createGif()
	createPng()
	createJpg()
}

func createGif() {
	colorPalette := []color.Color{
		color.RGBA{R: 255, G: 0, B: 0, A: 255},
		color.RGBA{R: 0, G: 255, B: 0, A: 255},
		color.RGBA{R: 0, G: 0, B: 255, A: 255},
		color.RGBA{R: 128, G: 128, B: 128, A: 255},
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
	}

	testImage := image.NewPaletted(image.Rect(0, 0, 200, 100), colorPalette)

	for x := 0; x < testImage.Rect.Dx(); x++ {
		for y := 0; y < testImage.Rect.Dy(); y++ {
			colorIndex := uint8(x * len(colorPalette) / testImage.Rect.Dx())
			testImage.SetColorIndex(x, y, colorIndex)
		}
	}

	outputFile, err := os.Create("test.gif")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	imageOptions := &gif.Options{}
	gif.Encode(outputFile, testImage, imageOptions)
}

func createPng() {
	testImage := image.NewRGBA(image.Rect(0, 0, 10, 10))

	outputFile, err := os.Create("test.png")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	for x := 0; x < testImage.Rect.Dx()/2; x++ {
		for y := 0; y < testImage.Rect.Dy()/2; y++ {
			testImage.Set(x, y, color.RGBA{G: 255, A: 255})
		}
	}

	var buff bytes.Buffer // buff.Bytes()

	png.Encode(&buff, testImage)

	buff.WriteTo(outputFile)
}

func createJpg() {
	testImage := image.NewRGBA(image.Rect(0, 0, 200, 100))

	outputFile, err := os.Create("test.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	for i := 0; i < len(testImage.Pix); i += 4 {
		var pixColor uint8 = 0
		if i > len(testImage.Pix)/2 {
			pixColor = 255
		}
		testImage.Pix[i] = 255        // Red
		testImage.Pix[i+1] = pixColor // Green
		testImage.Pix[i+2] = 128      // Blue
		testImage.Pix[i+3] = 255      // Alpha
	}

	imageOptions := &jpeg.Options{
		Quality: 95,
	}
	jpeg.Encode(outputFile, testImage, imageOptions)
}
