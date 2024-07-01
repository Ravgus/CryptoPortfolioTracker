package structs

type HistoryItem struct {
	TotalPrice float64 `csv:"total_price"`
	Date       string  `csv:"date"`
}
