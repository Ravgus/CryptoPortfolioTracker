package interfaces

import "github.com/Ravgus/CryptoPortfolioTracker/internal/structs"

type Portfolio interface {
	GetCoins() []structs.JsonCoin
}
