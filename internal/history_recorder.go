package internal

import (
	"encoding/json"
	"fmt"
	"github.com/Ravgus/CryptoPortfolioTracker/internal/structs"
	"io"
	"log"
	"os"
)

func GetHistory(args ...int) []structs.HistoryItem {
	var size int

	if len(args) != 0 {
		size = args[0]
	}

	jsonFile, err := os.Open("history.json")

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("history.json does not exist!")

			return nil
		} else {
			log.Fatal(err)
		}
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var history []structs.HistoryItem

	if err := json.Unmarshal(byteValue, &history); err != nil {
		log.Fatal(err)
	}

	if &size != nil && len(history) > size {
		return history[len(history)-size:]
	}

	return history
}

func AppendHistory(totalPrice float64, date string) {
	fileName := "history.json"

	file, err := os.ReadFile(fileName)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("creating history.json file...")

			file = []byte("[]")
		} else {
			log.Fatal(err)
		}
	}

	var history []structs.HistoryItem

	if err := json.Unmarshal(file, &history); err != nil {
		log.Fatal(err)
	}

	newItem := structs.HistoryItem{TotalPrice: totalPrice, Date: date}

	history = append(history, newItem)

	updatedData, err := json.MarshalIndent(history, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(fileName, updatedData, 0644); err != nil {
		log.Fatal(err)
	}
}
