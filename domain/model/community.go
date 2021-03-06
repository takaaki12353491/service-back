package model

import (
	"mime/multipart"
	"path/filepath"
	"service-back/errs"
	"service-back/validator"
	"strings"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Community struct {
	Model
	OwnerID     string
	Owner       User   `validate:"required"`
	Name        string `validate:"required"`
	Description string
	IconName    string
	Icon        *multipart.FileHeader
	HeaderName  string
	Header      *multipart.FileHeader
	Members     []User
	Invitees    []User
	Applicants  []User
}

func NewCommunity(owner *User, name, description string, icon, header *multipart.FileHeader) (*Community, error) {
	id := uuid.New().String()
	iconExt := filepath.Ext(icon.Filename)
	iconName := strings.Join([]string{"icon_", id, iconExt}, "")
	headerExt := filepath.Ext(header.Filename)
	headerName := strings.Join([]string{"header_", id, headerExt}, "")
	community := &Community{
		Model:       Model{ID: id},
		OwnerID:     owner.ID,
		Owner:       *owner,
		Name:        name,
		Description: description,
		IconName:    iconName,
		Icon:        icon,
		HeaderName:  headerName,
		Header:      header,
	}
	err := validator.Validate(community)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return community, nil
}

func (community *Community) Update(name, description string, icon, header *multipart.FileHeader) error {
	community.Name = name
	community.Description = description
	community.Icon = icon
	community.Header = header
	err := validator.Validate(community)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
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
