package controllers

import "github.com/ThiagoRodriguesdeSantana/desafio_conductor/go/repository"

//Controller to api
type Controller struct {
	Db repository.IRepositoryInterface
}

//NewController instance controller
func NewController(db repository.IRepositoryInterface) *Controller {
	return &Controller{
		Db: db,
	}
}
