package config

import "hackathon/types"

var Weights WeightsType

type WeightsType struct {
	PipelineWeight map[string]float64
	AggSumWeights  struct {
		Party2Category map[string]types.AggSumWeight
		Party2Party    types.AggSumWeight
		GeoZone        types.AggSumWeight
		FastOps        struct {
			TransferTransfer types.AggSumWeight
			TransferWithdraw types.AggSumWeight
			CreditTransfer   types.AggSumWeight
			CreditWithdraw   types.AggSumWeight
		}
	}
	TimeWeights struct {
		FastOps struct {
			TransferTransfer float64
			TransferWithdraw float64
			CreditTransfer   float64
			CreditWithdraw   float64
		}
	}
}
