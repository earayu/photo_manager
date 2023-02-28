package common

import (
	"image"
	"image/jpeg"
	"os"
)

type Operator interface {
	Open(inputPath string) (image.Image, error)
	NextImage(currentImage image.Image) (image.Image, error)
	Close(currentImage image.Image, outputPath string) error
}

type DefaultOperator struct {
}

func (d *DefaultOperator) Open(inputPath string) (image.Image, error) {
	// Open input file
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	// Decode input image
	inputImage, _, err := image.Decode(inputFile)
	if err != nil {
		return nil, err
	}
	return inputImage, nil
}

func (d *DefaultOperator) NextImage(currentImage image.Image) (image.Image, error) {
	panic("implement me")
}

func (d *DefaultOperator) Close(currentImage image.Image, outputPath string) error {
	// Create output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Encode output image
	err = jpeg.Encode(outputFile, currentImage, &jpeg.Options{Quality: 80})
	if err != nil {
		return err
	}
	return nil
}
