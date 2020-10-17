package outputport

import (
	"service-back/domain/model"
	"service-back/usecase/outputdata"
)

type CommunityOutputPort interface {
	Create(*model.Community) *outputdata.Community
}
