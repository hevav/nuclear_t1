package route

import (
	"bytes"
	"encoding/json"
	"hackathon/pipeline"
	"hackathon/types"
)

func proceed(r *bytes.Reader) error {
	var transaction types.Transaction
	err := json.NewDecoder(r).Decode(&transaction)
	if err != nil {
		return err
	}

	return pipeline.Proceed(&transaction)
}
