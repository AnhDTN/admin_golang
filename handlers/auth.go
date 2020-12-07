package handlers

import (
	"admin_golang/models"
	"admin_golang/pkg/error_custom"
	"admin_golang/schema"
	"admin_golang/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Auth struct {
	service services.IAuthService
}

func NewAuthHandler(service services.IAuthService) *Auth {
	return &Auth{service: service}
}

func (auth *Auth) Login(c *gin.Context) models.ResponseData {
	var item schema.LoginBodyParam

	if err := c.ShouldBindJSON(&item); err != nil {
		return models.ResponseData{
			Error: error_custom.InvalidParams.New(),
			Data:  nil,
		}
	}

	ctx := c.Request.Context()
	tokenInfo, err := auth.service.Login(ctx, &item)
	if err != nil {
		return models.ResponseData{
			Error: err,
			Data:  nil,
		}
	}

	return models.ResponseData{
		Error: error_custom.Success.New(),
		Data:  tokenInfo,
	}
}

func isValidator() {

}
func (auth *Auth) Register(c *gin.Context) models.ResponseData {
	var item schema.RegisterBodyParam
	if err := c.ShouldBindJSON(&item); err != nil {
		return models.ResponseData{
			Data:  nil,
			Error: error_custom.InvalidParams.New(),
		}
	}
	err := validator.New().Struct(item)

	if err != nil {
		return models.ResponseData{
			Data:  nil,
			Error: error_custom.InvalidParams.New(),
		}
	}

	ctx := c.Request.Context()
	user, err := auth.service.Register(ctx, &item)
	if err != nil {
		return models.ResponseData{
			Data:  nil,
			Error: err,
		}
	}
	return models.ResponseData{
		Data:  user,
		Error: error_custom.Success.New(),
	}
}

func (auth *Auth) Refresh(c *gin.Context) models.ResponseData {
	var bodyParam schema.RefreshTokenParam
	if err := c.ShouldBindJSON(&bodyParam); err != nil {
		return models.ResponseData{
			Data:  nil,
			Error: error_custom.InvalidParams.New(),
		}
	}
	validate := validator.New()
	if err := validate.Struct(bodyParam); err != nil {
		return models.ResponseData{
			Data:  nil,
			Error: error_custom.InvalidParams.New(),
		}
	}
	ctx := c.Request.Context()
	tokenInfo, err := auth.service.RefreshToken(ctx, &bodyParam)
	if err != nil {
		return models.ResponseData{
			Data:  nil,
			Error: err,
		}
	}
	return models.ResponseData{
		Data:  tokenInfo,
		Error: error_custom.Success.New(),
	}
}
func (auth *Auth) LogOut(c *gin.Context) models.ResponseData {
	err := auth.service.Logout(c)
	if err != nil {
		return models.ResponseData{
			Data:  nil,
			Error: error_custom.ErrorBadRequest.New(),
		}
	}
	return models.ResponseData{
		Data:  nil,
		Error: error_custom.Success.New(),
	}
}
