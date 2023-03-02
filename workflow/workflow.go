package workflow

import (
	"github.com/earayu/photo_manager/common"
	"github.com/earayu/photo_manager/config"
)

type Workflow struct {
	SourceImagePool common.Source
	OperatorChain   *common.OperatorChain
	Mixer           common.Mixer
}

func (w *Workflow) Run() {
	w.SourceImagePool.Open()
	for fileName, hasNext := w.SourceImagePool.Next(); hasNext; fileName, hasNext = w.SourceImagePool.Next() {
		//join w.SourceImagePool.InputDir and fileName
		sourceFileName := config.GlobalConfig.InputDir() + "/" + fileName
		processedFileName := config.GlobalConfig.OutputDir() + "/" + fileName
		w.OperatorChain.Process(sourceFileName, processedFileName)
	}
	resultImage, err := w.Mixer.Mix(w.SourceImagePool.GetProcessedImages())
	if err != nil {
		panic(err)
	}
	w.Mixer.Close(resultImage, config.GlobalConfig.FullNameTargetFileName("jpeg"))
}
