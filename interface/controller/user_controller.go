package controller

import (
	"net/http"
	"service-back/interface/database"
	"service-back/interface/presenter"
	inputdata "service-back/usecase/input/data"
	inputport "service-back/usecase/input/port"
	"service-back/usecase/interactor"
	"service-back/validator"

	log "github.com/sirupsen/logrus"
)

type userParam struct {
	Name     string
	Email    string
	Password string
}

type UserController struct {
	inputport inputport.UserInputPort
	param     *userParam
}

func NewUserController() *UserController {
	param := &userParam{}
	initializeParam(param)
	return &UserController{
		inputport: interactor.NewUserInteractor(
			presenter.NewUserPresenter(),
			database.NewUserDatabase(),
		),
		param: param,
	}
}

// SignUp ...
// @summary
// @description
// @tags User
// @accept mpfd
// @produce json
// @param name formData string true "name"
// @param email formData string true "email"
// @param password formData string true "password"
// @success 200
// @failure 400
// @router /signup [post]
func (ctrl *UserController) Signup(c Context) error {
	name := c.FormValue(ctrl.param.Name)
	email := c.FormValue(ctrl.param.Email)
	password := c.FormValue(ctrl.param.Password)
	iUser := &inputdata.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	err := validator.Validate(iUser)
	if err != nil {
		log.Error(err)
		return err
	}
	err = ctrl.inputport.Signup(iUser)
	if err != nil {
		log.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, nil)
}
