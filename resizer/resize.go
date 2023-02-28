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

func (t *ThumbnailResizer) NextImage(currentImage image.Image) (image.Image, error) {
	// Resize input image to output image
	outputImage := resize.Thumbnail(t.MaxWidth, t.MaxHeight, currentImage, resize.Lanczos3)
	return outputImage, nil
}

func ThumbnailImage(inputPath, outputPath string, maxWidth, maxHeight uint) (error, int, int) {
	t := ThumbnailResizer{
		MaxWidth:  uint(maxWidth),
		MaxHeight: uint(maxHeight),
	}
	image, err := t.Open(inputPath)
	if err != nil {
		return err, 0, 0
	}
	outputImage, err := t.NextImage(image)
	if err != nil {
		return err, 0, 0
	}
	t.Close(outputImage, outputPath)
	return nil, outputImage.Bounds().Dx(), outputImage.Bounds().Dy()
}
