package repository

import (
	"context"
	"errors"
	"github.com/EgorMizerov/kindergarten/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserMongo struct {
	coll *mongo.Collection
}

func NewUserMongo(coll *mongo.Collection) *UserMongo {
	return &UserMongo{coll: coll}
}

func (r *UserMongo) CreateUser(ctx context.Context, domain domain.User) (string, error) {
	res, err := r.coll.InsertOne(ctx, bson.M{
		"passwordHash": domain.PasswordHash,
		"email":        domain.Email,
		"firstName":    domain.FirstName,
		"middleName":   domain.MiddleName,
		"lastName":     domain.LastName,
		"dateOfBirth":  domain.DateOfBirth,
		"isAdmin":      domain.IsAdmin,
		"creationDate": time.Now().Unix(),
	})
	if err != nil {
		return "", err
	}

	objId, err := convertInsertedIDToString(res.InsertedID)
	if err != nil {
		return "", errors.New("error converting InsertedID to ObjectID")
	}

	return objId, err
}

func (r *UserMongo) GetUserById(ctx context.Context, id string) (domain.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, err
	}

	res := r.coll.FindOne(ctx, bson.M{"_id": objId})
	if res.Err() != nil {
		return domain.User{}, nil
	}

	var user domain.User
	err = res.Decode(&user)

	return user, err
}
