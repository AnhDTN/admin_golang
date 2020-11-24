package repo

import (
	"admin_golang/models"
	"admin_golang/schema"
)

type UserRepository interface {
	CreateUser(admin *schema.RegisterBodyParam) (*models.User, error)
	DeleteUser(admin *schema.AdminIdParam) error
	UpdateUser(admin *schema.AdminUpdateParam) (*models.User, error)
	GetUserById(param *schema.AdminIdParam) (*models.User, error)
	GetUserByToken(token string) (*models.User, error)
	DeleteUserToken(param *schema.AdminIdParam) error
	GetAllUser() ([]*models.User, error)
	RefreshToken(token string) (*models.User, error)
}
