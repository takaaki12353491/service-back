package model

import (
	"service-back/validator"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Community struct {
	Model
	OwnerID     string
	Owner       User   `validate:"required"`
	Name        string `validate:"required"`
	Description string
	Members     []User
}

func NewCommunity(owner *User, name, description string) (*Community, error) {
	id := uuid.New().String()
	community := &Community{
		Model:       Model{ID: id},
		OwnerID:     owner.ID,
		Owner:       *owner,
		Name:        name,
		Description: description,
	}
	err := validator.Validate(community)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return community, nil
}
