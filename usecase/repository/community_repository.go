//go:generate mockgen -source=$GOFILE -destination=../mock_repository/mock_$GOFILE
package repository

import "service-back/domain/model"

type CommunityRepository interface {
	FindAll() ([]model.Community, error)
	FindByID(string) (*model.Community, error)
	Store(*model.Community) error
}
