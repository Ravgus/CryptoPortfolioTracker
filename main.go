package main

import (
	"fmt"
	"github.com/Ravgus/CryptoPortfolioTracker/internal"
	"strconv"
	"sync"
)

func main() {
	internal.LoadEnv()

	var portfolio = internal.GetPortfolioFromJson()
	var totalAmount float64 = 0

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < len(portfolio.Coins); i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			coin := portfolio.Coins[i]
			price := internal.GetCoinPrice(coin.Name)

			mu.Lock()
			totalAmount += price * coin.Count
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println("Total Amount: " + strconv.FormatFloat(totalAmount, 'f', -1, 64))
}
