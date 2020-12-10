package repo

import (
	"admin_golang/models"
	"admin_golang/schema"
)

type IRoleRepository interface {
	CreateRole(role *schema.RoleBodyParam) (*models.Role, error)
	DeleteRole(role *schema.DeleteBodyParam) error
	UpdateRole(role *schema.Role) error
	GetRoleById(id string) (*models.Role, error)
}
