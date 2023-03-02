package common

type OperatorChain struct {
	DefaultOperator

	Ops []Operator

	//todo feature: add flag to control whether to save image to disk after each operator
	//each operator should also have a flag to control whether to save image to disk
}

func (o *OperatorChain) Process(inputPath, outputPath string) {
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
	err = o.Close(image, outputPath)
	if err != nil {
		panic(err)
	}
}
