package operator

import (
	"github.com/muesli/smartcrop"
	"github.com/muesli/smartcrop/nfnt"
	"image"
)

type SmartCrop struct {
	DefaultOperator

	Width  int
	Height int
}

func (s *SmartCrop) NextImage(currentImage image.Image) (image.Image, error) {
	analyzer := smartcrop.NewAnalyzer(nfnt.NewDefaultResizer())
	topCrop, err := analyzer.FindBestCrop(currentImage, s.Width, s.Height)
	if err != nil {
		panic(err)
	}
	type SubImager interface {
		SubImage(r image.Rectangle) image.Image
	}
	croppedimg := currentImage.(SubImager).SubImage(topCrop)
	return croppedimg, nil
}
