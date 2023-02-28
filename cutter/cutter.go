package cutter

import (
	"github.com/earayu/photo_manager/common"
	"image"
)

type CutterByRatio struct {
	common.DefaultOperator

	WidthWeight  int
	HeightWeight int
}

func (c *CutterByRatio) NextImage(currentImage image.Image) (image.Image, error) {
	targetRatio := float64(c.WidthWeight) / float64(c.HeightWeight)

	// Get input image dimensions
	inputWidth := currentImage.Bounds().Dx()
	inputHeight := currentImage.Bounds().Dy()

	// Calculate output image dimensions
	inputRatio := float64(inputWidth) / float64(inputHeight)
	outputWidth := inputWidth
	outputHeight := inputHeight
	if inputRatio > targetRatio {
		outputWidth = int(float64(inputHeight) * targetRatio)
	} else {
		outputHeight = int(float64(inputWidth) / targetRatio)
	}

	// Calculate the starting point for the image crop
	startX := (inputWidth - outputWidth) / 2
	startY := (inputHeight - outputHeight) / 2

	// Create the cropped image
	croppedImage := currentImage.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(startX, startY, startX+outputWidth, startY+outputHeight))
	return croppedImage, nil
}

func CutImageByRatio(inputPath, outputPath string, widthWeight int, heightWeight int) (error, int, int) {
	c := CutterByRatio{
		WidthWeight:  widthWeight,
		HeightWeight: heightWeight,
	}
	image, err := c.Open(inputPath)
	if err != nil {
		return err, 0, 0
	}
	croppedImage, err := c.NextImage(image)
	if err != nil {
		return err, 0, 0
	}
	c.Close(croppedImage, outputPath)
	return nil, croppedImage.Bounds().Dx(), croppedImage.Bounds().Dy()
}

type CutterBySize struct {
	common.DefaultOperator

	targetWidth  int
	targetHeight int
}

func (c *CutterBySize) NextImage(currentImage image.Image) (image.Image, error) {
	// Get input image dimensions
	inputWidth := currentImage.Bounds().Dx()
	inputHeight := currentImage.Bounds().Dy()

	// Calculate output image dimensions
	outputWidth := c.targetWidth
	outputHeight := c.targetHeight
	if outputWidth > inputWidth {
		outputWidth = inputWidth
	}
	if outputHeight > inputHeight {
		outputHeight = inputHeight
	}

	// Calculate the starting point for the image crop
	startX := (inputWidth - outputWidth) / 2
	startY := (inputHeight - outputHeight) / 2

	// Create the cropped image
	croppedImage := currentImage.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(startX, startY, startX+outputWidth, startY+outputHeight))
	return croppedImage, nil
}

// cutImage cuts an image to the specified dimensions
func cutImage(inputPath, outputPath string, targetWidth, targetHeight int) (error, int, int) {
	c := CutterBySize{
		targetWidth:  targetWidth,
		targetHeight: targetHeight,
	}
	image, err := c.Open(inputPath)
	if err != nil {
		return err, 0, 0
	}
	croppedImage, err := c.NextImage(image)
	if err != nil {
		return err, 0, 0
	}
	c.Close(croppedImage, outputPath)
	return nil, croppedImage.Bounds().Dx(), croppedImage.Bounds().Dy()
}
