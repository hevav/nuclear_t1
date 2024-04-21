package pipes

import "hackathon/types"

type InterBank struct{}

func (InterBank) Proceed(tx *types.Transaction) (float64, error) {
	if tx.InterBank {
		return 0, nil
	} else {
		return 1, nil
	}
}
