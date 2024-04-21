package main

type MarketType uint64

const (
	MarketTypeOffilne MarketType = iota
	MarketTypeOnline
)

type Market struct {
	Name string
	Type MarketType
}

type Transaction struct {
	Timestamp uint64
	Amount    int64
	MarketID  uint64
}

type Profile struct {
	Name string
}

func main() {

}
