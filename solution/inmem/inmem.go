package inmem

import (
	"context"

	"task/store"
)

type AccountsStore struct{}

func (s *AccountsStore) FindAll(ctx context.Context) ([]store.Account, error) {
	return Accounts, nil
}
