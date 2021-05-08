package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) mediaType() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/vnd.api+json")

		if ctx.ContentType() != "application/vnd.api+json" {
			ctx.AbortWithStatus(http.StatusUnsupportedMediaType)
			return
		}

		if ctx.GetHeader("Accept") == "" {
			ctx.AbortWithStatus(http.StatusNotAcceptable)
			return
		}

		ctx.Next()
	}
}

func (h *Handler) corsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		ctx.Next()
	}
}

func (h *Handler) optionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
			return
		}
		ctx.Next()
	}
}
