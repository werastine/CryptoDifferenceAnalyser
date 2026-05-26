// Package service - container for exchanges
package service

import (
	"net/http"
	"sync"

	"github.com/werastine/CryptoDifferenceAnalyser/internal/exchange"
	"github.com/werastine/CryptoDifferenceAnalyser/internal/market"
)

// Providers - container of exchanges
type Providers struct {
	waitGroup    *sync.WaitGroup
	bybitP       market.PriceProvider
	binanceP     market.PriceProvider
	hyperliquidP market.PriceProvider
}

// NewProviders - constructor for providers
func NewProviders(cl *http.Client, wg *sync.WaitGroup) *Providers {
	return &Providers{
		waitGroup:    wg,
		binanceP:     exchange.NewProviderBinance(cl),
		hyperliquidP: exchange.NewProviderHyperLiquid(cl),
		bybitP:       exchange.NewProviderByBit(cl),
	}
}

// Binance returns binance.ProviderBinance{}
func (p *Providers) Binance() market.PriceProvider {
	return p.binanceP
}

// Bybit returns bybit.ProviderByBit{}
func (p *Providers) Bybit() market.PriceProvider {
	return p.bybitP
}

// HyperLiquid returns hyperliquid.ProviderHyperLiquid{}
func (p *Providers) HyperLiquid() market.PriceProvider {
	return p.hyperliquidP
}

// GetWaiyGroup returns wg
func (p *Providers) GetWaiyGroup() *sync.WaitGroup {
	return p.waitGroup
}
