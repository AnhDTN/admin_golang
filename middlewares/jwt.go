package middlewares

import (
	"admin_golang/models"
	jwt2 "admin_golang/pkg/auth/jwt"
	"admin_golang/pkg/error_custom"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func validateToken(accessToken string) (map[string]interface{}, error) {
	jwtToken := strings.Replace(accessToken, "Bearer ", "", -1)
	tokenData := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(jwtToken, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwt2.DefaultRefreshKey), nil
	})
	if err != nil {
		log.Print("Parse Token With Claims Error: ", err)
		return nil, err
	}
	if !token.Valid {
		log.Print("Token Invalid")
		return nil, errors.New("token expire")
	}
	var result map[string]interface{}
	copier.Copy(&result, tokenData["payload"])
	return result, nil
}

func JWTHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			handleResponse(c, http.StatusUnauthorized, models.ResponseData{Data: nil, Error: error_custom.ErrorTokenInvalid.New()}, error_custom.ErrorTokenInvalid.New())
			c.Abort()
			return
		}
		_, err := validateToken(token)
		if err != nil {
			if err.Error() == "token expire" {
				handleResponse(c, http.StatusUnauthorized, nil, err)
			} else {
				handleResponse(c, http.StatusUnauthorized, nil, err)
			}
			c.Abort()
			return
		}
		c.Next()
	}
}

func handleResponse(c *gin.Context, status int, data interface{}, err error) {
	c.JSON(status, models.ResponseResult(data, err))
}
