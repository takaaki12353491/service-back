package outputport

import (
	"service-back/domain/model"
	"service-back/usecase/outputdata"
)

type CommunityOutputPort interface {
	Index([]model.Community) []outputdata.Community
	Show(*model.Community) *outputdata.Community
	Create(*model.Community) *outputdata.Community
}
