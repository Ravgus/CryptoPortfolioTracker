package main

import (
	"fmt"
	"github.com/Ravgus/CryptoPortfolioTracker/internal"
)

const (
	HistoryCheckIterations = 10
)

func main() {
	internal.LoadEnv()

	var portfolio = internal.GetPortfolioFromJson()

	currentPrice := internal.GetPortfolioPrice(portfolio)

	fmt.Println("Total Amount: " + internal.FloatToString(currentPrice))

	history := internal.GetHistory(HistoryCheckIterations)

	internal.CheckPortfolioPriceChange(currentPrice, history)

	internal.AppendHistory(currentPrice, internal.GenerateDate())
}
