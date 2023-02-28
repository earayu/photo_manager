package main

import (
	"github.com/earayu/photo_manager/common"
	"github.com/earayu/photo_manager/cutter"
	"github.com/earayu/photo_manager/resizer"
	flag "github.com/spf13/pflag"
	"image"
)

type Heart struct {
	common.DefaultOperator

	width  int
	height int
}

func (d *Heart) NextImage(currentImage image.Image) (image.Image, error) {
	c := cutter.CutterByRatio{
		WidthWeight:  1,
		HeightWeight: 1,
	}

	t := resizer.ThumbnailResizer{
		MaxWidth:  uint(d.width),
		MaxHeight: uint(d.height),
	}

	//call c.NextImage, then use the result to call t.NextImage
	if i, err := c.NextImage(currentImage); err != nil {
		return nil, err
	} else {
		return t.NextImage(i)
	}
}

func main() {
	var baseDir string
	flag.StringVar(&baseDir, "base_dir", "~/Documents/GitHub/photo_manager", "The base directory of the project")
	flag.Parse()
	h := Heart{
		width:  800,
		height: 800,
	}
	image, err := h.Open(baseDir + "/testdata/resizer/resizer.jpeg")
	if err != nil {
		panic(err)
	}
	outputImage, err := h.NextImage(image)
	if err != nil {
		panic(err)
	}
	err = h.Close(outputImage, baseDir+"/testdata/resizer/resizer_output.jpeg")
	if err != nil {
		panic(err)
	}
}
