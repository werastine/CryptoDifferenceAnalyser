// Package market contains the interface of all markets
package market

import "net/http"

// CoinToReturn - standatr for interface
type CoinToReturn struct {
	Symbol     string
	Price      float64
	STExchange string
}

// PriceProvider contains GetPrice
type PriceProvider interface {
	GetPrice(client *http.Client, coin string) (CoinToReturn, error)
}
