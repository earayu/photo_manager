package common

import "image"

type Filter struct {
	DefaultOperator
	Lambda func(currentImage *image.Image) bool
}

// NextImage if returns nil, then the image will be filtered out in common/operator_chain.go#Process
func (f *Filter) NextImage(currentImage *image.Image) (*image.Image, error) {
	if f.Lambda(currentImage) {
		return currentImage, nil
	} else {
		return nil, nil
	}
}
