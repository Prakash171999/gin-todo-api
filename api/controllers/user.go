package controllers

import (
	"boilerplate-api/api/responses"
	"boilerplate-api/api/services"
	"boilerplate-api/constants"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserController -> struct
type UserController struct {
	logger         infrastructure.Logger
	userService    services.UserService
	env            infrastructure.Env
}

// NewUserController -> constructor
func NewUserController(
	logger infrastructure.Logger,
	userService services.UserService,
	env infrastructure.Env,
) UserController {
	return UserController{
		logger:         logger,
		userService:    userService,
		env:            env,
	}
}

// CreateUser -> Create User
func (cc UserController) CreateUser(c *gin.Context) {
	user := models.User{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBind(&user); err != nil {
		cc.logger.Zap.Error("Error [CreateUser] (ShouldBindJson) : ", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed To Create User")
		return
	}

	if err := cc.userService.WithTrx(trx).CreateUser(user); err != nil {
		cc.logger.Zap.Error("Error [CreateUser] [db CreateUser]: ", err.Error())
		responses.ErrorJSON(c, http.StatusInternalServerError, "Failed To Create User")
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "User Created Sucessfully")
}

// GetAllUser -> Get All User
func (cc UserController) GetAllUser(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	users, count, err := cc.userService.GetAllUser(pagination)

	if err != nil {
		cc.logger.Zap.Error("Error finding user records", err.Error())
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to Find users")
		return
	}

	responses.JSONCount(c, http.StatusOK, users, count)
}
