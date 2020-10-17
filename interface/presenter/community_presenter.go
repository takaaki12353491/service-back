package presenter

import (
	"service-back/domain/model"
	"service-back/usecase/outputdata"
)

type CommunityPresenter struct {
}

func NewCommunityPresenter() *CommunityPresenter {
	return &CommunityPresenter{}
}

func (p *CommunityPresenter) Create(community *model.Community) *outputdata.Community {
	oOwner := &outputdata.User{
		ID:    community.Owner.ID,
		Name:  community.Owner.Name,
		Email: community.Owner.Email,
	}
	return &outputdata.Community{
		ID:          community.ID,
		Owner:       oOwner,
		Name:        community.Name,
		Description: community.Description,
	}
}
