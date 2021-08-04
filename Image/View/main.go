package main

// https://socketloop.com/tutorials/golang-simple-image-viewer-with-go-gtk

import (

	// Prior registration of encoding and decoding functions
	"image"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	// Image viewer utils

	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
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

	imageData, imageType, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// v := imageData.(*image.Paletted) - needs conversion
	// v := imageData.(*image.YCbCr) - needs conversion
	// v := imageData.(*image.RGBA) - works without conversion
	bounds := imageData.Bounds()
	v := image.NewNRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(v, v.Bounds(), imageData, bounds.Min, draw.Src)

	log.Printf("Image config type: %+v\n", imageType)

	// Image drawing
	gtk.Init(nil)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Basic Image Viewer")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		gtk.MainQuit()
	})

	hbox := gtk.NewHBox(false, 1)

	hpaned := gtk.NewHPaned()
	hbox.Add(hpaned)

	// The image name may be used
	frame1 := gtk.NewFrame("")
	framebox1 := gtk.NewHBox(false, 1)
	frame1.Add(framebox1)

	hpaned.Pack1(frame1, false, false)

	// Image data drawing structure
	pixbufdata := gdkpixbuf.PixbufData{
		Data:          v.Pix,
		Colorspace:    gdkpixbuf.GDK_COLORSPACE_RGB,
		HasAlpha:      true,
		BitsPerSample: 8,
		Width:         v.Rect.Dx(),
		Height:        v.Rect.Dy(),
		RowStride:     v.Stride,
	}

	pixbuf := gdkpixbuf.NewPixbufFromData(pixbufdata)

	image := gtk.NewImageFromPixbuf(pixbuf)

	// image := gtk.NewImageFromFile(os.Args[1])
	framebox1.Add(image)

	window.Add(hbox)
	imagePixBuffer := image.GetPixbuf()
	horizontalSize := imagePixBuffer.GetWidth()
	verticalSize := imagePixBuffer.GetHeight()
	window.SetSizeRequest(horizontalSize, verticalSize)
	window.ShowAll()
	gtk.Main()
}
