package handlers

import (
	"admin_golang/models"
	"admin_golang/pkg/error_custom"
	"admin_golang/schema"
	"admin_golang/services"

	"github.com/gin-gonic/gin"
)

type BasicUser struct {
	service services.IBasicUserService
}

func NewBasicUser(service services.IBasicUserService) *BasicUser {
	return &BasicUser{
		service: service,
	}
}

func (basic *BasicUser) CreateBasicUser(c *gin.Context) models.ResponseData {
	var param schema.BasicUserParam
	if err := c.ShouldBindJSON(&param); err != nil {
		return models.ResponseData{
			Data:  nil,
			Error: error_custom.InvalidParams.New(),
		}
	}
	err := basic.service.CreateBasicUser(&param)
	if err != nil {
		return models.ResponseData{
			Data:  nil,
			Error: err,
		}
	}
	return models.ResponseData{
		Data:  nil,
		Error: error_custom.Success.New(),
	}
}

func (basic *BasicUser) GetAllBasicUser() models.ResponseData {
	users, err := basic.service.GetAllBasicUser()
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
