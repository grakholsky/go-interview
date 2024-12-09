package web

import (
	"context"
	"encoding/json"
	"log"

	"task/store"
)

type AccountsHandler struct {
	AccountsStore store.Store[store.Account]
}

func (h *AccountsHandler) ServeHTTP(ctx context.Context) error {
	accounts, err := h.AccountsStore.FindAll(ctx)
	if err != nil {
		return err
	}

	respBody, err := json.Marshal(accounts)
	if err != nil {
		return err
	}

	// Writing response body to client
	log.Println(string(respBody))
	return nil
}
