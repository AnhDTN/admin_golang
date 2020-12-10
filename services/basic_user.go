package services

import (
	"admin_golang/models"
	"admin_golang/schema"
)

type IBasicUserService interface {
	CreateBasicUser(body *schema.BasicUserParam) error
	GetAllBasicUser() ([]*models.BasicUser, error)
}
