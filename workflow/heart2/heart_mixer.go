package main

import (
	"github.com/earayu/photo_manager/common"
	"image"
)

type HeartMixer struct {
	common.DefaultMixer
}

// combine all images in ImagePool to one image, make it a heart shape
func (m *HeartMixer) Mix() (*image.Image, error) {
	panic("implement me")
}
