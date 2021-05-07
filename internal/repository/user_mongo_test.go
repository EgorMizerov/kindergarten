package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/EgorMizerov/kindergarten/internal/domain"
	"github.com/EgorMizerov/kindergarten/pkg/database"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func testUsersData(ctx context.Context, coll *mongo.Collection) ([]string, error) {
	var ids []string
	data := []domain.User{
		domain.User{Email: "example@mail.com"},
		domain.User{Email: "example@gmail.com"},
		domain.User{Email: "example@microsoft.com"},
	}

	for _, v := range data {
		res, err := coll.InsertOne(ctx, bson.M{"email": v.Email})
		if err != nil {
			return nil, err
		}

		id, err := ConvertInsertedIDToString(res.InsertedID)
		if err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

func TestUserMongo_GetUserByEmail(t *testing.T) {
	ctx := context.Background()

	// connect to mongodb
	mongoClient, err := database.ConnectClient("mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&ssl=false")
	if err != nil {
		t.Fatal(err)
	}

	// create collection
	coll := mongoClient.Database("test").Collection("users")
	defer coll.Drop(ctx)

	// insert test data
	_, err = testUsersData(ctx, coll)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserMongo(coll)

	tests := []struct {
		name          string
		email         string
		expectedEmail string
		expectedErr   error
	}{
		{
			name:          "Ok",
			email:         "example@mail.com",
			expectedEmail: "example@mail.com",
		},
		{
			name:        "Users is not found",
			email:       "notfound@mail.com",
			expectedErr: errors.New("user is not found"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := repo.GetUserByEmail(ctx, test.email)

			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestUserMongo_GetUserById(t *testing.T) {
	ctx := context.Background()

	// connect to mongodb
	mongoClient, err := database.ConnectClient("mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&ssl=false")
	if err != nil {
		t.Fatal(err)
	}

	// create collection
	coll := mongoClient.Database("test").Collection("users")
	//defer coll.Drop(ctx)

	// insert test data
	ids, err := testUsersData(ctx, coll)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserMongo(coll)

	tests := []struct {
		name        string
		inputId     string
		expectedId  string
		expectedErr error
	}{
		{
			name:       "Ok",
			inputId:    ids[0],
			expectedId: fmt.Sprintf("ObjectID(\"%s\")", ids[0]),
		},
		{
			name:        "User is not found",
			inputId:     "60957c617eab0884aa340465",
			expectedId:  "ObjectID(\"000000000000000000000000\")",
			expectedErr: errors.New("user is not found"),
		},
		{
			name:        "Invalid ID",
			inputId:     "qqqw",
			expectedId:  "ObjectID(\"000000000000000000000000\")",
			expectedErr: errors.New("invalid id"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			user, err := repo.GetUserById(ctx, test.inputId)

			assert.Equal(t, test.expectedErr, err)
			assert.Equal(t, test.expectedId, user.Id.String())
		})
	}
}
