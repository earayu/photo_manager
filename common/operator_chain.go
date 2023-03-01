package common

type OperatorChain struct {
	DefaultOperator

	InputPath  string
	OutputPath string

	Ops []Operator
}

func (o *OperatorChain) Process() {
	image, err := o.Open(o.InputPath)
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
	err = o.Close(image, o.OutputPath)
	if err != nil {
		panic(err)
	}
}
