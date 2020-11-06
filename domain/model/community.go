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
	Invitees    []User
	Applicants  []User
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
	if !community.IsMember(owner) {
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

func (community *Community) Invite(user, invitee *User) error {
	if !community.IsOwner(user) {
		return errs.Forbidden.New("The user can't invite an user")
	}
	if !community.IsParticipant(invitee) {
		return errs.Conflict.New("The invitee is invalid")
	}
	community.Invitees = append(community.Invitees, *invitee)
	return nil
}

func (community *Community) Apply(applicant *User) error {
	if !community.IsParticipant(applicant) {
		return errs.Conflict.New("The applicant is invalid")
	}
	community.Applicants = append(community.Applicants, *applicant)
	return nil
}

func (community *Community) IsOwner(user *User) bool {
	return community.OwnerID == user.ID
}

func (community *Community) IsMember(user *User) bool {
	for _, member := range community.Members {
		if member.ID == user.ID {
			return true
		}
	}
	return false
}

func (community *Community) IsInvitee(user *User) bool {
	for _, invitee := range community.Invitees {
		if invitee.ID == user.ID {
			return true
		}
	}
	return false
}

func (community *Community) IsApplicant(userID string) bool {
	for _, applicant := range community.Applicants {
		if applicant.ID == userID {
			return true
		}
	}
	return false
}

func (community *Community) IsParticipant(user *User) bool {
	return community.IsOwner(user) || community.IsMember(user) || community.IsInvitee(user) || community.IsParticipant(user)
}
