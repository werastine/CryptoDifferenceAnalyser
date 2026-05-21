package main

import (
	"fmt"

	binance "github.com/werastine/CryptoDifferenceAnalyser/BinanceMRKT"
	hype "github.com/werastine/CryptoDifferenceAnalyser/HyperLiquid"
)

func main() {

	CoinHL, err := hype.GetPriceHyperLiquid("LINK")
	if err != nil {
		fmt.Println("Found error ir GetPriceHyperLiquid func", err)
	}

	CoinBN, err := binance.GetPriceBinance("LINK")
	if err != nil {
		fmt.Println("Found error ir GetPriceHyperLiquid func", err)
	}

	fmt.Println(CoinHL.Coin, ":", CoinHL.Price)
	fmt.Println(CoinBN.Coin, ":", CoinBN.Price)
}
