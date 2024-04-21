package pipeline

import "hackathon/types"

type Pipe interface {
	Proceed(transaction *types.Transaction) (float64, error)
}

type WeightedPipe struct {
	Pipe   Pipe
	Weight float64
}

var pipelineMap = map[string]Pipe{}
var pipeline = []WeightedPipe{}

func Init() {

}

func Proceed(transaction *types.Transaction) (float64, error) {
	result := 0.0

	pipesCount := len(pipeline)
	for _, pipe := range pipeline {
		pipeResult, err := pipe.Pipe.Proceed(transaction)
		if err != nil {
			return 0, err
		}

		if pipeResult != -1 {
			result += pipeResult * pipe.Weight
		} else {
			pipesCount -= 1
		}
	}

	result /= float64(pipesCount)
	return result, nil
}
