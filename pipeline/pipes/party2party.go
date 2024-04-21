package pipes

import (
	"hackathon/types"
)

type Party2Party struct{}

func (Party2Party) Proceed(tx *types.Transaction) (float64, error) {
	affiliate, err := db.GetPartyAffiliate(tx.FromID, tx.ToID)
	if err != nil {
		return 0, nil
	}

	affiliate.AggSum.Add(tx.Amount)

	return affiliate.AggSum.GetScore(), db.UpdatePartyAffiliate(affiliate)
}
