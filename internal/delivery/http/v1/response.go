package v1

import (
	"github.com/gin-gonic/gin"
	"log"
)

type response struct {
	Message string
}

func newResponse(ctx *gin.Context, statusCode int, message string) {
	log.Print(message)
	ctx.AbortWithStatusJSON(statusCode, response{message})
}
