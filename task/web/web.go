package web

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"task/store"
)

type AccountsHandler struct {
	// TODO: Add dependencies
}

func (h *AccountsHandler) ServeHTTP(ctx context.Context) error {
	accounts, err := h.AccountsStore.FindAll(ctx)
	if err != nil {
		return err
	}

	bits, err := json.Marshal(accounts)
	if err != nil {
		return fmt.Errorf("failed to marshal accounts: %w", err)
	}

	// Writing response body to client
	log.Println(string(bits))

	return nil
}
