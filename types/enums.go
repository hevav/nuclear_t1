package types

type MarketType int

const (
	MarketTypeUnclassified MarketType = iota
	MarketTypeOnline
	MarketTypeOffline
)

type TransactionFlag int

const (
	TransactionFlagNormal TransactionFlag = iota
	TransactionFlagReturn
)

type ProfileType int

const (
	ProfileTypePrivate ProfileType = iota
	ProfileTypeOrganization
)

type TransactionCategory int

const (
	TransactionCategoryOnlineEntertainment TransactionCategory = iota
	TransactionCategoryOfflineEntertainment
	TransactionCategoryRestaurants
	TransactionCategoryTravel
	TransactionCategoryGoods
	TransactionCategoryOnlineServices
	TransactionCategoryOfflineServices
	TransactionCategoryUnclassifiedServices
	TransactionCategoryCommission
	TransactionCategoryRegularPurchase
	TransactionCategoryCredit
	TransactionCategoryTransfer
	TransactionCategoryCash
)

type PartyType int

const (
	PartyTypeProfile PartyType = iota
	PartyTypeMarket
)
