package internal

import (
	"fmt"
	"github.com/Ravgus/CryptoPortfolioTracker/internal/structs"
	"os"
	"sync"
	"time"
)

const (
	SignificantChangePercent = 25
)

func CheckPortfolioPriceChange(currentPrice float64, history []structs.HistoryItem) {
	if checkByPrice(currentPrice) {
		return
	}

	checkByPercent(currentPrice, history)
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

func checkByPercent(currentPrice float64, history []structs.HistoryItem) {
	var wg sync.WaitGroup

	exitChan := make(chan struct{})

	var changePercent float64
	if len(os.Getenv("NOTIFICATION_CHANGE_PERCENT")) == 0 {
		changePercent = SignificantChangePercent
	} else {
		changePercent = StringToFloat(os.Getenv("NOTIFICATION_CHANGE_PERCENT"))
	}

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

			if percentageDifference > changePercent {
				close(exitChan) // send kill signal

				SendEmail(CreatePercentEmailBody(percentageDifference, history[i].Date))

				fmt.Println("Notification was sent!")
			}
		}(i)
	}

	wg.Wait()
}

func checkByPrice(currentPrice float64) bool {
	if len(os.Getenv("NOTIFICATION_CHANGE_PRICE")) != 0 {
		trackedPrice := StringToFloat(os.Getenv("NOTIFICATION_CHANGE_PRICE"))

		if currentPrice >= trackedPrice {
			SendEmail(CreatePriceEmailBody(trackedPrice))

			fmt.Println("Notification was sent!")

			return true
		}
	}

	return false
}
