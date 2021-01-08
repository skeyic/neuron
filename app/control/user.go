package control

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/glog"
	"github.com/skeyic/neuron/app/model"
	"github.com/skeyic/neuron/utils"
	"net/http"
)

// @Summary Add a new user
// @Description add a new user
// @Accept json
// @Produce json
// @Param body  body model.UserInput true "The user input info"
// @Success 200 {object} model.User	"The user info"
// @Failure 400 {object} utils.WebError "Bad request"
// @Failure 500 {object} utils.WebError "Internal error"
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
