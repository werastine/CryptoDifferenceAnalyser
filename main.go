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

	coinBB, err := bybit.GetByBitPrice(&sharedClient, "TON")
	if err != nil {
		log.Printf("[ERROR] Bybit: %v", err)
	}

	CoinHL, err := hyperliquid.GetPriceHyperLiquid(&sharedClient, "TON")
	if err != nil {
		fmt.Println("Found error in GetPriceHyperLiquid func: ", err)
	}

	CoinBN, err := binance.GetPriceBinance(&sharedClient, "TON")
	if err != nil {
		fmt.Println("Found error in GetPriceBinance func:", err)
	}

	fmt.Println(coinBB.STExchange, coinBB.Symbol, ":", CoinHL.Price)
	fmt.Println(CoinHL.STExchange, CoinHL.Coin, ":", CoinHL.Price)
	fmt.Println(CoinBN.STExchange, CoinBN.Coin, ":", CoinBN.Price)

	// fmt.Println("Difference is:", CoinBN.Price-CoinHL.Price)
}
