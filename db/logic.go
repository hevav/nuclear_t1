package db

import (
	"hackathon/types"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
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
