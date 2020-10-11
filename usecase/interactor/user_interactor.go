package interactor

import (
	"service-back/domain/model"
	"service-back/errs"
	inputdata "service-back/usecase/input/data"
	outputdata "service-back/usecase/output/data"
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

func (it *UserInteractor) Login(iLogin *inputdata.Login) (*outputdata.Login, error) {
	user, err := it.userRepository.FindByName(iLogin.Name)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !user.IsValidPassword(iLogin.Password) {
		errMsg := "The password is invalid"
		log.Error(errMsg)
		return nil, errs.Forbidden.New(errMsg)
	}
	return it.outputport.Login(user), nil
}
