package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenManager interface {
	NewJWT(ttl time.Duration, id string, role string) (string, error)
	Parse(accessToken string) (sub, role string, err error)
}

const (
	UserRole  = "user"
	AdminRole = "admin"
)

type Manager struct {
	signingKey string
	issuer     string
}

func NewManager(signingKey, issuer string) *Manager {
	return &Manager{
		signingKey: signingKey,
		issuer:     issuer,
	}
}

func (m *Manager) NewJWT(ttl time.Duration, id string, role string) (string, error) {
	claims := jwt.MapClaims{
		"exp":  time.Now().Add(ttl).Unix(),
		"iss":  m.issuer,
		"iat":  time.Now().Unix(),
		"sub":  id,
		"role": role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(m.signingKey))
}

func (m *Manager) Parse(accessToken string) (sub, role string, err error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signingKey), nil
	})
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", fmt.Errorf("error get user claims from token")
	}

	sub, ok = claims["sub"].(string)
	if !ok {
		return "", "", fmt.Errorf("error get sub from token")
	}

	role, ok = claims["role"].(string)
	if !ok {
		return "", "", fmt.Errorf("error get role from token")
	}

	return
}
