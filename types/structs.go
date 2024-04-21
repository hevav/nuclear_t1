package types

import "time"

type Snowflaked struct {
	ID uint64 `gorm:"primaryKey"`
}

type Market struct {
	Snowflaked
	Name     string              `db:"name"`
	Type     MarketType          `db:"type"`
	Category TransactionCategory `db:"category"`
}

type Transaction struct {
	Snowflaked
	Timestamp   time.Time        `db:"timestamp"`
	Amount      int64            `db:"amount"`
	FromID      uint64           `db:"from_id"`
	From        TransactionParty `gorm:"foreignKey:FromID"`
	ToID        uint64           `db:"to_id"`
	To          TransactionParty `gorm:"foreignKey:ToID"`
	Flag        TransactionFlag  `db:"flag"`
	GeoLat      float64          `db:"geo_lat"`
	GeoLon      float64          `db:"geo_lon"`
	GeoZoneID   uint64           `db:"geo_zone_id"`
	GeoZone     TransactionPartyGeoZoneVector
	PurposeText string              `db:"purpose_text"`
	Category    TransactionCategory `db:"category"`
	InterBank   bool
	Suspiscious bool
}

type Profile struct {
	Snowflaked
	Name string      `db:"name"`
	Type ProfileType `db:"type"`
}

type TransactionPartyGeoZoneVector struct {
	Snowflaked
	TransactionPartyID uint64 `db:"transaction_party_id"`
	TransactionParty   TransactionParty
	GeoLat             float64 `db:"geo_lat"`
	GeoLon             float64 `db:"geo_lon"`
	GeoRadius          uint64  `db:"geo_radius"`
	AggSum
}

type TransactionPartyCategoryVector struct {
	TransactionPartyID uint64              `db:"transaction_party_id"`
	VectorCategory     TransactionCategory `db:"vector_category"`
	AggSum
}

type TransactionParty struct {
	Snowflaked
	Type    PartyType `db:"type"`
	PartyID uint64    `db:"party_id"`
	AggSum
}

type PartyAffiliates struct {
	FromID uint64  `db:"from_id"`
	ToID   uint64  `db:"to_id"`
	Score  float64 `db:"score"`
	AggSum
}
