package common

import (
	"image"
	"image/draw"
)

type Mixer interface {
	Mix() (*image.Image, error)
}

type DefaultMixer struct {
	ImagePool []*image.Image
}

func (m *DefaultMixer) Mix() (*image.RGBA, error) {
	// Determine the maximum width and height of the images in the pool
	maxWidth := 0
	maxHeight := 0
	for _, img := range m.ImagePool {
		if (*img).Bounds().Dx() > maxWidth {
			maxWidth = (*img).Bounds().Dx()
		}
		if (*img).Bounds().Dy() > maxHeight {
			maxHeight = (*img).Bounds().Dy()
		}
	}

	// Create a new RGBA image to hold the combined images
	finalImg := image.NewRGBA(image.Rect(0, 0, maxWidth*len(m.ImagePool), maxHeight))

	// Combine the images into the final image
	for i, img := range m.ImagePool {
		// Calculate the position of the current image in the final image
		pos := image.Point{X: i * maxWidth, Y: 0}
		// Draw the current image onto the final image
		draw.Draw(finalImg, (*img).Bounds().Add(pos), *img, image.ZP, draw.Src)
	}

	return finalImg, nil
}
