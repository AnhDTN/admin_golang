package impl

import (
	"admin_golang/models"
	"admin_golang/repo"
	"admin_golang/schema"
	"admin_golang/services"
)

type UserService struct {
	userRepo repo.IUserRepository
	roleRepo repo.IRoleRepository
}

func NewUserService(userRepo repo.IUserRepository,
	roleRepo repo.IRoleRepository) services.IUserService {
	return &UserService{userRepo: userRepo, roleRepo: roleRepo}
}

func (u UserService) Create(body *schema.RegisterBodyParam) (*models.User, error) {
	user, err := u.userRepo.CreateUser(body)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserService) Delete(body *schema.UserIdParam) error {
	err := u.userRepo.DeleteUser(body)
	if err != nil {
		return err
	}
	return nil
}

func (u UserService) Update(body *schema.UserUpdateParam) (*models.User, error) {
	user, err := u.userRepo.UpdateUser(body)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserService) GetById(body *schema.UserIdParam) (*models.User, error) {
	user, err := u.userRepo.GetUserById(body)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserService) GetByToken(token string) (*models.User, error) {
	user, err := u.userRepo.GetUserByToken(token)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserService) GetAll() ([]*models.User, error) {
	user, err := u.userRepo.GetAllUser()
	if err != nil {
		return nil, err
	}
	return user, nil
}
