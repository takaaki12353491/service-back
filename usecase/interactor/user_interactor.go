package interactor

import (
	inputdata "service-back/usecase/input/data"
	outputport "service-back/usecase/output/port"
	"service-back/usecase/repository"
)

type UserInteractor struct {
	userRepository repository.UserRepository
	outputport     outputport.UserOutputPort
}

func NewUserInteractor(
	userRepository repository.UserRepository,
	outputport outputport.UserOutputPort,
) *UserInteractor {
	return &UserInteractor{
		userRepository: userRepository,
		outputport:     outputport,
	}
}

func (it *UserInteractor) SignUp(user *inputdata.User) {

}
