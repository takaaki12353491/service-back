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

type CommunityController struct {
	inputport inputport.CommunityInputPort
}

func NewCommunityController() *CommunityController {
	return &CommunityController{
		inputport: interactor.NewCommunityInteractor(
			presenter.NewCommunityPresenter(),
			datastore.NewCommunityDatastore(),
			datastore.NewUserDatastore(),
		),
	}
}

// Index ...
// @summary
// @description
// @tags Community
// @accept json
// @produce json
// @success 200 {array} outputdata.Community ""
// @failure 400 {string} string ""
// @router /communities [get]
func (ctrl *CommunityController) Index(c Context) error {
	oCommunities, err := ctrl.inputport.Index()
	if err != nil {
		log.Error(err)
		c.JSON(errs.StatusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, oCommunities)
}

// Show ...
// @summary
// @description
// @tags Community
// @accept json
// @produce json
// @param id path string true ""
// @success 200 {object} outputdata.Community ""
// @failure 400 {string} string ""
// @router /communities/{id} [get]
func (ctrl *CommunityController) Show(c Context) error {
	id := c.Param(pn.ID)
	oCommunity, err := ctrl.inputport.Show(id)
	if err != nil {
		log.Error(err)
		c.JSON(errs.StatusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, oCommunity)
}

// Create ...
// @summary
// @description
// @tags Community
// @accept mpfd
// @produce json
// @param Authorization header string true "jwt token"
// @param name formData string true "name"
// @param description formData string false "description"
// @success 200
// @failure 400
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
	err = ctrl.inputport.Create(iCommunity)
	if err != nil {
		log.Error(err)
		c.JSON(errs.StatusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, nil)
}
