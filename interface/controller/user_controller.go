package controller

import (
	"net/http"
	"service-back/errs"
	"service-back/interface/datastore"
	"service-back/interface/presenter"
	"service-back/usecase/inputdata"
	"service-back/usecase/inputport"
	"service-back/usecase/interactor"
	"service-back/validator"

	log "github.com/sirupsen/logrus"
)

type UserController struct {
	inputport inputport.UserInputPort
}

func NewUserController() *UserController {
	return &UserController{
		inputport: interactor.NewUserInteractor(
			presenter.NewUserPresenter(),
			datastore.NewUserDatastore(),
		),
	}
}

// Signup ...
// @summary
// @description
// @tags Auth
// @accept mpfd
// @produce json
// @param name formData string true "name"
// @param email formData string true "email"
// @param password formData string true "password"
// @success 200
// @failure 400
// @router /signup [post]
func (ctrl *UserController) Signup(c Context) error {
	name := c.FormValue(pn.Name)
	email := c.FormValue(pn.Email)
	password := c.FormValue(pn.Password)
	iUser := &inputdata.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	err := validator.Validate(iUser)
	if err != nil {
		log.Error(err)
		c.JSON(errs.StatusCode(err), errs.Cause(err).Error())
		return err
	}
	err = ctrl.inputport.Signup(iUser)
	if err != nil {
		log.Error(err)
		c.JSON(errs.StatusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

// Login ...
// @summary Login
// @description
// @tags Auth
// @accept mpfd
// @produce json
// @param identity formData string true "identity"
// @param password formData string true "password"
// @success 200 {object} outputdata.Login ""
// @failure 409 {string} string ""
// @router /login [post]
func (ctrl *UserController) Login(c Context) error {
	identity := c.FormValue(pn.Identity)
	password := c.FormValue(pn.Password)
	iLogin := &inputdata.Login{
		Identity: identity,
		Password: password,
	}
	oLogin, err := ctrl.inputport.Login(iLogin)
	if err != nil {
		log.Error(err)
		c.JSON(errs.StatusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, oLogin)
}
