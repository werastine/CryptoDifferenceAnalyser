// Package bybit contains the logic of request on ByBit exchange
package bybit

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type tickerByBit struct {
	Result struct {
		List []struct {
			Symbol    string `json:"symbol"`
			MarkPrice string `json:"markPrice"`
		} `json:"list"`
	} `json:"result"`
}

// CoinToReturn contains data about the coin wich user going to recieve
type CoinToReturn struct {
	Symbol     string
	Price      float64
	STExchange string
}

// GetByBitPrice - public request for coin price
func GetByBitPrice(client *http.Client, coin string) (CoinToReturn, error) {

	url := fmt.Sprintf("https://api.bybit.com/v5/market/tickers?category=inverse&symbol=%sUSDT", coin)

	resp, err := client.Get(url)
	if err != nil {
		return CoinToReturn{}, fmt.Errorf("send get request for %s: %w", coin, err)
	}

	defer func() {
		if errClose := resp.Body.Close(); errClose != nil {
			log.Printf("[ERROR] failed to close the body: %v", errClose)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return CoinToReturn{}, fmt.Errorf("status code of %s, %d", coin, resp.StatusCode)
	}

	var ticker tickerByBit
	if err := json.NewDecoder(resp.Body).Decode(&ticker); err != nil {
		return CoinToReturn{}, fmt.Errorf("decoding data of %s: %w", coin, err)
	}

	if len(ticker.Result.List) == 0 {
		return CoinToReturn{}, fmt.Errorf("no ticker data found in response")
	}

	tickerData := ticker.Result.List[0]

	priceF, err := strconv.ParseFloat(tickerData.MarkPrice, 64)
	if err != nil {
		return CoinToReturn{}, fmt.Errorf("converting price from string to float: %w", err)
	}

	return CoinToReturn{Symbol: tickerData.Symbol, Price: priceF, STExchange: "ByBit"}, nil
}
