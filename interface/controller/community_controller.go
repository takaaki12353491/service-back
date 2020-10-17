package controller

import (
	"net/http"
	"service-back/errs"
	"service-back/interface/database"
	"service-back/interface/presenter"
	"service-back/usecase/inputdata"
	"service-back/usecase/inputport"
	"service-back/usecase/interactor"
	"service-back/validator"

	log "github.com/sirupsen/logrus"
)

type CommunityController struct {
	inputport inputport.CommunityInputPort
}

func NewCommunityController() *CommunityController {
	return &CommunityController{
		inputport: interactor.NewCommunityInteractor(
			presenter.NewCommunityPresenter(),
			database.NewCommunityDatabase(),
			database.NewUserDatabase(),
		),
	}
}

// Create ...
// @summary
// @description
// @tags Community
// @accept mpfd
// @produce json
// @param name formData string true "name"
// @param description formData string true "description"
// @param password formData string true "password"
// @success 200 {object} outputdata.Community ""
// @failure 400 {string} string ""
// @router /communities [post]
func (ctrl *CommunityController) Create(c Context) error {
	userID := c.UserID()
	name := c.FormValue(pn.Name)
	description := c.FormValue(pn.Description)
	iCommunity := &inputdata.Community{
		UserID:      userID,
		Name:        name,
		Description: description,
	}
	err := validator.Validate(iCommunity)
	if err != nil {
		log.Error(err)
		c.JSON(errs.StatusCode(err), errs.Cause(err).Error())
		return err
	}
	oCommunity, err := ctrl.inputport.Create(iCommunity)
	if err != nil {
		log.Error(err)
		c.JSON(errs.StatusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, oCommunity)
}
