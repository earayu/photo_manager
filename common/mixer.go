package common

import (
	"image"
)

type Mixer interface {
	Mix(imagePool []*image.Image) (*image.Image, error)
}
