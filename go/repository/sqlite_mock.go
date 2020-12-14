package repository

import "github.com/ThiagoRodriguesdeSantana/desafio_conductor/go/model"

//SqliteDbMock mock to db
type SqliteDbMock struct {
	OnFindAllAccounts            func() ([]model.Account, error)
	OnFindAccountByID            func(id string) (*model.Account, error)
	OnFindTransactionByAccountID func(id string) ([]model.Transaction, error)
}

//FindAllAccounts mock function
func (m *SqliteDbMock) FindAllAccounts() ([]model.Account, error) {
	return m.OnFindAllAccounts()
}

//FindAccountByID mock function
func (m *SqliteDbMock) FindAccountByID(id string) (*model.Account, error) {
	return m.OnFindAccountByID(id)
}

//FindTransactionByAccountID mock function
func (m *SqliteDbMock) FindTransactionByAccountID(id string) ([]model.Transaction, error) {
	return m.OnFindTransactionByAccountID(id)
}
