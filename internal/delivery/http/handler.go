package http

import (
	v1 "github.com/EgorMizerov/kindergarten/internal/delivery/http/v1"
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

func (h *Handler) Init() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	handlerV1 := v1.NewHandler(h.service)
	api := router.Group("/api")
	{
		handlerV1.InitAPIV1(api)
	}

	router.GET("/favicon.ico", func(ctx *gin.Context) { ctx.File("./favicon.ico") })

	return router
}
