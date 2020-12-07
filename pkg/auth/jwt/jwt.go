package jwt

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const DefaultKey = "admin_golang_token"
const DefaultRefreshKey = "admin_golang_refresh_token"

type JWTAuth interface {
	GenerateToken(userId string) (TokenInfo, error)
	ParseToken(refreshToken string, refresh bool) (*jwt.StandardClaims, error)
	ParseUserId(refreshToken string, refresh bool) (string, error)
	RefreshToken(refreshToken string) (TokenInfo, error)
}

type options struct {
}

type JWTOptions struct {
	options *options
}

func (J *JWTOptions) GenerateToken(userId string) (TokenInfo, error) {
	accessToken, err := J.generateAccessToken(userId)
	if err != nil {
		return nil, err
	}
	refreshToken, err := J.generateRefreshToken(userId)
	if err != nil {
		return nil, err
	}
	tokenInfo := &tokenInfo{
		TokenType:    "Bearer",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return tokenInfo, nil
}

func NewJWTAuth() *JWTOptions {
	return &JWTOptions{
		options: nil,
	}
}
func (J *JWTOptions) ParseToken(refreshToken string, refresh bool) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &jwt.StandardClaims{}, func(j *jwt.Token) (interface{}, error) {
		if _, ok := j.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Print("Invalidate token")
			return nil, nil
		}
		return []byte(DefaultRefreshKey), nil
	})
	if err != nil {
		log.Print("Parse Token error_custom: ", err)
		return nil, err
	}
	return token.Claims.(*jwt.StandardClaims), nil
}

func (J *JWTOptions) ParseUserId(refreshToken string, refresh bool) (string, error) {
	claim, err := J.ParseToken(refreshToken, refresh)
	if err != nil {
		return "", err
	}
	return claim.Subject, nil
}

func (J *JWTOptions) RefreshToken(refreshToken string) (TokenInfo, error) {
	userId, err := J.ParseUserId(refreshToken, true)
	if err != nil {
		log.Print("Refresh Token Error: ", err)
		// if token expire, call Generate token again
		return nil, err
	}
	accessToken, err := J.generateAccessToken(userId)
	if err != nil {
		return nil, err
	}
	tokenInfo := &tokenInfo{
		TokenType:    "Breare",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return tokenInfo, nil
}

func (J *JWTOptions) generateAccessToken(userId string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   userId,
		ExpiresAt: now.Add(time.Duration(3000) * time.Second).Unix(),
		IssuedAt:  now.Unix(),
		NotBefore: now.Unix(),
	})
	tokenString, err := token.SignedString([]byte(DefaultKey))
	if err != nil {
		log.Print("Generate Token Error: ", err)
		return "", err
	}
	return tokenString, nil
}

func (J *JWTOptions) generateRefreshToken(userId string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   userId,
		ExpiresAt: now.Add(time.Duration(24) * time.Hour).Unix(),
		IssuedAt:  now.Unix(),
		NotBefore: now.Unix(),
	})
	tokenString, err := token.SignedString([]byte(DefaultRefreshKey))
	if err != nil {
		log.Print("Generate Token Error: ", err)
		return "", err
	}

	return tokenString, nil
}
