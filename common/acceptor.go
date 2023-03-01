package common

import "image"

type Acceptor struct {
	DefaultOperator
	Accept func(currentImage *image.Image)
}

func (a *Acceptor) NextImage(currentImage *image.Image) (*image.Image, error) {
	a.Accept(currentImage)
	return currentImage, nil
}
