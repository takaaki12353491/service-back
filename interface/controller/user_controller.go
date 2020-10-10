package controller

import (
	"net/http"
	"service-back/interface/database"
	"service-back/interface/presenter"
	inputport "service-back/usecase/input/port"
	"service-back/usecase/interactor"
)

type UserController struct {
	inputport inputport.UserInputPort
}

func NewUserController() *UserController {
	return &UserController{
		inputport: interactor.NewUserInteractor(
			presenter.NewUserPresenter(),
			database.NewUserDatabase(),
		),
	}
}

// SignUp ...
// @summary
// @description
// @tags User
// @accept json
// @produce json
// @router /signup [post]
func (ctrl *UserController) SignUp(c Context) error {
	return c.JSON(http.StatusOK, nil)
}
