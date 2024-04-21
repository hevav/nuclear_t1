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

func (as AggSum) GetScore(limitsum, limitabs uint64, weightsum, weightabs, weightcnt float64) float64 {

}

func (as AggSum) Add(amount int64) {
	as.Sum += uint64(amount)
	if amount > 0 {
		as.AbsoluteSum += uint64(amount)
	} else {
		as.AbsoluteSum -= uint64(amount)
	}
	as.Count++
}
