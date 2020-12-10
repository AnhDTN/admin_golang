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
			auth *handler.Auth,
			user *handler.User) error {
			versionApi := engine.Group("/api/v1")
			authentic := versionApi.Group("/auth")
			{
				authentic.POST("/login", httpResponse.Wrap(auth.Login))
				authentic.POST("/register", httpResponse.Wrap(auth.Register))
				authentic.POST("/token/refresh", httpResponse.Wrap(auth.Refresh))
				authentic.POST("/logout", httpResponse.Wrap(auth.LogOut))
			}

			// U can save map[string]interface account
			internal := versionApi.Group("/internal", gin.BasicAuth(gin.Accounts{
				"NamAnh":     "123",
				"NamAnh123":  "123",
				"NamAnh1234": "123",
			}))
			{
				internal.GET("/users", httpResponse.Wrap(user.GetAllUser))
			}
			return nil
		},
	)
	return err
}
