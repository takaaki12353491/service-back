package repository

import "service-back/domain/model"

type CommunityRepository interface {
	Store(*model.Community) error
}
