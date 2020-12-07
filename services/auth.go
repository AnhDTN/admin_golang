package services

import (
	"admin_golang/models"
	"admin_golang/schema"
	"context"
)

type IAuthService interface {
	Login(ctx context.Context, param *schema.LoginBodyParam) (*models.User, error)
	Register(ctx context.Context, param *schema.RegisterBodyParam) (*models.User, error)
	RefreshToken(ctx context.Context, bodyParam *schema.RefreshTokenParam) (*schema.UserTokenInfo, error)
	Logout(ctx context.Context) error
}
