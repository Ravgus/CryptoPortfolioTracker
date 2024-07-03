package internal

import (
	"encoding/json"
	"fmt"
	"github.com/Ravgus/CryptoPortfolioTracker/internal/structs"
	"io"
	"log"
	"os"
)

func GetPortfolioFromJson() structs.Portfolio {
	jsonFile, err := os.Open("portfolio.json")

	if err != nil {
		fmt.Println("portfolio.json file not found!")

		os.Exit(2)
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var portfolio structs.Portfolio

	if err := json.Unmarshal(byteValue, &portfolio); err != nil {
		log.Fatal(err)
	}

	return portfolio
}
