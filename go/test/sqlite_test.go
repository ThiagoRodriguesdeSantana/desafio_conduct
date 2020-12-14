package test

import (
	// "os"
	"os"
	"testing"

	"github.com/ThiagoRodriguesdeSantana/desafio_conductor/go/model"
	"github.com/ThiagoRodriguesdeSantana/desafio_conductor/go/repository"
)

var (
	mockDb *repository.SqliteDbMock
)

func TestRepository(t *testing.T) {

	pathDbTest := "db_test.db"

	mockDb = &repository.SqliteDbMock{}

	t.Run("test Init Data Base - Data Base empty", func(t *testing.T) {

		os.Remove(pathDbTest)
		os.Create(pathDbTest)

		mockDb.OnFindAllAccounts = func() ([]model.Account, error) {
			realDb := repository.InitDB(pathDbTest)
			defer realDb.CloseDb()

			return realDb.FindAllAccounts()
		}

		resp, err := mockDb.FindAllAccounts()

		if err == nil {
			t.Errorf("error")
		}

		if len(resp) > 0 {
			t.Fail()
		}

		os.Remove(pathDbTest)

	})

	t.Run("test FindAllAccounts - Accounts not found", func(t *testing.T) {

		mockDb.OnFindAllAccounts = func() ([]model.Account, error) {
			return nil, nil
		}

		resp, err := mockDb.FindAllAccounts()

		if err != nil {
			t.Errorf("error")
		}

		if resp != nil {
			t.Fail()
		}

	})

	t.Run("test FindAllAccounts - Accounts funded", func(t *testing.T) {

		mockDb.OnFindAllAccounts = func() ([]model.Account, error) {

			realDb := repository.InitDB(pathDbTest)
			defer realDb.CloseDb()

			return realDb.FindAllAccounts()

		}

		resp, err := mockDb.FindAllAccounts()

		if err != nil {
			t.Errorf("error")
		}

		if resp == nil {
			t.Fail()
		}

	})

	t.Run("test OnFindAccountByID - Account not found", func(t *testing.T) {

		mockDb.OnFindAccountByID = func(id string) (*model.Account, error) {
			return nil, nil
		}

		resp, err := mockDb.FindAccountByID("")

		if err != nil {
			t.Errorf("error")
		}

		if resp != nil {
			t.Fail()
		}

	})

	t.Run("test OnFindAccountByID - Account funded", func(t *testing.T) {

		mockDb.OnFindAccountByID = func(id string) (*model.Account, error) {

			realDb := repository.InitDB(pathDbTest)
			defer realDb.CloseDb()

			accounts, _ := realDb.FindAllAccounts()

			account, _ := realDb.FindAccountByID(accounts[0].Id)

			return account, nil
		}

		resp, err := mockDb.FindAccountByID("")

		if err != nil {
			t.Errorf("error")
		}

		if resp == nil {
			t.Fail()
		}

	})

	t.Run("test FindTransactionByAccountID - transaction not found", func(t *testing.T) {

		mockDb.OnFindTransactionByAccountID = func(id string) ([]model.Transaction, error) {
			return nil, nil
		}

		resp, err := mockDb.FindTransactionByAccountID("")

		if err != nil {
			t.Errorf("error")
		}

		if resp != nil {
			t.Fail()
		}

	})

	t.Run("test FindTransactionByAccountID - transaction funded", func(t *testing.T) {

		mockDb.OnFindTransactionByAccountID = func(id string) ([]model.Transaction, error) {

			realDb := repository.InitDB(pathDbTest)
			defer realDb.CloseDb()

			accounts, _ := realDb.FindAllAccounts()

			transaction, _ := realDb.FindTransactionByAccountID(accounts[0].Id)

			return transaction, nil
		}

		resp, err := mockDb.FindAccountByID("")

		if err != nil {
			t.Errorf("error")
		}

		if resp == nil {
			t.Fail()
		}

	})

	t.Run("Remove DB", func(t *testing.T) {
		os.Remove(pathDbTest)
	})

}
