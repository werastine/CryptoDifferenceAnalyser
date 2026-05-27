// Package api contains handlers
package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/werastine/CryptoDifferenceAnalyser/internal/analysis"
	"github.com/werastine/CryptoDifferenceAnalyser/internal/market"
	"github.com/werastine/CryptoDifferenceAnalyser/service"
)

// Handler of client
type Handler struct {
	providers *service.Providers
}

// NewHandler for getting client
func NewHandler(pr *service.Providers) *Handler {
	return &Handler{
		providers: pr,
	}
}

// RegisterRoutes regists all routes in main, to make all handlers usable
func (h *Handler) RegisterRoutes() {
	http.HandleFunc("/search", h.SearchHandler)
}

// SearchHandler handles operation of getting prices from exchanges
func (h *Handler) SearchHandler(w http.ResponseWriter, r *http.Request) {
	respBody, err := readBodyToString(r)
	if err != nil {
		log.Printf("[ERROR %v]", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println("resp body", respBody)

	storage := h.getPrices(respBody)
	for key := range storage {
		fmt.Println(key.STExchange, key.Symbol, key.Price)
	}

	spreadData := analysis.Spread(storage)

	msg := fmt.Sprintf("Buy price(min): %f, %s\nSell price(max): %f, %s\nMaximum spread is %f",
		spreadData.BuyPrice,
		spreadData.BuyExchange,
		spreadData.SellPrice,
		spreadData.SellEcchange,
		spreadData.SellPrice-spreadData.BuyPrice,
	)
	if err := h.sendData(msg, w); err != nil {
		log.Println("[ERROR] writing response", err)
	}
}

func (h *Handler) sendData(msg string, w http.ResponseWriter) error {
	w.WriteHeader(http.StatusAccepted)
	_, err := w.Write([]byte(msg))
	return err
}

func readBodyToString(r *http.Request) (string, error) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		return "", fmt.Errorf("search handler: %w", err)
	}
	res := string(httpRequestBody)
	return res, nil
}

// Add error return for getPrices() query (line 41)
// getPrice returns price in channels to SearchHandler
func (h *Handler) getPrices(symbol string) map[market.CoinToReturn]struct{} {
	storage := make(map[market.CoinToReturn]struct{})

	wg := h.providers.GetWaiyGroup()
	transferPoint := make(chan market.CoinToReturn)

	wg.Add(1)
	go h.getPriceBinance(wg, transferPoint, symbol)

	wg.Add(1)
	go h.getPriceByBit(wg, transferPoint, symbol)

	wg.Add(1)
	go h.getPriceHyperLiquid(wg, transferPoint, symbol)

	go func() {
		wg.Wait()
		close(transferPoint)
	}()

	for val := range transferPoint {
		storage[val] = struct{}{}
	}

	return storage
}

func (h *Handler) getPriceBinance(
	wg *sync.WaitGroup,
	ch chan<- market.CoinToReturn,
	symbol string,
) {
	defer wg.Done()

	BN := h.providers.Binance()
	coinData, err := BN.GetPrice(symbol)
	if err != nil {
		log.Printf("[ERROR] Binance: %v", err)
		return
	}
	ch <- coinData
}

func (h *Handler) getPriceByBit(
	wg *sync.WaitGroup,
	ch chan<- market.CoinToReturn,
	symbol string,
) {
	defer wg.Done()

	BB := h.providers.Bybit()
	coinData, err := BB.GetPrice(symbol)
	if err != nil {
		log.Printf("[ERROR] ByBit: %v", err)
		return
	}

	ch <- coinData
}

func (h *Handler) getPriceHyperLiquid(
	wg *sync.WaitGroup,
	ch chan<- market.CoinToReturn,
	symbol string,
) {
	defer wg.Done()

	HL := h.providers.HyperLiquid()
	coinData, err := HL.GetPrice(symbol)
	if err != nil {
		log.Printf("[ERROR] HyperLiquid: %v", err)
		return
	}

	ch <- coinData
}
