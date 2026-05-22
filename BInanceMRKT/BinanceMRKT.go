// Package binancemrkt contains the logic of request on Binance exchange
package binancemrkt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type url struct {
	url string
}

// Func to add Coin name and USDT into link
func insert(url string, idx int, insertion string) string {
	res := url[:idx] + insertion + "USDT" + url[idx:]
	return res
}

func newURL() url {
	return url{url: "https://api.binance.com/api/v3/ticker/price?symbol="}
}

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

	urlData := newURL()
	urlData.url = insert(urlData.url, 51, Coin)

	response, err := client.Get(urlData.url)
	if err != nil {
		fmt.Println("Fail got no answer", err)
		return CoinToReturn{}, err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			fmt.Println("fail to close the body", err)
		}
	}()

	if response.StatusCode != http.StatusOK {
		err := fmt.Errorf("fail bad status %d", response.StatusCode)
		return CoinToReturn{}, err
	}

	var ticker binanceTicker
	if err := json.NewDecoder(response.Body).Decode(&ticker); err != nil {
		fmt.Println("Got an error in decoding response.Body")
		return CoinToReturn{}, err
	}

	fprice, err := strconv.ParseFloat(ticker.Price, 64)
	if err != nil {
		fmt.Println("Fail to convert string to float64", err)
		return CoinToReturn{}, err
	}

	coin := CoinToReturn{Coin: ticker.Symbol, Price: fprice, STExchange: "Binance"}

	return coin, nil
}
