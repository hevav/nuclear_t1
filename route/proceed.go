package route

import (
	"encoding/json"
	"hackathon/pipeline"
	"hackathon/types"
	"io"
)

type Result struct {
	Score       float64
	Suspiscious bool
}

func proceed(r io.Reader, w io.Writer) error {
	var transaction types.Transaction
	err := json.NewDecoder(r).Decode(&transaction)
	if err != nil {
		return err
	}

	fraudScore, err := pipeline.Proceed(&transaction)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(Result{
		Score:       fraudScore,
		Suspiscious: transaction.Suspiscious,
	})
}
