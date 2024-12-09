package rates

import (
	"context"

	"task/store"
)

var prices = map[string]float64{
	"Bitcoin": 62544.30,
}

type AccountsStore struct {
	AccountsStore store.Store[store.Account]
}

func (s *AccountsStore) FindAll(ctx context.Context) ([]store.Account, error) {
	accounts, err := s.AccountsStore.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	for i, a := range accounts {
		if price, ok := prices[a.Asset]; ok {
			accounts[i].BalanceUSD = a.Balance * price
		}
	}

	return accounts, nil
}
