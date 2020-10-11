package outputport

import (
	"service-back/domain/model"
	outputdata "service-back/usecase/output/data"
)

type UserOutputPort interface {
	Login(*model.User) *outputdata.Login
}
