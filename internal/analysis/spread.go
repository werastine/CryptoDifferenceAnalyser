// Package analysis contains metrics
package analysis

import (
	"fmt"

	"github.com/werastine/CryptoDifferenceAnalyser/internal/market"
)

// SpreadData contains data of prices wich we recieved
type SpreadData struct {
	BuyPrice     float64
	BuyExchange  string
	BuyCoin      string
	SellPrice    float64
	SellExchange string
	SellCoin     string
}

// Spread returns buy price and sell price
func Spread(spread []market.CoinToReturn) (*SpreadData, error) {
	sd := SpreadData{}

	if len(spread) <= 1 {
		return &sd, fmt.Errorf("cannot count spread, recieved less than 2 exchanges")
	}

	for key := 0; key != len(spread); key++ {
		if sd.SellPrice == 0 && sd.BuyPrice == 0 {
			sd.BuyPrice = spread[key].Price
			sd.BuyExchange = spread[key].STExchange
			sd.BuyCoin = spread[key].Symbol

			sd.SellPrice = spread[key].Price
			sd.SellExchange = spread[key].STExchange
			sd.SellCoin = spread[key].Symbol
		}

		if sd.BuyPrice > spread[key].Price {
			sd.BuyPrice = spread[key].Price
			sd.BuyExchange = spread[key].STExchange
			sd.BuyCoin = spread[key].Symbol
		}

		if sd.SellPrice < spread[key].Price {
			sd.SellPrice = spread[key].Price
			sd.SellExchange = spread[key].STExchange
			sd.SellCoin = spread[key].Symbol
		}
	}

	return &sd, nil
}
