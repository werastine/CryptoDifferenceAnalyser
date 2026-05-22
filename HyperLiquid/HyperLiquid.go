// Package hyperliquid contains the logic of request on HyperLiquid exchange
package hyperliquid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// InfoRequest is requied request for hyperliquid
type InfoRequest struct {
	Type string `json:"type"`
}

// Data struct contains url from hyperliquid
type Data struct {
	url string
}

// constructor for Data
func newData() *Data {
	return &Data{
		url: "https://api.hyperliquid.xyz/info",
	}
}

// CoinToReturn contains data about the coin wich user going to recieve
type CoinToReturn struct {
	Coin       string
	Price      float64
	STExchange string
}

// GetPriceHyperLiquid func - global price getter
func GetPriceHyperLiquid(client *http.Client, coin string) (CoinToReturn, error) {

	data := newData()

	requestBody := InfoRequest{
		Type: "allMids",
	}

	jsonBytes, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Print("Fail marhsal json", err)
		return CoinToReturn{}, err
	}

	response, err := client.Post(data.url, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Print("Fail in request", err)
		return CoinToReturn{}, err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			fmt.Println("fail to close the body", err)
		}
	}()

	if response.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(response.Body)
		err := fmt.Errorf("fail, wrong status code %d, error is: %s", response.StatusCode, string(bodyBytes))
		return CoinToReturn{}, err
	}

	var prices map[string]string
	if err = json.NewDecoder(response.Body).Decode(&prices); err != nil {
		fmt.Println("Fail decode json", err)
		return CoinToReturn{}, err
	}

	price, ok := prices[coin]
	if !ok {
		// fmt.Println("wrong request, can't find:", coin)
		return CoinToReturn{}, fmt.Errorf("wrong request, can't find: %s", coin)
	}

	fprice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Println("Fail convert string price to float64")
		return CoinToReturn{}, err
	}

	return CoinToReturn{Price: fprice, Coin: coin, STExchange: "HyperLiquid"}, nil
}
