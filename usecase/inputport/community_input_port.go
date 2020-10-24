package inputport

import (
	"service-back/usecase/inputdata"
	"service-back/usecase/outputdata"
)

type CommunityInputPort interface {
	Index() ([]outputdata.Community, error)
	Show(string) (*outputdata.Community, error)
	Create(*inputdata.Community) (*outputdata.Community, error)
}
