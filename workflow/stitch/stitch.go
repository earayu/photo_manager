package main

import (
	"github.com/earayu/photo_manager/common"
	"github.com/earayu/photo_manager/common/operator"
	"github.com/earayu/photo_manager/config"
	"github.com/earayu/photo_manager/mixer"
	"github.com/earayu/photo_manager/workflow"
	flag "github.com/spf13/pflag"
	"image"
	"strings"
)

type StitchWorkflow struct {
	workflow.Workflow
}

func NewStitchWorkflow(inputDir, ouputDir string) *StitchWorkflow {

	creator := mixer.StitchMixerCreator{
		PhotoCountInRowSide:    3,
		PhotoCountInColumnSide: 3,
	}
	mixer := creator.Create()

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
	}

	chain := &operator.OperatorChain{
		Ops: []operator.Operator{
			&operator.Filter{
				Filter: func(img image.Image) bool {
					return true
				},
			},
			&operator.CutterByRatio{
				WidthWeight:  1,
				HeightWeight: 1,
			},
			&operator.ThumbnailResizer{
				MaxWidth:  800,
				MaxHeight: 800,
			},
			&common.Acceptor{
				Accept: func(img image.Image) {
					mixer.AddImages(img)
				},
			},
		},
	}

	return &StitchWorkflow{
		workflow.Workflow{
			SourceImagePool: source,
			OperatorChain:   chain,
			Mixer:           mixer,
		},
	}
}

func main() {
	flag.Parse()

	w := NewStitchWorkflow(config.GlobalConfig.InputDir(), config.GlobalConfig.OutputDir())
	w.Run()

}
