package common

import "sync"

type OperatorChain struct {
	DefaultOperator

	Ops []Operator

	//SkipClose default true
	AutoSave bool
	Wg       sync.WaitGroup
}

func (o *OperatorChain) Process(inputPath, outputPath string) {
	defer o.Wg.Done()
	image, err := o.Open(inputPath)
	if err != nil {
		panic(err)
	}
	for _, operator := range o.Ops {
		if image == nil {
			break
		}
		image, err = operator.NextImage(image)
		if err != nil {
			panic(err)
		}
	}
	if image == nil {
		return
	}
	if !o.AutoSave {
		return
	}
	err = o.Close(image, outputPath)
	if err != nil {
		panic(err)
	}
}

// wait for all process to finish
func (o *OperatorChain) Wait() {
	o.Wg.Wait()
}
