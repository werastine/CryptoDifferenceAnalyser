// Package binancemrkt contains the logic of request on Binance exchange
package binancemrkt

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type binanceTicker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// CoinToReturn contains data about the coin wich user going to recieve
type CoinToReturn struct {
	Coin       string
	Price      float64
	STExchange string
}

// GetPriceBinance func - public request
func GetPriceBinance(client *http.Client, Coin string) (CoinToReturn, error) {

	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%sUSDT", strings.ToUpper(Coin))

	response, err := client.Get(url)
	if err != nil {
		return CoinToReturn{}, fmt.Errorf("send a get request: %w", err)
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Println("fail to close the body", err)
		}
	}()

	if response.StatusCode != http.StatusOK {
		return CoinToReturn{}, fmt.Errorf("fail bad status %d", response.StatusCode)
	}

	var ticker binanceTicker
	if err := json.NewDecoder(response.Body).Decode(&ticker); err != nil {
		return CoinToReturn{}, fmt.Errorf("decoding json file: %w", err)
	}

	fprice, err := strconv.ParseFloat(ticker.Price, 64)
	if err != nil {
		return CoinToReturn{}, fmt.Errorf("converting string to float64: %w", err)
	}

	return CoinToReturn{Coin: ticker.Symbol, Price: fprice, STExchange: "Binance"}, nil
}
