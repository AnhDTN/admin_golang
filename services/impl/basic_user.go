package impl

import (
	"admin_golang/models"
	"admin_golang/repo"
	"admin_golang/schema"
	"admin_golang/services"
)

type BasicUserService struct {
	service repo.IBasicUserRepo
}

func NewBasicUserService(service repo.IBasicUserRepo) services.IBasicUserService {
	return &BasicUserService{service: service}
}

func (b BasicUserService) CreateBasicUser(body *schema.BasicUserParam) error {
	err := b.service.Create(body)
	if err != nil {
		return err
	}
	return nil
}

func (b BasicUserService) GetAllBasicUser() ([]*models.BasicUser, error) {
	users, err := b.service.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
