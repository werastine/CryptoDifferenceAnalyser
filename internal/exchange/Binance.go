// Package exchange contains the logic of request on Binance exchange
package exchange

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/werastine/CryptoDifferenceAnalyser/internal/market"
)

// ProviderBinance structure for interface
type ProviderBinance struct{}

type binanceTicker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// GetPrice func - public request
func (ProviderBinance) GetPrice(client *http.Client, Coin string) (market.CoinToReturn, error) {

	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%sUSDT", strings.ToUpper(Coin))

	response, err := client.Get(url)
	if err != nil {
		return market.CoinToReturn{}, fmt.Errorf("send a get request: %w", err)
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Println("fail to close the body", err)
		}
	}()

	if response.StatusCode != http.StatusOK {
		return market.CoinToReturn{}, fmt.Errorf("fail bad status %d", response.StatusCode)
	}

	var ticker binanceTicker
	if err := json.NewDecoder(response.Body).Decode(&ticker); err != nil {
		return market.CoinToReturn{}, fmt.Errorf("decoding json file: %w", err)
	}

	fprice, err := strconv.ParseFloat(ticker.Price, 64)
	if err != nil {
		return market.CoinToReturn{}, fmt.Errorf("converting string to float64: %w", err)
	}

	return market.CoinToReturn{Symbol: ticker.Symbol, Price: fprice, STExchange: "Binance"}, nil
}
