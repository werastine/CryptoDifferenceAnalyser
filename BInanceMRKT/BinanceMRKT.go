package binancemrkt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Url struct {
	url string
}

// Func to add Coin name and USDT into link
func insert(url string, idx int, insertion string) string {
	res := url[:idx] + insertion + "USDT" + url[idx:]
	return res
}

func newUrl() Url {
	return Url{url: "https://api.binance.com/api/v3/ticker/price?symbol="}
}

type binanceTicker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type CoinToReturn struct {
	Coin  string
	Price float64
}

func GetPriceBinance(Coin string) (CoinToReturn, error) {

	client := &http.Client{Timeout: 10 * time.Second}

	urlData := newUrl()
	urlData.url = insert(urlData.url, 51, Coin)

	response, err := client.Get(urlData.url)
	if err != nil {
		fmt.Println("Fail got no answer", err)
		return CoinToReturn{}, err
	}

	defer response.Body.Close()

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

	coin := CoinToReturn{Coin: ticker.Symbol, Price: fprice}

	return coin, nil
}
