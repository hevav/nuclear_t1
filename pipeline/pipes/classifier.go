package pipes

import (
	"hackathon/db"
	"hackathon/types"
	"strings"
)

type Classifier struct{}

var classifierMap = map[string]types.TransactionCategory{
	"продукт":  types.TransactionCategoryGoods,
	"техник":   types.TransactionCategoryGoods,
	"отель":    types.TransactionCategoryTravel,
	"билет":    types.TransactionCategoryTravel,
	"ресторан": types.TransactionCategoryRestaurants,
	"кафе":     types.TransactionCategoryRestaurants,
	"бистро":   types.TransactionCategoryRestaurants,
	"комисс":   types.TransactionCategoryCommission,
	"сбор":     types.TransactionCategoryCommission,
	"налог":    types.TransactionCategoryCommission,
	"услуг":    types.TransactionCategoryUnclassifiedServices,
}

func (Classifier) Proceed(tx *types.Transaction) (float64, error) {
	var market types.Market
	var err error
	if tx.Category == types.TransactionCategoryRegularPurchase {
		if tx.To.Type == types.PartyTypeMarket {
			if market.ID == 0 {
				market, err = db.GetMarketByPartyID(tx.ToID)
				if err != nil {
					return 0, err
				}
			}

			if market.ID != 0 {
				tx.Category = market.Category
			}
		}
	}

	if tx.Category == types.TransactionCategoryRegularPurchase {
		tokens := strings.Split(tx.PurposeText, " ")
		doBreak := false
		for _, token := range tokens {
			for classificator, category := range classifierMap {
				if strings.HasPrefix(token, classificator) {
					tx.Category = category
					doBreak = true
					break
				}
			}

			if doBreak {
				break
			}
		}
	}

	if tx.Category == types.TransactionCategoryUnclassifiedServices {
		if tx.To.Type == types.PartyTypeMarket {
			if market.ID == 0 {
				market, err = db.GetMarketByPartyID(tx.ToID)
				if err != nil {
					return 0, err
				}
			}

			if market.Type == types.MarketTypeOnline {
				tx.Category = types.TransactionCategoryOnlineServices
			} else if market.Type == types.MarketTypeOffline {
				tx.Category = types.TransactionCategoryOfflineServices
			}
		}
	}

	return -1, nil
}
