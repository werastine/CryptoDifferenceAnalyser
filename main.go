// Package main
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/werastine/CryptoDifferenceAnalyser/service"
)

func main() {
	sharedClient := &http.Client{Timeout: 5 * time.Second}
	srv := service.NewProviders()

	BN := srv.Binance()
	HL := srv.HyperLiquid()
	BB := srv.Bybit()

	CoinBB, err := BB.GetPrice(sharedClient, "TON")
	if err != nil {
		log.Printf("[ERROR] Bybit: %v", err)
	}

	CoinHL, err := HL.GetPrice(sharedClient, "link")
	if err != nil {
		log.Printf("[ERROR] HyperLiquid: %v", err)
	}

	CoinBN, err := BN.GetPrice(sharedClient, "link")
	if err != nil {
		log.Printf("[ERROR] Binance: %v", err)
	}

	fmt.Println(CoinBB.STExchange, CoinBB.Symbol, ":", CoinBB.Price)
	fmt.Println(CoinHL.STExchange, CoinHL.Symbol, ":", CoinHL.Price)
	fmt.Println(CoinBN.STExchange, CoinBN.Symbol, ":", CoinBN.Price)
}
