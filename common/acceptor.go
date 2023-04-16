package common

import (
	"github.com/earayu/photo_manager/common/operator"
	"image"
)

type Acceptor struct {
	operator.DefaultOperator
	Accept func(currentImage image.Image)
}

func (a *Acceptor) NextImage(currentImage image.Image) (image.Image, error) {
	a.Accept(currentImage)
	return currentImage, nil
}
