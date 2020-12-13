package repository

import "github.com/ThiagoRodriguesdeSantana/desafio_conductor/go-server-server/go/model"

//IRepositoryInterface inteface data base
type IRepositoryInterface interface {
	FindAllAccounts() ([]model.Account, error)
	FindAccountByID(id string) (*model.Account, error)
	FindTransactionByAccountID(id string) ([]model.Transaction, error)
}
