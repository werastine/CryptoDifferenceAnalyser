// Package market contains the interface of all markets
package market

// CoinToReturn - standatr for interface
type CoinToReturn struct {
	Symbol     string
	Price      float64
	STExchange string
	Err        error
}

// PriceProvider contains GetPrice
type PriceProvider interface {
	GetPrice(coin string) (CoinToReturn, error)
}
