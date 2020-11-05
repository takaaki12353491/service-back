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

func (p *CommunityPresenter) Index(cummunities []model.Community) []outputdata.Community {
	oCommunities := []outputdata.Community{}
	for _, community := range cummunities {
		oCommunity := p.convert(&community)
		oCommunities = append(oCommunities, *oCommunity)
	}
	return oCommunities
}

func (p *CommunityPresenter) Show(community *model.Community) *outputdata.Community {
	return p.convert(community)
}

func (p *CommunityPresenter) Edit(community *model.Community) *outputdata.Community {
	return p.convert(community)
}

func (p *CommunityPresenter) convert(community *model.Community) *outputdata.Community {
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
