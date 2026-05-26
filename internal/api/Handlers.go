// Package api contain handlers
package api

import (
	"fmt"
	"io"
	"log"
	"net/http"

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

	h.getPrices(respBody) // make err check, etc...
}

func readBodyToString(r *http.Request) (string, error) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		return "", fmt.Errorf("search handler: %w", err)
	}
	res := string(httpRequestBody)
	return res, nil
}

// getprice is func
func (h *Handler) getPrices(coin string) {
	// there will be 3 goroutines sending request for 3 stock exchanges
	// h.providers.Binance().GetPrice()
}
