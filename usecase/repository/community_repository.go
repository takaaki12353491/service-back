package repository

import "service-back/domain/model"

type CommunityRepository interface {
	FindAll() ([]model.Community, error)
	Store(*model.Community) error
}
