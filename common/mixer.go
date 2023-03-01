package common

import "image"

type Mixer interface {
	Mix() (*image.Image, error)
}

type DefaultMixer struct {
	ImagePool []*image.Image
}

func (m *DefaultMixer) Mix() (*image.Image, error) {
	panic("implement me")
}
