package interactor

import (
	"service-back/domain/model"
	"service-back/errs"
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

func (it *CommunityInteractor) Show(id string) (*outputdata.Community, error) {
	community, err := it.communityRepository.FindByID(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return it.outputport.Show(community), nil
}

func (it *CommunityInteractor) Create(iCommunity *inputdata.Community) error {
	user, err := it.userRepository.FindByID(iCommunity.UserID)
	if err != nil {
		log.Error(err)
		return err
	}
	community, err := model.NewCommunity(user, iCommunity.Name, iCommunity.Description)
	if err != nil {
		log.Error(err)
		return err
	}
	err = it.communityRepository.Store(community)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (it *CommunityInteractor) Edit(id string, userID string) (*outputdata.Community, error) {
	community, err := it.communityRepository.FindByID(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if community.OwnerID != userID {
		errMsg := "The user can't get the community"
		log.Error(errMsg)
		return nil, errs.Forbidden.New(errMsg)
	}
	return it.outputport.Edit(community), nil
}

func (it *CommunityInteractor) Update(iUpdateCommunity *inputdata.UpdateCommunity) error {
	community, err := it.communityRepository.FindByID(iUpdateCommunity.ID)
	if err != nil {
		log.Error(err)
		return err
	}
	user, err := it.userRepository.FindByID(iUpdateCommunity.UserID)
	if err != nil {
		log.Error(err)
		return err
	}
	if !community.IsOwner(user) {
		return errs.Forbidden.New("The user can't update the community")
	}
	community.Update(iUpdateCommunity.Name, iUpdateCommunity.Description)
	return nil
}
