package interactor

import (
	"service-back/domain/model"
	inputdata "service-back/usecase/input/data"
	outputport "service-back/usecase/output/port"
	"service-back/usecase/repository"

	log "github.com/sirupsen/logrus"
)

type UserInteractor struct {
	userRepository repository.UserRepository
	outputport     outputport.UserOutputPort
}

func NewUserInteractor(
	outputport outputport.UserOutputPort,
	userRepository repository.UserRepository,
) *UserInteractor {
	return &UserInteractor{
		userRepository: userRepository,
		outputport:     outputport,
	}
}

func (it *UserInteractor) Signup(iUser *inputdata.User) error {
	user, err := model.NewUser(iUser.Name, iUser.Email, iUser.Password)
	if err != nil {
		log.Error(err)
		return err
	}
	err = it.userRepository.Store(user)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
