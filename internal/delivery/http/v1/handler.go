package v1

import (
	"github.com/EgorMizerov/kindergarten/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitAPIV1(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		v1.GET("/version", func(ctx *gin.Context) {
			ctx.String(200, "v1")
		})
	}
}
