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

	objId, err := ConvertInsertedIDToString(res.InsertedID)
	if err != nil {
		return "", errors.New("error converting InsertedID to ObjectID")
	}

	return objId, err
}

func (r *UserMongo) GetUserById(ctx context.Context, id string) (domain.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, errors.New("invalid id")
	}

	res := r.coll.FindOne(ctx, bson.M{"_id": objId})
	if res.Err() != nil {
		if res.Err().Error() == "mongo: no documents in result" {
			return domain.User{}, errors.New("user is not found")
		}

		return domain.User{}, res.Err()
	}

	var user domain.User
	err = res.Decode(&user)

	return user, err
}

func (r *UserMongo) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	res := r.coll.FindOne(ctx, bson.M{"email": email})
	if res.Err() != nil {
		if res.Err().Error() == "mongo: no documents in result" {
			return domain.User{}, errors.New("user is not found")
		}

		return domain.User{}, res.Err()
	}

	var user domain.User
	err := res.Decode(&user)
	if err != nil {
		return domain.User{}, err
	}

	return user, err
}
