package operator

import (
	"github.com/nfnt/resize"
	"image"
)

type ThumbnailResizer struct {
	DefaultOperator

	MaxWidth  uint
	MaxHeight uint
}

func (t *ThumbnailResizer) NextImage(currentImage *image.Image) (*image.Image, error) {
	// Resize input image to output image
	outputImage := resize.Thumbnail(t.MaxWidth, t.MaxHeight, (*currentImage), resize.Lanczos3)
	return &outputImage, nil
}
