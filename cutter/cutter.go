package cutter

import (
	"image"
	"image/jpeg"
	"os"
)

func CutImageByRatio(inputPath, outputPath string, widthWeight int, heightWeight int) (error, int, int) {
	targetRatio := float64(widthWeight) / float64(heightWeight)
	// Open input file
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err, 0, 0
	}
	defer inputFile.Close()

	// Decode input image
	inputImage, _, err := image.Decode(inputFile)
	if err != nil {
		return err, 0, 0
	}

	// Get input image dimensions
	inputWidth := inputImage.Bounds().Dx()
	inputHeight := inputImage.Bounds().Dy()

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
	croppedImage := inputImage.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(startX, startY, startX+outputWidth, startY+outputHeight))

	// Create output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err, 0, 0
	}
	defer outputFile.Close()

	// Encode output image
	err = jpeg.Encode(outputFile, croppedImage, &jpeg.Options{Quality: 80})
	if err != nil {
		return err, 0, 0
	}

	return nil, croppedImage.Bounds().Dx(), croppedImage.Bounds().Dy()
}

// cutImage cuts an image to the specified dimensions
func cutImage(inputPath, outputPath string, targetWidth, targetHeight int) (error, int, int) {
	// Open input file
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err, 0, 0
	}
	defer inputFile.Close()

	// Decode input image
	inputImage, _, err := image.Decode(inputFile)
	if err != nil {
		return err, 0, 0
	}

	// Get input image dimensions
	inputWidth := inputImage.Bounds().Dx()
	inputHeight := inputImage.Bounds().Dy()

	// Calculate output image dimensions
	outputWidth := targetWidth
	outputHeight := targetHeight
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
	croppedImage := inputImage.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(startX, startY, startX+outputWidth, startY+outputHeight))

	// Create output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err, 0, 0
	}
	defer outputFile.Close()

	// Encode output image
	err = jpeg.Encode(outputFile, croppedImage, &jpeg.Options{Quality: 80})
	if err != nil {
		return err, 0, 0
	}

	return nil, croppedImage.Bounds().Dx(), croppedImage.Bounds().Dy()
}
