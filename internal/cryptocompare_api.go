package internal

import (
	"encoding/json"
	"github.com/Ravgus/CryptoPortfolioTracker/internal/structs"
	"io"
	"log"
	"net/http"
)

func GetCoinPrice(name string) float64 {
	resp, err := http.Get("https://min-api.cryptocompare.com/data/price?fsym=" + name + "&tsyms=USD")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(err)
	}

	responseData, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObject structs.CryptoCompareCoin

	if err := json.Unmarshal(responseData, &responseObject); err != nil {
		log.Fatal(err)
	}

	return responseObject.Price
}
