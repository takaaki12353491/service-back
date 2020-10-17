package inputport

import (
	"service-back/usecase/inputdata"
	"service-back/usecase/outputdata"
)

type CommunityInputPort interface {
	Create(*inputdata.Community) (*outputdata.Community, error)
}
