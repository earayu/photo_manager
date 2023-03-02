package main

import (
	"github.com/earayu/photo_manager/common"
	"github.com/earayu/photo_manager/config"
	"github.com/earayu/photo_manager/cutter"
	"github.com/earayu/photo_manager/mixer"
	"github.com/earayu/photo_manager/resizer"
	"github.com/earayu/photo_manager/workflow"
	flag "github.com/spf13/pflag"
	"image"
	"strings"
)

type CombinationWorkflow struct {
	workflow.Workflow
}

// NewCombinationWorkflow creates a new workflow and initializes all the components.
func NewCombinationWorkflow(inputDir, ouputDir string) *CombinationWorkflow {

	source := &common.FileSystemSource{
		InputDir:  inputDir,
		OutputDir: ouputDir,
		SourceFilter: func(fileName string) bool {
			//whether file name ends with element in sourceType
			for _, suffix := range config.GlobalConfig.SourceType() {
				if strings.HasSuffix(fileName, "."+suffix) {
					return true
				}
			}
			return false
		},
		ProcessedImages: make([]*image.Image, 0),
	}

	chain := &common.OperatorChain{
		Ops: []common.Operator{
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
					source.ProcessedImages = append(source.ProcessedImages, img)
				},
			},
		},
	}

	mixer := &mixer.CombinationMixer{}

	return &CombinationWorkflow{
		workflow.Workflow{
			SourceImagePool: source,
			OperatorChain:   chain,
			Mixer:           mixer,
		},
	}
}

func main() {
	flag.Parse()

	w := NewCombinationWorkflow(config.GlobalConfig.InputDir(), config.GlobalConfig.OutputDir())
	w.Run()

}
