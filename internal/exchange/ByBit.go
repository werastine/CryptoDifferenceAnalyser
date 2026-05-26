// Package exchange contains the logic of request on ByBit exchange
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

// ProviderByBit structure for interface
type ProviderByBit struct {
	client *http.Client
}

func NewProviderByBit(cl *http.Client) *ProviderByBit {
	return &ProviderByBit{client: cl}
}

type tickerByBit struct {
	Result struct {
		List []struct {
			Symbol    string `json:"symbol"`
			MarkPrice string `json:"markPrice"`
		} `json:"list"`
	} `json:"result"`
}

// GetPrice - public request for coin price
func (b *ProviderByBit) GetPrice(coin string) (market.CoinToReturn, error) {

	url := fmt.Sprintf("https://api.bybit.com/v5/market/tickers?category=inverse&symbol=%sUSDT", strings.ToUpper(coin))

	resp, err := b.client.Get(url)
	if err != nil {
		return market.CoinToReturn{}, fmt.Errorf("send get request for %s: %w", coin, err)
	}

	defer func() {
		if errClose := resp.Body.Close(); errClose != nil {
			log.Printf("[ERROR] failed to close the body: %v", errClose)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return market.CoinToReturn{}, fmt.Errorf("status code of %s, %d", coin, resp.StatusCode)
	}

	var ticker tickerByBit
	if err := json.NewDecoder(resp.Body).Decode(&ticker); err != nil {
		return market.CoinToReturn{}, fmt.Errorf("decoding data of %s: %w", coin, err)
	}

	if len(ticker.Result.List) == 0 {
		return market.CoinToReturn{}, fmt.Errorf("no ticker data found in response")
	}

	tickerData := ticker.Result.List[0]

	priceF, err := strconv.ParseFloat(tickerData.MarkPrice, 64)
	if err != nil {
		return market.CoinToReturn{}, fmt.Errorf("converting price from string to float: %w", err)
	}

	return market.CoinToReturn{Symbol: tickerData.Symbol, Price: priceF, STExchange: "ByBit"}, nil
}
