package inputport

import (
	inputdata "service-back/usecase/input/data"
	outputdata "service-back/usecase/output/data"
)

type UserInputPort interface {
	Signup(*inputdata.User) error
	Login(*inputdata.Login) (*outputdata.Login, error)
}
