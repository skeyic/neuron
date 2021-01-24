package control

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/glog"
	"github.com/skeyic/neuron/app/model"
	"github.com/skeyic/neuron/utils"
	"net/http"
)

// @Summary Add a new user
// @Tags User
// @Description add a new user
// @Accept json
// @Produce json
// @Param body  body model.UserInput true "The user input info"
// @Success 200 {object} model.User	"The user info"
// @Failure 400 {object} utils.WebResponse "Bad request"
// @Failure 500 {object} utils.WebResponse "Internal error"
// @Router /users [post]
func NewUser(c *gin.Context) {
	var (
		err       error
		userInput = &model.UserInput{}
	)

	err = c.ShouldBindWith(userInput, binding.JSON)
	if err != nil {
		glog.Errorf("Bind request json error: %v", err)
		utils.NewBadRequestError(c, "Bind request json error")
		return
	}

	err = userInput.Validate()
	if err != nil {
		glog.Errorf("Validate user input error: %v", err)
		utils.NewBadRequestError(c, "Validate user input error")
		return
	}

	user := userInput.ToUser()
	err = user.Save()
	if err != nil {
		glog.Errorf("Save user error: %v", err)
		utils.NewBadRequestError(c, "Save user error")
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary Get all users
// @Tags User
// @Description get all users
// @Accept json
// @Produce json
// @Success 200 {object} []model.User	"All the user infos"
// @Failure 400 {object} utils.WebResponse "Bad request"
// @Failure 500 {object} utils.WebResponse "Internal error"
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var (
		users []*model.User
	)

	users = model.TheUsersMaster.GetUsers()

	c.JSON(http.StatusOK, users)
}

// @Summary Get user by ID
// @Tags User
// @Description get user by ID
// @Accept json
// @Produce json
// @Param id path string true "The user ID"
// @Success 200 {object} model.User	"The user info"
// @Failure 400 {object} utils.WebResponse "Bad request"
// @Failure 404 {object} utils.WebResponse "Not found"
// @Failure 500 {object} utils.WebResponse "Internal error"
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	var (
		user *model.User
	)

	ID := c.Param("id")

	user = model.TheUsersMaster.GetUser(ID)
	if user == nil {
		glog.Errorf("user %s not found", ID)
		utils.NewNotFoundError(c, fmt.Sprintf("user %s not found", ID))
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary Add BarkAlert postman to user
// @Tags User
// @Description send alert to user
// @Accept json
// @Produce json
// @Param id path string true "The user ID"
// @Param alertBody body model.BarkAlert true "The bark alert postman"
// @Success 200 {object} model.User "The user after add bark alert postman"
// @Failure 400 {object} utils.WebResponse "Bad request"
// @Failure 404 {object} utils.WebResponse "Not found"
// @Failure 500 {object} utils.WebResponse "Internal error"
// @Router /users/{id}/bark [post]
func AddBarkAlertPostmanToUser(c *gin.Context) {
	var (
		err       error
		user      *model.User
		barkAlert = &model.BarkAlert{}
	)

	ID := c.Param("id")

	user = model.TheUsersMaster.GetUser(ID)
	if user == nil {
		glog.Errorf("user %s not found", ID)
		utils.NewNotFoundError(c, fmt.Sprintf("user %s not found", ID))
		return
	}

	err = c.ShouldBindWith(&barkAlert, binding.JSON)
	if err != nil {
		glog.Error("failed to load alert")
		utils.NewBadRequestError(c, "failed to load alert")
		return
	}

	err = user.NewAlertService(barkAlert)
	if err != nil {
		glog.Error("failed to save new alert")
		utils.NewBadRequestError(c, "failed to save new alert")
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary Send alert to user
// @Tags User
// @Description send alert to user
// @Accept json
// @Produce json
// @Param id path string true "The user ID"
// @Param alertBody body model.AlertBody true "The alert body"
// @Success 200 {object} utils.WebResponse "Ok"
// @Failure 400 {object} utils.WebResponse "Bad request"
// @Failure 404 {object} utils.WebResponse "Not found"
// @Failure 500 {object} utils.WebResponse "Internal error"
// @Router /users/{id}/send [post]
func SendAlertToUser(c *gin.Context) {
	var (
		err       error
		user      *model.User
		alertBody = &model.AlertBody{}
	)

	ID := c.Param("id")

	user = model.TheUsersMaster.GetUser(ID)
	if user == nil {
		glog.Errorf("user %s not found", ID)
		utils.NewNotFoundError(c, fmt.Sprintf("user %s not found", ID))
		return
	}

	err = c.ShouldBindWith(&alertBody, binding.JSON)
	if err != nil {
		glog.Error("failed to load alert")
		utils.NewBadRequestError(c, "failed to load alert")
		return
	}

	user.Send(alertBody)

	c.JSON(http.StatusOK, user)
}
