package db

import (
	"hackathon/config"
	"hackathon/types"

	"github.com/sony/sonyflake"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var flake = sonyflake.NewSonyflake(sonyflake.Settings{})

func Init() {
	var err error
	// TODO: switch sqlite/postgre/mysql
	db, err = gorm.Open(sqlite.Open(config.Settings.Integration.LocalDatabase.Host), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&types.Market{})
	db.AutoMigrate(&types.PartyAffiliates{})
	db.AutoMigrate(&types.Profile{})
	db.AutoMigrate(&types.Transaction{})
	db.AutoMigrate(&types.TransactionParty{})
	db.AutoMigrate(&types.TransactionPartyGeoZoneVector{})
	db.AutoMigrate(&types.TransactionPartyCategoryVector{})
}

func GetPartyAffiliate(from_id, to_id uint64) (types.PartyAffiliates, error) {
	var affiliate types.PartyAffiliates
	result := db.Where("FromID =? AND ToID =?", from_id, to_id).First(&affiliate)
	if result.Error != nil {
		return affiliate, result.Error
	}
	return affiliate, nil
}

func GetCategoryVector(party_id uint64, category types.TransactionCategory) (types.TransactionPartyCategoryVector, error) {
	var vector types.TransactionPartyCategoryVector
	result := db.Where("PartyID =? AND Category =?", party_id, category).First(&vector)
	if result.Error != nil {
		return vector, result.Error
	}
	return vector, nil
}

func UpdateCategoryVector(vector *types.TransactionPartyCategoryVector) error {
	return db.Save(vector).Error
}

func UpdatePartyAffiliate(affiliate *types.PartyAffiliates) error {
	return db.Save(affiliate).Error
}

func GetMarketByPartyID(party_id uint64) (types.Market, error) {
	var market types.Market
	result := db.Where("ID =?", party_id).First(&market)
	if result.Error != nil {
		return market, result.Error
	}
	return market, nil
}

func GetGeoZonesByParty(party_id uint64) ([]types.TransactionPartyGeoZoneVector, error) {
	var zones []types.TransactionPartyGeoZoneVector
	result := db.Where("PartyID =?", party_id).Find(&zones)
	if result.Error != nil {
		return zones, result.Error
	}
	return zones, nil
}

func GetGeoZonePoints(zone_id uint64) ([]float64, error) {
	var zones []float64
	var txs []types.Transaction
	result := db.Where("GeoZoneID =?", zone_id).Find(&txs)
	if result.Error != nil {
		return zones, result.Error
	}

	for _, tx := range txs {
		zones = append(zones, tx.GeoLat, tx.GeoLon)
	}
	return zones, nil
}

func UpdateGeoZoneVector(vector *types.TransactionPartyGeoZoneVector) error {
	setSnowflakeID(&vector.Snowflaked)
	return db.Save(vector).Error
}

func CreateNewGeoZoneVector(tx *types.Transaction, vector *types.TransactionPartyGeoZoneVector) error {
	setSnowflakeID(&vector.Snowflaked)
	err := db.Save(vector).Error
	if err != nil {
		return err
	}

	tx.GeoZoneID = vector.ID
	setSnowflakeID(&tx.Snowflaked)
	return db.Save(tx).Error
}

func setSnowflakeID(sf *types.Snowflaked) {
	if sf.ID == 0 {
		sf.ID, _ = flake.NextID()
	}
}
