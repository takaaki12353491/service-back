package interactor

import (
	"service-back/domain/model"
	inputdata "service-back/usecase/input/data"
	outputdata "service-back/usecase/output/data"
	outputport "service-back/usecase/output/port"
	"service-back/usecase/repository"

	log "github.com/sirupsen/logrus"
)

type CommunityInteractor struct {
	outputport          outputport.CommunityOutputPort
	communityRepository repository.CommunityRepository
	userRepository      repository.UserRepository
}

func NewCommunityInteractor() *CommunityInteractor {
	return &CommunityInteractor{}
}

func (it *CommunityInteractor) Create(iCommunity *inputdata.Community) (*outputdata.Community, error) {
	user, err := it.userRepository.FindByID(iCommunity.UserID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	community, err := model.NewCommunity(user, iCommunity.Name, iCommunity.Description)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = it.communityRepository.Store(community)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return it.outputport.Create(community), nil
}
