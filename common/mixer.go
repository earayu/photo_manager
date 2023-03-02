package common

import (
	"image"
)

type Mixer interface {
	Closer
	Mix(imagePool []*image.Image) (*image.Image, error)
}
