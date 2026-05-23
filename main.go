// Package main
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	binance "github.com/werastine/CryptoDifferenceAnalyser/BinanceMRKT"
	bybit "github.com/werastine/CryptoDifferenceAnalyser/ByBit"
	hyperliquid "github.com/werastine/CryptoDifferenceAnalyser/HyperLiquid"
)

func main() {

	sharedClient := http.Client{Timeout: 5 * time.Second}

	coinBB, err := bybit.GetByBitPrice(&sharedClient, "LINK")
	if err != nil {
		log.Printf("[ERROR] Bybit: %v", err)
	}

	CoinHL, err := hyperliquid.GetPriceHyperLiquid(&sharedClient, "link")
	if err != nil {
		log.Printf("[ERROR] HyperLiquid: %v", err)
	}

	CoinBN, err := binance.GetPriceBinance(&sharedClient, "link")
	if err != nil {
		log.Printf("[ERROR] Binance: %v", err)
	}

	fmt.Println(coinBB.STExchange, coinBB.Symbol, ":", coinBB.Price)
	fmt.Println(CoinHL.STExchange, CoinHL.Coin, ":", CoinHL.Price)
	fmt.Println(CoinBN.STExchange, CoinBN.Coin, ":", CoinBN.Price)
}
