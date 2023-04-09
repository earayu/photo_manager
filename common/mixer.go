package common

import (
	"image"
)

// Mixer accepts a list of images and mix them into one image
type Mixer struct {
	DefaultOperator
	//processed images
	images  []*image.Image
	mixFunc func(images []*image.Image) (*image.Image, error)
}

func (d *Mixer) AddImages(images ...*image.Image) {
	d.images = append(d.images, images...)
}

func (d *Mixer) GetImages() []*image.Image {
	return d.images
}

func (d *Mixer) Mix() (*image.Image, error) {
	return d.mixFunc(d.images)
}

// CreateMixer returns a Mixer that mixes images using the given function
func CreateMixer(f func(images []*image.Image) (*image.Image, error)) *Mixer {
	return &Mixer{
		images:  make([]*image.Image, 0),
		mixFunc: f,
	}
}
