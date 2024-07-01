package internal

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func StringToInt(data string) int {
	result, err := strconv.Atoi(data)
	if err != nil {
		fmt.Println("Can't parse string to int!")
		os.Exit(1)
	}

	return result
}

func FloatToString(data float64) string {
	return strconv.FormatFloat(data, 'f', -1, 64)
}

func PercentageDifference(oldValue, newValue float64) float64 {
	if oldValue == 0 {
		return 0
	}

	difference := newValue - oldValue
	percentageDifference := (difference / oldValue) * 100

	return math.Abs(percentageDifference)
}
