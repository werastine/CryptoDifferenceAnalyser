// Package exchange contains the logic of request on HyperLiquid exchange
package exchange

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/werastine/CryptoDifferenceAnalyser/internal/market"
)

// ProviderHyperLiquid structure for interface
type ProviderHyperLiquid struct{}

// InfoRequest is requied request for hyperliquid
type infoRequest struct {
	Type string `json:"type"`
}

// GetPrice func - global price getter
func (ProviderHyperLiquid) GetPrice(client *http.Client, coin string) (market.CoinToReturn, error) {

	url := "https://api.hyperliquid.xyz/info"

	coin = strings.ToUpper(coin)

	requestBody := infoRequest{
		Type: "allMids",
	}

	jsonBytes, err := json.Marshal(requestBody)
	if err != nil {
		return market.CoinToReturn{}, fmt.Errorf("marhsaling json file: %w", err)
	}

	response, err := client.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return market.CoinToReturn{}, fmt.Errorf("post request: %w", err)
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Printf("closing response body: %v", err)
		}
	}()

	if response.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(response.Body)
		return market.CoinToReturn{}, fmt.Errorf("status code %d: %s", response.StatusCode, string(bodyBytes))
	}

	var prices map[string]string
	if err = json.NewDecoder(response.Body).Decode(&prices); err != nil {
		return market.CoinToReturn{}, fmt.Errorf("decoding json: %w", err)
	}

	price, ok := prices[coin]
	if !ok {
		return market.CoinToReturn{}, fmt.Errorf("searching %s in map", coin)
	}

	fprice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return market.CoinToReturn{}, fmt.Errorf("converting string to float64: %w", err)
	}

	return market.CoinToReturn{Price: fprice, Symbol: coin, STExchange: "HyperLiquid"}, nil
}
