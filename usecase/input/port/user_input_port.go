package inputport

import inputdata "service-back/usecase/input/data"

type UserInputPort interface {
	Signup(*inputdata.User) error
}
