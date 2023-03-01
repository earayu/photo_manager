package common

import (
	"errors"
	"image"
	"image/color"
)

//type HeartMixer struct {
//	common.DefaultMixer
//}
//
//// combine all images in ImagePool to one image, make it a heart shape
//func (m *HeartMixer) Mix() (*image.Image, error) {
//	panic("implement me")
//}

type HeartMixer struct {
	DefaultOperator
	ImagePool []*image.Image
}

func (m *HeartMixer) Mix() (*image.RGBA, error) {
	// Check if ImagePool is not empty
	if len(m.ImagePool) == 0 {
		return nil, errors.New("ImagePool is empty")
	}

	// Create a new RGBA image with dimensions equal to the first image in ImagePool
	width := (*m.ImagePool[0]).Bounds().Max.X
	height := (*m.ImagePool[0]).Bounds().Max.Y
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))

	// Loop through each pixel in the new image and calculate the average color
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			var r, g, b, a uint32
			for _, img := range m.ImagePool {
				// Retrieve the color of the corresponding pixel in the current image
				c := (*img).At(x, y)
				cr, cg, cb, ca := c.RGBA()

				// Accumulate the colors
				r += cr
				g += cg
				b += cb
				a += ca
			}

			// Divide the accumulated colors by the number of images to get the average color
			numImages := uint32(len(m.ImagePool))
			r /= numImages
			g /= numImages
			b /= numImages
			a /= numImages

			// Set the color of the current pixel in the new image to the average color
			rgba.SetRGBA64(x, y, color.RGBA64{uint16(r), uint16(g), uint16(b), uint16(a)})
		}
	}

	// Return the new image and nil error to indicate a successful mix
	return rgba, nil
}
