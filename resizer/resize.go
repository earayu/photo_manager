package resizer

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func ThumbnailImage(inputPath, outputPath string, maxWidth, maxHeight uint) (error, int, int) {
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

	// Determine output size while maintaining aspect ratio
	//inputWidth := inputImage.Bounds().Dx()
	//inputHeight := inputImage.Bounds().Dy()

	// Resize input image to output image
	outputImage := resize.Thumbnail(maxWidth, maxHeight, inputImage, resize.Lanczos3)

	// Create output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err, 0, 0
	}
	defer outputFile.Close()

	// Encode output image
	err = jpeg.Encode(outputFile, outputImage, &jpeg.Options{Quality: 80})
	if err != nil {
		return err, 0, 0
	}

	return nil, outputImage.Bounds().Dx(), outputImage.Bounds().Dy()
}
