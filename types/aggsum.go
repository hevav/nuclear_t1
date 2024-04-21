package types

type AggSum struct {
	Sum         uint64 `db:"sum"`
	AbsoluteSum uint64 `db:"absolute_sum"`
	Count       uint64 `db:"count"`
}

type AggSumWeight struct {
	SumWeight         float64
	AbsoluteSumWeight float64
	CountWeight       float64
	SumLimit          uint64
}

func (as *AggSum) GetScore(asw AggSumWeight) float64 {
	weightedSum := float64(as.Sum)*asw.SumWeight + float64(as.AbsoluteSum)*asw.AbsoluteSumWeight + float64(as.Count)*asw.CountWeight
	riskScore := 1 - (weightedSum / float64(asw.SumLimit))

	if riskScore < 0 {
		riskScore = 0
	} else if riskScore > 1 {
		riskScore = 1
	}

	return riskScore
}

func (as *AggSum) Add(amount int64) {
	as.Sum += uint64(amount)
	if amount > 0 {
		as.AbsoluteSum += uint64(amount)
	} else {
		as.AbsoluteSum -= uint64(amount)
	}
	as.Count++
}
