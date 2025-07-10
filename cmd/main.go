package main

import (
	"image"
	"image/color"
	"log"

	"github.com/sabaruto/code-to-video/pkg/render"
)

func main() {
	const width, height = 256, 256

	// Create a colored image of the given width and height.
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	renderer := render.PNGRenderer{
		Width:  256,
		Height: 256,
	}

	for y := range height {
		for x := range width {
			// img.Set(x, y, color.NRGBA{
			// 	R: uint8((x + y) << 0 & 255),
			// 	G: uint8((x + y) << 1 & 255),
			// 	B: uint8((x + y) << 2 & 255),
			// 	A: 255,
			// })
			img.Set(x, y, color.NRGBA{
				R: 10,
				G: 0,
				B: 0,
				A: 255,
			})
		}
	}

	if err := renderer.Render(*img); err != nil {
		log.Fatal(err)
	}
}
