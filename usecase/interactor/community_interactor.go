package interactor

import (
	"service-back/domain/model"
	"service-back/usecase/inputdata"
	"service-back/usecase/outputdata"
	"service-back/usecase/outputport"
	"service-back/usecase/repository"

	log "github.com/sirupsen/logrus"
)

type CommunityInteractor struct {
	outputport          outputport.CommunityOutputPort
	communityRepository repository.CommunityRepository
	userRepository      repository.UserRepository
}

func NewCommunityInteractor(
	outputport outputport.CommunityOutputPort,
	communityRepository repository.CommunityRepository,
	userRepository repository.UserRepository,
) *CommunityInteractor {
	return &CommunityInteractor{
		outputport:          outputport,
		communityRepository: communityRepository,
		userRepository:      userRepository,
	}
}

func (it *CommunityInteractor) Index() ([]outputdata.Community, error) {
	communities, err := it.communityRepository.FindAll()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return it.outputport.Index(communities), nil
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
