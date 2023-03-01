package main

import (
	"github.com/earayu/photo_manager/common"
	"github.com/earayu/photo_manager/cutter"
	"github.com/earayu/photo_manager/resizer"
	flag "github.com/spf13/pflag"
	"image"
)

type Heart2 struct {
	common.OperatorChain
}

func main() {
	var baseDir string
	flag.StringVar(&baseDir, "base_dir", "~/Documents/GitHub/photo_manager", "The base directory of the project")
	flag.Parse()
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
}
