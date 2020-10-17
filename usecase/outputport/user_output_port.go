package outputport

import (
	"service-back/domain/model"
	"service-back/usecase/outputdata"
)

type UserOutputPort interface {
	Login(*model.User) *outputdata.Login
}
