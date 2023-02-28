package main

import (
	"github.com/earayu/photo_manager/common"
	"github.com/earayu/photo_manager/cutter"
	"github.com/earayu/photo_manager/resizer"
	flag "github.com/spf13/pflag"
)

type Heart2 struct {
	common.OperatorChain
}

func main() {
	var baseDir string
	flag.StringVar(&baseDir, "base_dir", "~/Documents/GitHub/photo_manager", "The base directory of the project")
	flag.Parse()
	ops := []common.Operator{
		&cutter.CutterByRatio{
			WidthWeight:  1,
			HeightWeight: 1,
		},
		&resizer.ThumbnailResizer{
			MaxWidth:  300,
			MaxHeight: 300,
		},
	}
	h := Heart2{
		common.OperatorChain{
			InputPath:  baseDir + "/testdata/resizer/resizer.jpeg",
			OutputPath: baseDir + "/testdata/resizer/resizer_output.jpeg",
			Ops:        ops,
		},
	}
	h.Process()
}
