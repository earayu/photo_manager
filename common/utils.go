package common

import (
	"image"
)

// compile
// convert a *image.RGBA to an *image.Image
func RgbaToImage(rgba *image.RGBA) *image.Image {
	i := rgba.SubImage(image.Rect(rgba.Bounds().Min.X, rgba.Bounds().Min.Y, rgba.Bounds().Max.X, rgba.Bounds().Max.Y))
	return &i
}
