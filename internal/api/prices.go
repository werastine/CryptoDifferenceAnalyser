package api

import (
	"log"
	"sync"

	"github.com/werastine/CryptoDifferenceAnalyser/internal/market"
)

func (h *Handler) collectPrices(symbol string) map[market.CoinToReturn]struct{} {
	storage := make(map[market.CoinToReturn]struct{})
	wg := &sync.WaitGroup{}

	transferPoint := make(chan market.CoinToReturn)

	wg.Add(1)
	go h.fetchBinance(wg, transferPoint, symbol)

	wg.Add(1)
	go h.fetchByBit(wg, transferPoint, symbol)

	wg.Add(1)
	go h.fetchHyperLiquid(wg, transferPoint, symbol)

	go func() {
		wg.Wait()
		close(transferPoint)
	}()

	for val := range transferPoint {
		if val.Err != nil {
			continue
		}
		storage[val] = struct{}{}

	}

	return storage
}

func (h *Handler) fetchBinance(
	wg *sync.WaitGroup,
	ch chan<- market.CoinToReturn,
	symbol string,
) {
	defer wg.Done()

	BN := h.providers.Binance()
	coinData, err := BN.GetPrice(symbol)
	if err != nil {
		log.Printf("[ERROR] Binance: %v", err)
		coinData.Err = err
		return
	}
	ch <- coinData
}

func (h *Handler) fetchByBit(
	wg *sync.WaitGroup,
	ch chan<- market.CoinToReturn,
	symbol string,
) {
	defer wg.Done()

	BB := h.providers.Bybit()
	coinData, err := BB.GetPrice(symbol)
	if err != nil {
		log.Printf("[ERROR] ByBit: %v", err)
		coinData.Err = err
		return
	}

	ch <- coinData
}

func (h *Handler) fetchHyperLiquid(
	wg *sync.WaitGroup,
	ch chan<- market.CoinToReturn,
	symbol string,
) {
	defer wg.Done()

	HL := h.providers.HyperLiquid()
	coinData, err := HL.GetPrice(symbol)
	if err != nil {
		log.Printf("[ERROR] HyperLiquid: %v", err)
		coinData.Err = err
		return
	}

	ch <- coinData
}
