package services

import (
	"admin_golang/models"
	"admin_golang/schema"
)

type IUserService interface {
	Create(body *schema.RegisterBodyParam) (*models.User, error)
	Delete(body *schema.UserIdParam) error
	Update(body *schema.UserUpdateParam) (*models.User, error)
	GetById(body *schema.UserIdParam) (*models.User, error)
	GetByToken(token string) (*models.User, error)
	GetAll() ([]*models.User, error)
}
