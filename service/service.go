// Package service - container for exchanges
package service

import (
	"github.com/werastine/CryptoDifferenceAnalyser/internal/exchange"
	"github.com/werastine/CryptoDifferenceAnalyser/internal/market"
)

// Providers - container of exchanges
type Providers struct {
	bybitP       market.PriceProvider
	binanceP     market.PriceProvider
	hyperliquidP market.PriceProvider
}

// NewProviders - constructor for providers
func NewProviders() *Providers {
	return &Providers{
		binanceP:     exchange.ProviderBinance{},
		hyperliquidP: exchange.ProviderHyperLiquid{},
		bybitP:       exchange.ProviderByBit{},
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
