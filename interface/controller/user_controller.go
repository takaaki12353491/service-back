package controller

import (
	"net/http"
	"service-back/errs"
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

// Signup ...
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
		c.String(errs.StatusCode(err), errs.Cause(err).Error())
		return err
	}
	err = ctrl.inputport.Signup(iUser)
	if err != nil {
		log.Error(err)
		c.String(errs.StatusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

// Login ...
// @summary Login
// @description Generate cookie for login discrimination if email and password match DB. Otherwise redirect to sign in page.
// @tags UserAuth
// @accept mpfd
// @produce json
// @param name formData string true "name"
// @param password formData string true "password"
// @success 200 {object} outputdata.Login ""
// @failure 409 {string} string ""
// @router /login [post]
func (ctrl *UserController) Login(c Context) error {
	name := c.FormValue(ctrl.param.Name)
	password := c.FormValue(ctrl.param.Password)
	iSignIn := &inputdata.Login{
		Name:     name,
		Password: password,
	}
	oLogin, err := ctrl.inputport.Login(iSignIn)
	if err != nil {
		log.Error(err)
		c.String(errs.StatusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, oLogin)
}
