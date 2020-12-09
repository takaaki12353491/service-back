//go:generate mockgen -source=$GOFILE -destination=../mock_repository/mock_$GOFILE
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
