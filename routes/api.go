package routes

import (
	handler "admin_golang/handlers"
	"admin_golang/pkg/auth/jwt"
	httpResponse "admin_golang/pkg/http_response"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func CreateApi(engine *gin.Engine, container *dig.Container) error {
	err := container.Invoke(
		func(jwt jwt.JWTAuth,
			auth *handler.Auth) error {
			{
				engine.POST("/login", httpResponse.Wrap(auth.Login))
				engine.POST("/register", httpResponse.Wrap(auth.Register))
				engine.POST("/token/refresh", httpResponse.Wrap(auth.Refresh))
				engine.POST("/logout", httpResponse.Wrap(auth.LogOut))
			}
			engine.Group("/admin")
			return nil
		},
	)
	return err
}
