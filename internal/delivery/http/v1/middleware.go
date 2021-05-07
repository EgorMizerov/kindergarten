package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "empty Authorization header")
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "invalid auth header")
			return
		}

		if len(authHeaderParts[1]) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "token is empty")
			return
		}

		sub, role, err := h.tokenManager.Parse(authHeaderParts[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Set("sub", sub)
		ctx.Set("role", role)

		ctx.Next()
	}
}
