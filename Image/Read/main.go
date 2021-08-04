package main

import (
	"image"
	"image/color"

	// Prior registration of encoding and decoding functions
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Printf("Usage : %s <image file>\n", os.Args[0])
		os.Exit(0)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	// Do not forget
	// Use file.Seek(0, 0) when multiple functions access file

	// (optional) Image data
	imageConfig, imageConfigType, err := image.DecodeConfig(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Image config: %+v\n", imageConfig)
	log.Printf("Image config type: %+v\n", imageConfigType)

	file.Seek(0, 0)

	imageData, imageType, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	switch imageData.ColorModel() {
	case color.RGBAModel:
		log.Println("RGBA")
	case color.RGBA64Model:
		log.Println("RGBA64Model")
	case color.NRGBAModel:
		log.Println("NRGBAModel")
	case color.NRGBA64Model:
		log.Println("NRGBA64Model")
	case color.AlphaModel:
		log.Println("AlphaModel")
	case color.Alpha16Model:
		log.Println("Alpha16Model")
	case color.GrayModel:
		log.Println("GrayModel")
	case color.Gray16Model:
		log.Println("Gray16Model")
	case color.CMYKModel:
		log.Println("CMYKModel")
	case color.YCbCrModel:
		log.Println("YCbCrModel")
	case color.NYCbCrAModel:
		log.Println("NYCbCrAModel")
	default:
		if v := imageData.ColorModel().(color.Palette); v != nil {
			log.Println("Palette")
		} else {
			log.Printf("Unknown model: %+v\n", imageData.ColorModel())
		}
	}

	log.Printf("Image bounds: %+v\n", imageData.Bounds())
	log.Printf("Color at (0,0): %+v\n", imageData.At(0, 0))
	r, g, b, a := imageData.At(0, 0).RGBA()
	log.Printf("RGBA Color at (0,0): %d %d %d %d\n", r, g, b, a)

	switch imageType {
	case "png":
		log.Printf("PNG")
	case "jpeg":
		log.Printf("JPEG")
	case "gif":
		log.Printf("GIF")
	default:
		log.Printf("Unknown Image type: %s", imageType)
	}
}
