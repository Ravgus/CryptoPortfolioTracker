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
		log.Fatal(err)
	}

	fmt.Println("Successfully Opened portfolio.json")

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var portfolio structs.Portfolio

	err = json.Unmarshal(byteValue, &portfolio)

	if err != nil {
		log.Fatal(err)
	}

	return portfolio
}
