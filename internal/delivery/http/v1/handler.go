package v1

import (
	"github.com/EgorMizerov/kindergarten/internal/service"
	"github.com/EgorMizerov/kindergarten/pkg/auth"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service      *service.Service
	tokenManager auth.TokenManager
}

func NewHandler(service *service.Service, manager auth.TokenManager) *Handler {
	return &Handler{
		service:      service,
		tokenManager: manager,
	}
}

func (h *Handler) InitAPIV1(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initAuthRoutes(v1)
		v1.GET("/version", func(ctx *gin.Context) {
			ctx.String(200, "v1")
		})
	}
}
