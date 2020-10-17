package inputport

import (
	"service-back/usecase/inputdata"
	"service-back/usecase/outputdata"
)

type UserInputPort interface {
	Signup(*inputdata.User) error
	Login(*inputdata.Login) (*outputdata.Login, error)
}
