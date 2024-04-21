package pipes

import (
	"hackathon/types"
)

type Party2Category struct{}

func (Party2Category) Proceed(tx *types.Transaction) (float64, error) {
	vector, err := db.GetCategoryVector(tx.FromID, tx.Category)
	if err != nil {
		return 0, nil
	}

	vector.AggSum.Add(tx.Amount)

	return vector.AggSum.GetScore(), db.UpdateCategoryVector(vector)
}
