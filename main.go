package main

import (
	// "fmt"

	hype "DifferenceAnalyser/hyperliquid"
	"fmt"
)

func main() {

	btcInfo, err := hype.GetPriceHyperLiquid("TON")
	if err != nil {
		fmt.Println("Found error ir GetPriceHyperLiquid func", err)
	}

	if btcInfo.Price >= 77000 {
		fmt.Printf("%s price is %f and it's more than 77000", btcInfo.Coin, btcInfo.Price)
	} else {
		fmt.Printf("%s price is %f, its less then 77000", btcInfo.Coin, btcInfo.Price)
	}
}
