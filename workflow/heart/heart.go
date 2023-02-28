package main

import (
	"github.com/earayu/photo_manager/cutter"
	"github.com/earayu/photo_manager/resizer"
	flag "github.com/spf13/pflag"
)

func main() {
	var baseDir string
	flag.StringVar(&baseDir, "base_dir", "~/Documents/GitHub/photo_manager", "The base directory of the project")
	flag.Parse()
	err, x, y := cutter.CutImageByRatio(
		baseDir+"/testdata/resizer/resizer.jpeg",
		baseDir+"/testdata/resizer/resizer_output.jpeg",
		1,
		1,
	)
	if err != nil {
		panic(err)
	}
	println(x, y)

	err, x, y = resizer.ThumbnailImage(
		baseDir+"/testdata/resizer/resizer_output.jpeg",
		baseDir+"/testdata/resizer/resizer_output.jpeg",
		1500,
		1500,
	)
	if err != nil {
		panic(err)
	}

	println(x, y)
}
