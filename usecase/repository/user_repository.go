package repository

import (
	"service-back/domain/model"
)

type UserRepository interface {
	FindByNameOrEmail(name, email string) (*model.User, error)
	FindByID(string) (*model.User, error)
	Store(*model.User) error
	Update(*model.User) error
}
