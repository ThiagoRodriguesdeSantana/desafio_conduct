package repository

import (
	"fmt"
	"testing"
)

func TestRepository(t *testing.T) {

	db := InitDB()
	t.Run("test query", func(t *testing.T) {

		resp, err := db.FindAllAccounts()

		if err != nil {
			t.Errorf("error")
		}

		fmt.Println(resp[0].Id)

	})

	t.Run("test query", func(t *testing.T) {

		resp, err := db.FindAllAccounts()
		account, err := db.FindAccountByID(resp[0].Id)

		if err != nil {
			t.Errorf("error")
		}

		fmt.Println(account.Id)

	})

	t.Run("test query", func(t *testing.T) {

		resp, err := db.FindAllAccounts()
		transaction, err := db.FindTransactionByAccountID(resp[0].Id)

		if err != nil {
			t.Errorf("error")
		}

		fmt.Println(transaction[0].Id)

	})
}
