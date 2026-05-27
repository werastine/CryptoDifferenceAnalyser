// Package api contains handlers
package api

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/werastine/CryptoDifferenceAnalyser/internal/analysis"
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

	storage := h.collectPrices(respBody)

	spreadData := analysis.Spread(storage)

	msg := fmt.Sprintf("Buy price(min): %f, %s\nSell price(max): %f, %s\nMaximum spread is %f",
		spreadData.BuyPrice,
		spreadData.BuyExchange,
		spreadData.SellPrice,
		spreadData.SellEcchange,
		spreadData.SellPrice-spreadData.BuyPrice,
	)
	if err := h.response(msg, w); err != nil {
		log.Println("[ERROR] writing response", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// response sends a header and a byte response
func (h *Handler) response(msg string, w http.ResponseWriter) error {
	w.WriteHeader(http.StatusAccepted)
	_, err := w.Write([]byte(msg))
	return err
}

// readBodyToString reads body of request and transfers bytes into string
func readBodyToString(r *http.Request) (string, error) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		return "", fmt.Errorf("search handler: %w", err)
	}
	res := string(httpRequestBody)
	return res, nil
}
