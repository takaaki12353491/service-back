package repository

import (
	"service-back/domain/model"
)

type UserRepository interface {
	FindByID(uint) (*model.User, error)
	FindByName(string) (*model.User, error)
	Store(*model.User) error
	Update(*model.User) error
}
