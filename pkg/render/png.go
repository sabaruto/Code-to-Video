package render

import (
	"errors"
	"image"
	"image/png"
	"log"
	"os"
)

type PNGRenderer struct {
	Width  int
	Height int
	canvas *image.NRGBA
}

func (r *PNGRenderer) initCanvas() {
	r.canvas = image.NewNRGBA(image.Rect(0, 0, r.Width, r.Height))
}

func (r *PNGRenderer) GetCanvas() *image.NRGBA {
	if r.canvas == nil {
		r.initCanvas()
	}

	return r.canvas
}

func (r *PNGRenderer) getSize() image.Point {
	if r.canvas == nil {
		r.initCanvas()
	}

	return r.canvas.Rect.Size()
}

func (r *PNGRenderer) Render(newImage image.NRGBA) error {
	if newImage.Rect.Size() != r.getSize() {
		return errors.New("incorrect size given")
	}

	for y := range r.Height {
		for x := range r.Width {
			r.canvas.Set(x, y, newImage.At(x, y))
		}
	}

	f, err := os.Create("image.png")
	if err != nil {
		return err
	}

	if err := png.Encode(f, r.canvas); err != nil {
		if closeErr := f.Close(); closeErr != nil {
			log.Fatal(closeErr)
		}

		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
