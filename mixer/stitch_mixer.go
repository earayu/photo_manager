package mixer

import (
	"errors"
	"github.com/earayu/photo_manager/common"
	"image"
	"image/draw"
)

type StitchMixer struct {
	common.DefaultOperator
}

func (m *StitchMixer) Mix(imagePool []*image.Image) (*image.Image, error) {
	// Make sure we have at least two images to stitch
	if len(imagePool) < 2 {
		return nil, errors.New("need at least two images to stitch")
	}

	// Get the size of the first image and calculate the size of the stitched image
	w := (*imagePool[0]).Bounds().Max.X
	h := (*imagePool[0]).Bounds().Max.Y
	for _, img := range imagePool[1:] {
		bounds := (*img).Bounds()
		if bounds.Max.X > w {
			w = bounds.Max.X
		}
		if bounds.Max.Y > h {
			h = bounds.Max.Y
		}
	}

	// Create the output image
	outImg := image.NewRGBA(image.Rect(0, 0, w*len(imagePool), h))

	// Draw each image onto the output image
	for i, img := range imagePool {
		bounds := (*img).Bounds()
		x := i * w
		y := 0
		r := image.Rect(x, y, x+bounds.Max.X, y+bounds.Max.Y)
		draw.Draw(outImg, r, *img, bounds.Min, draw.Src)
	}

	// Return the new image and nil error to indicate a successful mix
	return common.RgbaToImage(outImg), nil
}
