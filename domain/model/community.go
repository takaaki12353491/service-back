package model

import (
	"service-back/validator"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Community struct {
	Model
	Owner   User   `validate:"required"`
	Name    string `validate:"required"`
	Members []User
}

func NewCommunity(owner *User, name string) (*Community, error) {
	id := uuid.New().String()
	community := &Community{
		Model: Model{ID: id},
		Owner: *owner,
		Name:  name,
	}
	err := validator.Validate(community)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return community, nil
}
