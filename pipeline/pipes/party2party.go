package pipes

import (
	"hackathon/config"
	"hackathon/db"
	"hackathon/types"
)

type Party2Party struct{}

func (Party2Party) Proceed(tx *types.Transaction) (float64, error) {
	affiliate, err := db.GetPartyAffiliate(tx.FromID, tx.ToID)
	if err != nil {
		return 0, nil
	}

	affiliate.AggSum.Add(tx.Amount)

	return affiliate.AggSum.GetScore(config.Weights.AggSumWeights.Party2Party) * affiliate.Score, db.UpdatePartyAffiliate(&affiliate)
}
