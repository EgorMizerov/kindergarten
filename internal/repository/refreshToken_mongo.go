package repository

import (
	"context"
	"errors"
	"github.com/EgorMizerov/kindergarten/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RefreshTokenMongo struct {
	coll *mongo.Collection
}

func NewRefreshTokenMongo(coll *mongo.Collection) *RefreshTokenMongo {
	return &RefreshTokenMongo{coll: coll}
}

func (r *RefreshTokenMongo) SetRefreshTokenByID(ctx context.Context, session domain.Session, id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.coll.UpdateOne(ctx,
		bson.M{"_id": objId},
		bson.M{"$set": bson.M{
			"session": bson.M{
				"token": session.Token,
				"exp":   session.Exp,
			},
		}})

	return err
}

func (r *RefreshTokenMongo) GetRefreshTokenById(ctx context.Context, id string) (domain.Session, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Session{}, errors.New("invalid id")
	}

	res := r.coll.FindOne(ctx, bson.M{"_id": objId}, options.FindOne().SetProjection(bson.M{"session": 1, "_id": 0}))
	if res.Err() != nil {
		if res.Err().Error() == "mongo: no documents in result" {
			return domain.Session{}, errors.New("user is not found")
		}

		return domain.Session{}, res.Err()
	}

	var session struct{ Session domain.Session }
	err = res.Decode(&session)

	return session.Session, err
}
