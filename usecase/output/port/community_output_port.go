package outputport

import (
	"service-back/domain/model"
	outputdata "service-back/usecase/output/data"
)

type CommunityOutputPort interface {
	Create(*model.Community) *outputdata.Community
}
