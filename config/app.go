package config

import (
	"admin_golang/pkg/auth/jwt"
	"log"

	"go.uber.org/dig"

	handlerDig "admin_golang/handlers"
	repoImpl "admin_golang/repo/impl"
	serverImpl "admin_golang/services/impl"
)

func BuildDig() *dig.Container {
	container := dig.New()
	//auth, err := InitAuth()
	_ = container.Provide(func() jwt.JWTAuth {
		return jwt.NewJWTAuth()
	})
	err := repoImpl.Inject(container)
	if err != nil {
		log.Print("Inject Repo To Container Error")
	}
	err = serverImpl.Inject(container)
	if err != nil {
		log.Print("Inject Server To Container Error")
	}
	err = handlerDig.Inject(container)
	if err != nil {
		log.Print("Inject Handler To Container Error")
	}
	return container
}
