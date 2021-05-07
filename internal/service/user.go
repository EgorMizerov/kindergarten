package service

import (
	"context"
	"github.com/EgorMizerov/kindergarten/internal/domain"
	"github.com/EgorMizerov/kindergarten/internal/repository"
)

type UserService struct {
	userRepo repository.User
}

func NewUserService(userRepo repository.User) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(ctx context.Context, domain domain.User) (string, error) {
	return s.userRepo.CreateUser(ctx, domain)
}

func (s *UserService) GetUserById(ctx context.Context, id string) (domain.User, error) {
	return s.userRepo.GetUserById(ctx, id)
}
