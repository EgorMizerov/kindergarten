package v1

import "github.com/EgorMizerov/kindergarten/internal/service"

type Handler struct {
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{}
}
