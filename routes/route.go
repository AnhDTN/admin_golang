package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func InitGinEngine(container *dig.Container) *gin.Engine {
	r := gin.Default()

	err := CreateApi(r, container)
	if err != nil {
		log.Print("Create Api Error", err)
	}
	return r
}
