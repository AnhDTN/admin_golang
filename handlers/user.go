package handlers

import (
	"admin_golang/models"
	"admin_golang/pkg/error_custom"
	"admin_golang/services"

	"github.com/gin-gonic/gin"
)

type User struct {
	user services.IUserService
}

func NewUserHandler(user services.IUserService) *User {
	return &User{user: user}
}

func (u User) GetAllUser(c *gin.Context) models.ResponseData {
	_ = c.MustGet(gin.AuthUserKey).(string)
	users, err := u.user.GetAll()
	if err != nil {
		return models.ResponseData{
			Data:  nil,
			Error: err,
		}
	}
	return models.ResponseData{
		Data:  users,
		Error: error_custom.Success.New(),
	}
}
