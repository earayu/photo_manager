package common

import "image"

type Reducer interface {
	Reduce(currentImage *image.Image, nextImage *image.Image) (*image.Image, error)
}
