package service

import (
	"context"
	"errors"
	"github.com/EgorMizerov/kindergarten/internal/domain"
	"github.com/EgorMizerov/kindergarten/internal/repository"
	"github.com/EgorMizerov/kindergarten/pkg/auth"
	"github.com/EgorMizerov/kindergarten/pkg/hash"
	"github.com/google/uuid"
	"time"
)

type AuthService struct {
	userRepo     repository.User
	refreshRepo  repository.RefreshToken
	tokenManager auth.TokenManager
}

func NewAuthService(userRepo repository.User, refreshRepo repository.RefreshToken, manager auth.TokenManager) *AuthService {
	return &AuthService{userRepo: userRepo, refreshRepo: refreshRepo, tokenManager: manager}
}

func (s *AuthService) SignIn(ctx context.Context, input domain.User) (string, string, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return "", "", err
	}

	if user.PasswordHash != hash.PasswordHash(input.PasswordHash) {
		return "", "", errors.New("wrong password")
	}

	id, err := repository.ConvertInsertedIDToString(user.Id)
	if err != nil {
		return "", "", errors.New("invalid id")
	}

	accessToken, err := s.tokenManager.NewJWT(time.Hour*2, id, "user") // 2 hour
	refreshToken := uuid.New().String()

	err = s.refreshRepo.SetRefreshTokenByID(ctx, domain.Session{
		Token: refreshToken,
		Exp:   time.Now().Add(time.Hour * 1440).Unix(), // 60 days
	}, id)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err
}

func (s *AuthService) SignUp(ctx context.Context, domain domain.User) (string, error) {
	domain.PasswordHash = hash.PasswordHash(domain.PasswordHash)
	return s.userRepo.CreateUser(ctx, domain)
}

func (s *AuthService) Refresh(ctx context.Context, id, token string) (string, string, error) {
	currentToken, err := s.refreshRepo.GetRefreshTokenById(ctx, id)
	if err != nil {
		return "", "", err
	}

	if currentToken.Token != token {
		return "", "", errors.New("invalid token")
	}

	if currentToken.Exp < time.Now().Unix() {
		return "", "", errors.New("token expired")
	}

	accessToken, err := s.tokenManager.NewJWT(time.Hour*2, id, "user") // 2 hour
	refreshToken := uuid.New().String()

	err = s.refreshRepo.SetRefreshTokenByID(ctx, domain.Session{
		Token: refreshToken,
		Exp:   time.Now().Add(time.Hour * 1440).Unix(), // 60 days
	}, id)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err
}

func (h *AuthService) Logout(ctx context.Context, id string) error {
	return h.refreshRepo.SetRefreshTokenByID(ctx, domain.Session{
		Token: uuid.New().String(),
		Exp:   time.Now().Unix(),
	}, id)
}
