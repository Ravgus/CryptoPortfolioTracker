package structs

type Portfolio struct {
	Coins []JsonCoin `json:"coins"`
}

func (portfolio Portfolio) GetCoins() []JsonCoin {
	return portfolio.Coins
}
