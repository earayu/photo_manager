package main

import (
	"github.com/earayu/photo_manager/common"
	"github.com/earayu/photo_manager/cutter"
	"github.com/earayu/photo_manager/resizer"
	flag "github.com/spf13/pflag"
	"image"
	"image/jpeg"
	"os"
)

type Heart2 struct {
	common.OperatorChain
}

func main() {
	var baseDir string
	flag.StringVar(&baseDir, "base_dir", "~/Documents/GitHub/photo_manager", "The base directory of the project")
	flag.Parse()
	//0. define HeartMixer
	hm := common.HeartMixer{
		ImagePool: make([]*image.Image, 0),
	}

	//1. define operators
	ops := []common.Operator{
		&common.Filter{
			Filter: func(img *image.Image) bool {
				return true
			},
		},
		&cutter.CutterByRatio{
			WidthWeight:  1,
			HeightWeight: 1,
		},
		&resizer.ThumbnailResizer{
			MaxWidth:  300,
			MaxHeight: 300,
		},
		&common.Acceptor{
			Accept: func(img *image.Image) {
				hm.ImagePool = append(hm.ImagePool, img)
			},
		},
	}

	//2. read input directory
	s := common.FileSystemSource{
		InputDir:  baseDir + "/src",
		OutputDir: baseDir + "/dest",
		SourceFilter: func(fileName string) bool {
			return true
		},
	}
	s.Open()

	for fileName, hasNext := s.Next(); hasNext; fileName, hasNext = s.Next() {
		//3. process each image
		h := Heart2{
			common.OperatorChain{
				InputPath:  baseDir + "/src/" + fileName,
				OutputPath: baseDir + "/dest/" + fileName,
				Ops:        ops,
			},
		}
		h.Process()
	}

	//4. mix images
	i, err := hm.Mix()
	if err != nil {
		panic(err)
	}
	Close(rgbaToImage(i), baseDir+"/dest/heart2.jpg")
}

func Close(currentImage image.Image, outputPath string) error {
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

func rgbaToImage(rgba *image.RGBA) image.Image {
	// Create a new *image.RGBA that is backed by the same pixel data as the original *image.RGBA,
	// but with a different color model.
	bounds := rgba.Bounds()
	newRgba := image.NewRGBA(bounds)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			newRgba.Set(x, y, rgba.At(x, y))
		}
	}

	// Return the new *image.RGBA as an *image.Image by using a type assertion.
	return image.Image(newRgba)
}
