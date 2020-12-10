package repo

import (
	"admin_golang/models"
	"admin_golang/schema"
)

type IBasicUserRepo interface {
	Create(param *schema.BasicUserParam) error
	GetAll() ([]*models.BasicUser, error)
}
