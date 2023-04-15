package resizer

import (
	"github.com/earayu/photo_manager/common"
	"github.com/nfnt/resize"
	"image"
)

type ThumbnailResizer struct {
	common.DefaultOperator

	MaxWidth  uint
	MaxHeight uint
}

func (t *ThumbnailResizer) NextImage(currentImage *image.Image) (*image.Image, error) {
	// Resize input image to output image
	outputImage := resize.Thumbnail(t.MaxWidth, t.MaxHeight, (*currentImage), resize.Lanczos3)
	return &outputImage, nil
}
