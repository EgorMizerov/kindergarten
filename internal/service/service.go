package service

import (
	"context"
	"github.com/EgorMizerov/kindergarten/internal/domain"
	"github.com/EgorMizerov/kindergarten/internal/repository"
	"github.com/EgorMizerov/kindergarten/internal/storage"
	"github.com/EgorMizerov/kindergarten/pkg/auth"
	"io"
)

type User interface {
	CreateUser(ctx context.Context, domain domain.User) (string, error)
	GetUserById(ctx context.Context, id string) (domain.User, error)
}

type Auth interface {
	SignIn(ctx context.Context, user domain.User) (string, string, error)
	SignUp(ctx context.Context, domain domain.User) (string, error)
	Refresh(ctx context.Context, id, token string) (string, string, error)
	Logout(ctx context.Context, id string) error
}

type ProfilePhoto interface {
	SetProfilePhoto(ctx context.Context, id string, r io.Reader) error
	GetProfilePhoto(ctx context.Context, id string, w io.Writer) error
}

type Service struct {
	user         User
	ProfilePhoto ProfilePhoto
	Auth         Auth
}

func NewService(repo *repository.Repository, storage storage.Storage, auth auth.TokenManager) *Service {
	return &Service{
		user:         NewUserService(repo.User),
		ProfilePhoto: NewProfilePhotoService(storage),
		Auth:         NewAuthService(repo.User, repo.RefreshToken, auth),
	}
}
