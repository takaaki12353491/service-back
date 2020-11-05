package model

import (
	"mime/multipart"
	"service-back/errs"
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
	LogoURL     string
	Logo        *multipart.FileHeader
	HeaderURL   string
	Header      *multipart.FileHeader
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

func (community *Community) NewProject(owner *User, name, description string) (*Project, error) {
	isMemeber := community.IsMember(owner)
	if !isMemeber {
		return nil, errs.Forbidden.New("The user is not member")
	}
	id := uuid.New().String()
	project := &Project{
		Model:       Model{ID: id},
		CommunityID: community.ID,
		Community:   *community,
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
	return project, nil
}

func (community *Community) IsMember(user *User) bool {
	for _, member := range community.Members {
		if member.ID == user.ID {
			return true
		}
	}
	return false
}
