package controllers

import "github.com/ThiagoRodriguesdeSantana/desafio_conductor/go-server-server/go/repository"

type Controller struct {
	Db repository.IRepositoryInterface
}

//NewController instance controller
func NewController(db repository.IRepositoryInterface) *Controller {
	return &Controller{
		Db: db,
	}
}
