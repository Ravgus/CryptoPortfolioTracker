package internal

import (
	"log"
	"math"
	"strconv"
)

func StringToInt(data string) int {
	result, err := strconv.Atoi(data)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

func StringToFloat(data string) float64 {
	result, err := strconv.ParseFloat(data, 64)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

func FloatToString(data float64) string {
	return strconv.FormatFloat(data, 'f', 2, 64)
}

func PercentageDifference(oldValue, newValue float64) float64 {
	if oldValue == 0 {
		return 0
	}

	difference := newValue - oldValue
	percentageDifference := (difference / oldValue) * 100

	return math.Abs(percentageDifference)
}
