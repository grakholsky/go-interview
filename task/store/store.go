package store

import "context"

type Store[T any] interface {
	FindAll(context.Context) ([]T, error)
}

type Account struct {
	ID         uint64  `json:"id"`
	Asset      string  `json:"asset"`
	Balance    float64 `json:"balance"`
	BalanceUSD float64 `json:"balance_usd"`
}
