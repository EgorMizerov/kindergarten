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
}

type Repository struct {
	User User
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		User: NewUserMongo(db.Collection("users")),
	}
}

func convertInsertedIDToString(insertedID interface{}) (string, error) {
	objId, ok := insertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("error converting InsertedID to ObjectID")
	}

	id := strings.Split(objId.String(), "\"")[1]

	return id, nil
}
