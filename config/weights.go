package config

import "hackathon/types"

var Weights WeightsType

type WeightsType struct {
	PipelineWeight map[string]float64
	AggSumWeights  struct {
		Party2Category map[types.TransactionCategory]types.AggSumWeight
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
			TransferTransfer uint32
			TransferWithdraw uint32
			CreditTransfer   uint32
			CreditWithdraw   uint32
		}
	}
}
