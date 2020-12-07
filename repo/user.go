package repo

import (
	"admin_golang/models"
	"admin_golang/schema"
)

type UserRepository interface {
	Login(body *schema.LoginBodyParam) (*models.User, error)
	CreateUser(admin *schema.RegisterBodyParam) (*models.User, error)
	DeleteUser(admin *schema.UserIdParam) error
	UpdateUser(admin *schema.UserUpdateParam) (*models.User, error)
	UpdateRefreshToken(param *schema.UserUpdateBodyParam) error
	GetUserById(param *schema.UserIdParam) (*models.User, error)
	GetUserByToken(token string) (*models.User, error)
	DeleteUserToken(param *schema.UserIdParam) error
	GetAllUser() ([]*models.User, error)
	RefreshToken(token string) (*models.User, error)
}
