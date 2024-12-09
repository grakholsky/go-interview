package inmem

import "task/store"

var Accounts = []store.Account{
	{
		ID:      1,
		Asset:   "Bitcoin",
		Balance: 0.1,
	},
}
