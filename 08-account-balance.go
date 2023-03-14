package main

import "fmt"

type Account struct {
	Balance int
}

func AccountBalance() {
	accounts := []Account{
		{Balance: 1000},
		{Balance: 2000},
	}

	for _, account := range accounts {
		account.Balance += 1000
	}

	fmt.Println(accounts)
}

/*
 * Again, for-range loop uses a single variable that is populated with different value for each iteration. It's not a pointer
 * to the original value, it's a copy. This means that we change the Balance for the copy and then immediately discard it.
 */
