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

// NewCombinationWorkflow creates a new workflow and initializes all the components.
func NewCombinationWorkflow(inputDir, ouputDir string) *CombinationWorkflow {

	source := &common.FileSystemSource{
		InputDir:  inputDir,
		OutputDir: ouputDir,
		SourceFilter: func(fileName string) bool {
			//whether file name ends with element in sourceType
			for _, suffix := range sourceType {
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
		sourceFileName := w.SourceImagePool.InputDir + "/" + fileName
		processedFileName := w.SourceImagePool.OutputDir + "/" + fileName
		w.OperatorChain.Process(sourceFileName, processedFileName)
	}
	resultImage, err := w.Mixer.Mix(w.SourceImagePool.ProcessedImages)
	if err != nil {
		panic(err)
	}
	w.Mixer.Close(resultImage, w.SourceImagePool.OutputDir+"/"+targetFileName+".jpeg")
}

func main() {
	flag.StringVar(&inputDir, "input_dir", "~/Documents/GitHub/photo_manager/testdata/resizer", "The base directory of the project")
	flag.StringVar(&outputDir, "output_dir", "", "The base directory of the project")
	flag.StringVar(&targetFileName, "target_file_name", "target", "The base directory of the project")
	flag.StringArrayVar(&sourceType, "source_type", []string{"jpg", "jpeg", "png"}, "The base directory of the project")
	flag.Parse()

	if outputDir == "" {
		outputDir = inputDir + "/output"
	}
	w := NewCombinationWorkflow(inputDir, outputDir)
	w.Run()
}
