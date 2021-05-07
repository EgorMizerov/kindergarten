package service

import "github.com/EgorMizerov/kindergarten/internal/repository"

type Service struct {
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
