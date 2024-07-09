package internal

import (
	"fmt"
	"github.com/Ravgus/CryptoPortfolioTracker/internal/structs"
	"sync"
	"time"
)

const (
	SignificantChangePercent = 25
)

func CheckPortfolioPriceChange(currentPrice float64, history []structs.HistoryItem) {
	var wg sync.WaitGroup

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
			percentageDifference := PercentageDifference(totalPrice, currentPrice)

			if percentageDifference > SignificantChangePercent {
				close(exitChan) // send kill signal

				SendEmail(CreateEmailBody(percentageDifference, history[i].Date))

				fmt.Println("Notification was sent!")
			}
		}(i)
	}

	wg.Wait()
}

func GetPortfolioPrice(portfolio structs.Portfolio) float64 {
	var currentPrice float64 = 0

	fmt.Println("Calculating portfolio sum...")

	for i := 0; i < len(portfolio.Coins); i++ {
		coin := portfolio.Coins[i]
		price := GetCoinPrice(coin.Name)

		// because of cryptocompare api limits
		time.Sleep(500 * time.Millisecond)

		currentPrice += price * coin.Count
	}

	return currentPrice
}
