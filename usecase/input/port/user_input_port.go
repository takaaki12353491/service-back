package inputport

import inputdata "service-back/usecase/input/data"

type UserInputPort interface {
	SignUp(*inputdata.User) error
}
