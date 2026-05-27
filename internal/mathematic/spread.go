// Package math
package mathematic

import (
	"github.com/werastine/CryptoDifferenceAnalyser/internal/market"
)

type spreadData struct {
	BuyPrice     float64
	BuyExchange  string
	BuyCoin      string
	SellPrice    float64
	SellEcchange string
	SellCoin     string
}

// Spread returns buy price and sell price
func Spread(spread map[market.CoinToReturn]struct{}) *spreadData {
	sd := spreadData{}

	for key := range spread {
		if sd.SellPrice == 0 && sd.BuyPrice == 0 {
			sd.BuyPrice = key.Price
			sd.BuyExchange = key.STExchange
			sd.BuyCoin = key.Symbol

			sd.SellPrice = key.Price
			sd.SellEcchange = key.STExchange
			sd.SellCoin = key.Symbol
		}

		if sd.BuyPrice > key.Price {
			sd.BuyPrice = key.Price
			sd.BuyExchange = key.STExchange
			sd.BuyCoin = key.Symbol
		}

		if sd.SellPrice < key.Price {
			sd.SellPrice = key.Price
			sd.SellEcchange = key.STExchange
			sd.SellCoin = key.Symbol
		}
	}

	return &sd
}
