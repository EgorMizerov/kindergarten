package repository

import (
	"context"
	"errors"
	"github.com/EgorMizerov/kindergarten/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type User interface {
	CreateUser(ctx context.Context, domain domain.User) (string, error)
	GetUserById(ctx context.Context, id string) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
}

type RefreshToken interface {
	SetRefreshTokenByID(ctx context.Context, session domain.Session, id string) error
	GetRefreshTokenById(ctx context.Context, id string) (domain.Session, error)
}

type Repository struct {
	User         User
	RefreshToken RefreshToken
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		User:         NewUserMongo(db.Collection("users")),
		RefreshToken: NewRefreshTokenMongo(db.Collection("users")),
	}
}

func ConvertInsertedIDToString(insertedID interface{}) (string, error) {
	objId, ok := insertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("error converting InsertedID to ObjectID")
	}

	id := strings.Split(objId.String(), "\"")[1]

	return id, nil
}
