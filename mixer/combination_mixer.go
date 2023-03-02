package mixer

import (
	"errors"
	"github.com/earayu/photo_manager/common"
	"image"
	"image/color"
)

type CombinationMixer struct {
	common.DefaultOperator
}

func (m *CombinationMixer) Mix(imagePool []*image.Image) (*image.Image, error) {
	// Check if imagePool is not empty
	if len(imagePool) == 0 {
		return nil, errors.New("imagePool is empty")
	}

	// Create a new RGBA image with dimensions equal to the first image in imagePool
	width := (*imagePool[0]).Bounds().Max.X
	height := (*imagePool[0]).Bounds().Max.Y
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))

	// Loop through each pixel in the new image and calculate the average color
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			var r, g, b, a uint32
			for _, img := range imagePool {
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
			numImages := uint32(len(imagePool))
			r /= numImages
			g /= numImages
			b /= numImages
			a /= numImages

			// Set the color of the current pixel in the new image to the average color
			rgba.SetRGBA64(x, y, color.RGBA64{uint16(r), uint16(g), uint16(b), uint16(a)})
		}
	}

	// Return the new image and nil error to indicate a successful mix
	return common.RgbaToImage(rgba), nil
}
