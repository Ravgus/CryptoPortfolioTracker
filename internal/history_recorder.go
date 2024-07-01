package internal

import (
	"encoding/csv"
	"github.com/Ravgus/CryptoPortfolioTracker/internal/structs"
	"github.com/gocarina/gocsv"
	"log"
	"os"
)

func GetHistory(args ...int) []structs.HistoryItem {
	var size int

	if len(args) != 0 {
		size = args[0]
	}

	file, err := os.Open("history.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var history []structs.HistoryItem

	if err := gocsv.UnmarshalFile(file, &history); err != nil {
		log.Fatal(err)
	}

	if &size != nil && len(history) > size {
		return history[len(history)-size:]
	}

	return history
}

func AppendHistory(totalPrice float64, date string) {
	file, err := os.OpenFile("history.csv", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	row := []string{FloatToString(totalPrice), date}
	err = writer.Write(row)

	if err != nil {
		log.Fatal(err)
	}
}
