package pipes

import "hackathon/types"

type TransactionFlag struct{}

func (TransactionFlag) Proceed(tx *types.Transaction) (float64, error) {
	if tx.Flag == types.TransactionFlagNormal {
		return 0, nil
	} else {
		return 1, nil
	}
}
