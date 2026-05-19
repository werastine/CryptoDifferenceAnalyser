package hyperliquid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// this is requied request for hyperliquid
type InfoRequest struct {
	Type string `json:"type"`
}

// local data for hyperliquid.go
type Data struct {
	url string
}

// constructor for Data
func newData() *Data {
	return &Data{
		url: "https://api.hyperliquid.xyz/info",
	}
}

type Coin struct {
	Coin  string
	Price float64
}

// Coin price getter
func GetPriceHyperLiquid(coin string) (Coin, error) {
	data := newData()

	requestBody := InfoRequest{
		Type: "allMids",
	}

	jsonBytes, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Print("Fail marhsl json", err)
		return Coin{}, err
	}

	response, err := http.Post(data.url, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Print("Fail in request", err)
		return Coin{}, err
	}

	defer response.Body.Close()

	var prices map[string]string
	if err = json.NewDecoder(response.Body).Decode(&prices); err != nil {
		fmt.Println("Fail decode json", err)
		return Coin{}, err
	}

	price, ok := prices[coin]
	if !ok {
		fmt.Println("There is no", coin)
	}
	fprice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Println("Fail convert string price to float64")
	}

	return Coin{Price: fprice, Coin: coin}, nil
}
