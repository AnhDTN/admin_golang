package http_response

import (
	"admin_golang/models"
	"admin_golang/pkg/error_custom"

	"github.com/gin-gonic/gin"
)

const (
	DataField    = "data"
	CodeField    = "code"
	MessageField = "message"
)

type GinHandlerFn func(c *gin.Context) models.ResponseData

func Wrap(fn GinHandlerFn) gin.HandlerFunc {
	return func(c *gin.Context) {
		// handle req
		res := fn(c)
		Translate(c, res)
	}
}

func Translate(c *gin.Context, response models.ResponseData) {
	result := gin.H{}
	customError, ok := response.Error.(error_custom.CustomError)
	if ok {
		result[MessageField] = customError.Error()
		result[CodeField] = customError.Type()
	} else {
		result[MessageField] = response.Error.Error()
		result[CodeField] = -1
	}

	result[DataField] = response.Data

	c.JSON(200, result)
}
