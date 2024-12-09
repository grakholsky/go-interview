package main

import (
	"context"
	"log"

	"task/inmem"
	"task/rates"
	"task/web"
)

func main() {
	handler := &web.AccountsHandler{
		AccountsStore: &rates.AccountsStore{
			AccountsStore: &inmem.AccountsStore{},
		},
	}

	err := handler.ServeHTTP(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
