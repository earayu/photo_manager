package common

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"image"
	"math/rand"
)

// compile
// convert a *image.RGBA to an image.Image
func RgbaToImage(rgba *image.RGBA) image.Image {
	return rgba.SubImage(image.Rect(rgba.Bounds().Min.X, rgba.Bounds().Min.Y, rgba.Bounds().Max.X, rgba.Bounds().Max.Y))
}

func Shuffle(imagePool []image.Image) []image.Image {
	return ShuffleN(imagePool, len(imagePool))
}

// shuffle the images in the imagePool
func ShuffleN(imagePool []image.Image, expectCount int) []image.Image {
	var shuffledImagePool []image.Image
	rand.Seed(timestamppb.Now().GetSeconds())
	for i := 0; i < expectCount; i++ {
		index := rand.Intn(len(imagePool))
		shuffledImagePool = append(shuffledImagePool, imagePool[index])
		imagePool = append(imagePool[:index], imagePool[index:]...)
		if len(imagePool) == 0 {
			imagePool = append(shuffledImagePool)
		}
	}
	return shuffledImagePool
}
