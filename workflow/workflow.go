package main

import (
	"github.com/earayu/photo_manager/common"
	"github.com/earayu/photo_manager/cutter"
	"github.com/earayu/photo_manager/resizer"
	flag "github.com/spf13/pflag"
	"image"
	"strings"
)

type CombinationWorkflow struct {
	SourceImagePool *common.FileSystemSource
	OperatorChain   *common.OperatorChain
	Mixer           *common.CombinationMixer
}

// NewWorkflow creates a new workflow and initializes all the components
func NewWorkflow(inputDir, ouputDir string) *CombinationWorkflow {

	source := &common.FileSystemSource{
		InputDir:  inputDir,
		OutputDir: ouputDir,
		SourceFilter: func(fileName string) bool {
			return strings.HasSuffix(fileName, ".jpg") || strings.HasSuffix(fileName, ".jpeg")
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

	mixer := &common.CombinationMixer{}

	return &CombinationWorkflow{
		source,
		chain,
		mixer,
	}
}

func (w *CombinationWorkflow) Run() {
	w.SourceImagePool.Open()
	for fileName, hasNext := w.SourceImagePool.Next(); hasNext; fileName, hasNext = w.SourceImagePool.Next() {
		//join w.SourceImagePool.InputDir and fileName
		w.OperatorChain.InputPath = w.SourceImagePool.InputDir + "/" + fileName
		w.OperatorChain.OutputPath = w.SourceImagePool.OutputDir + "/" + fileName
		w.OperatorChain.Process()
	}
	resultImage, err := w.Mixer.Mix(w.SourceImagePool.ProcessedImages)
	if err != nil {
		panic(err)
	}
	w.Mixer.Close(resultImage, w.SourceImagePool.OutputDir+"/heart2.jpg")
}

func main() {
	var baseDir string
	flag.StringVar(&baseDir, "base_dir", "~/Documents/GitHub/photo_manager", "The base directory of the project")
	flag.Parse()

	w := NewWorkflow(baseDir+"/src", baseDir+"/dest")
	w.Run()
}
