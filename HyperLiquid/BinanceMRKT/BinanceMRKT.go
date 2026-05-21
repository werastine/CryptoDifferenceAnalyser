package binancemrkt

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Url struct {
	url string
}

// Func to add Coin name and USDT into link
func Insert(url string, idx int, insertion string) string {
	res := url[:idx] + insertion + "USDT" + url[idx:]
	return res
}

func NewUrl() Url {
	return Url{url: "https://api.binance.com/api/v3/ticker/price?symbol="}
}

type BinanceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func GetPriceBinance(Coin string) {

	client := &http.Client{Timeout: 10 * time.Second}

	urlData := NewUrl()
	urlData.url = Insert(urlData.url, 51, Coin)

	response, err := client.Get(urlData.url)
	if err != nil {
		fmt.Println("Fail got no answer", err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(response.Body)
		fmt.Println("Fail request", response.StatusCode, string(bodyBytes))
		return
	}

	bodyBytes, err := io.ReadAll(response.Body)
}
