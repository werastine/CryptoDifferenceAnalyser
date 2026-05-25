// Package main
package main

import (
	"fmt"
	"log"
	"net/http"

	"time"

	handlers "github.com/werastine/CryptoDifferenceAnalyser/Handlers"
	"github.com/werastine/CryptoDifferenceAnalyser/service"
)

func main() {
	sharedClient := &http.Client{Timeout: 5 * time.Second}
	srv := service.NewProviders()

	handlers.RegisterRoutes()

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Println("[ERROR] connecting to a port:", err)
	}

	// wg := sync.WaitGroup{}

	BN := srv.Binance()
	HL := srv.HyperLiquid()
	BB := srv.Bybit()

	CoinBB, err := BB.GetPrice(sharedClient, "Link")
	if err != nil {
		log.Printf("[ERROR] Bybit: %v", err)
	}

	CoinHL, err := HL.GetPrice(sharedClient, "Link")
	if err != nil {
		log.Printf("[ERROR] HyperLiquid: %v", err)
	}

	CoinBN, err := BN.GetPrice(sharedClient, "Link")
	if err != nil {
		log.Printf("[ERROR] Binance: %v", err)
	}

	fmt.Println(CoinBB.STExchange, CoinBB.Symbol, ":", CoinBB.Price)
	fmt.Println(CoinHL.STExchange, CoinHL.Symbol, ":", CoinHL.Price)
	fmt.Println(CoinBN.STExchange, CoinBN.Symbol, ":", CoinBN.Price)

}
