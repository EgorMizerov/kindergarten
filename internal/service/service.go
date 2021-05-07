package service

import (
	"github.com/EgorMizerov/kindergarten/internal/repository"
	"github.com/EgorMizerov/kindergarten/internal/storage"
	"github.com/EgorMizerov/kindergarten/pkg/auth"
)

type Service struct {
}

func NewService(repo *repository.Repository, storage storage.Storage, auth auth.TokenManager) *Service {
	return &Service{}
}
