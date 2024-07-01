package main

import (
	"fmt"
	"github.com/Ravgus/CryptoPortfolioTracker/internal"
	"sync"
)

func main() {
	internal.LoadEnv()

	var portfolio = internal.GetPortfolioFromJson()
	var currentPrice float64 = 0

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < len(portfolio.Coins); i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			coin := portfolio.Coins[i]
			price := internal.GetCoinPrice(coin.Name)

			mu.Lock()
			currentPrice += price * coin.Count
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println("Total Amount: " + internal.FloatToString(currentPrice))

	history := internal.GetHistory(10)

	exitChan := make(chan struct{})

	for i := 0; i < len(history); i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			select {
			case <-exitChan:
				// exit
				return
			default:
				// continue
			}

			totalPrice := history[i].TotalPrice
			percentageDifference := internal.PercentageDifference(totalPrice, currentPrice)

			if percentageDifference > 25 {
				close(exitChan) // send kill signal

				internal.SendEmail(internal.CreateEmailBody(percentageDifference, history[i].Date))
			}
		}(i)
	}

	wg.Wait()

	internal.AppendHistory(currentPrice, internal.GenerateDate())
}
