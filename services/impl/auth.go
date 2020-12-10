package impl

import (
	"admin_golang/models"
	"admin_golang/pkg/auth/jwt"
	"admin_golang/repo"
	"admin_golang/schema"
	"admin_golang/services"
	"context"
	"log"
)

type AuthService struct {
	jwt      jwt.JWTAuth
	userRepo repo.IUserRepository
	roleRepo repo.IRoleRepository
}

func NewAuthService(jwt jwt.JWTAuth, userRepo repo.IUserRepository, roleRepo repo.IRoleRepository) services.IAuthService {
	return &AuthService{jwt: jwt, userRepo: userRepo, roleRepo: roleRepo}
}

func (a *AuthService) Login(ctx context.Context, param *schema.LoginBodyParam) (*models.User, error) {
	user, err := a.userRepo.Login(param)
	if err != nil {
		return nil, err
	}
	token, err := a.jwt.GenerateToken(user.Id)
	if err != nil {
		log.Print("jwt.GenerateToken Error: ", err)
		return nil, err
	}
	values := schema.UserUpdateBodyParam{RefreshToken: token.GetRefreshToken()}
	err = a.userRepo.UpdateRefreshToken(&values)
	if err != nil {
		log.Print("userRepo.UpdateRefreshToken Error: ", err)
		return nil, err
	}
	user.PassWord = ""
	user.AccessToken = token.GetTokenType() + " " + token.GetAccessToken()
	user.RefreshToken = token.GetTokenType() + " " + token.GetRefreshToken()
	return user, nil
}

func (a *AuthService) Register(ctx context.Context, param *schema.RegisterBodyParam) (*models.User, error) {

	if param.RoleId == "" {
		role, err := a.roleRepo.GetRoleById("1")

		if err != nil {
			return nil, err
		}
		param.RoleId = role.Id
	}
	user, err := a.userRepo.CreateUser(param)

	if err != nil {
		log.Print("userRepo.CreateUser Error: ", err)
		return nil, err
	}

	token, err := a.jwt.GenerateToken(user.Id)
	if err != nil {
		log.Print("jwt.GenerateToken Error: ", err)
		return nil, err
	}
	values := schema.UserUpdateBodyParam{RefreshToken: token.GetRefreshToken()}
	err = a.userRepo.UpdateRefreshToken(&values)
	if err != nil {
		log.Print("userRepo.UpdateRefreshToken Error: ", err)
		return nil, err
	}
	tokenInfo := schema.UserTokenInfo{
		AccessToken:  token.GetAccessToken(),
		RefreshToken: token.GetRefreshToken(),
		TokenType:    token.GetTokenType(),
	}
	user.RefreshToken = tokenInfo.TokenType + " " + tokenInfo.RefreshToken
	user.AccessToken = tokenInfo.TokenType + " " + tokenInfo.AccessToken
	return user, nil
}

func (a *AuthService) RefreshToken(ctx context.Context, bodyParam *schema.RefreshTokenParam) (*schema.UserTokenInfo, error) {
	user, err := a.userRepo.GetUserByToken(bodyParam.RefreshToken)
	if err != nil {
		return nil, err
	}
	token, err := a.jwt.RefreshToken(bodyParam.RefreshToken)
	if err != nil {
		return nil, err
	}
	body := schema.UserUpdateBodyParam{Id: user.Id, RefreshToken: bodyParam.RefreshToken}
	err = a.userRepo.UpdateRefreshToken(&body)
	tokenInfo := schema.UserTokenInfo{
		RefreshToken: token.GetRefreshToken(),
		AccessToken:  token.GetAccessToken(),
		TokenType:    token.GetTokenType(),
	}
	return &tokenInfo, nil
}

func (a *AuthService) Logout(ctx context.Context) error {
	//_, err := a.userRepo.DeleteUserToken()
	//if err != nil {
	//	return err
	//}

	return nil
}
