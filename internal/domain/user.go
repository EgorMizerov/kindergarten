package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id primitive.ObjectID `bson:"_id"`

	PasswordHash string
	Email        string
	Phone        string

	FirstName    string
	MiddleName   string
	LastName     string
	DateOfBirth  int64
	IsAdmin      bool
	CreationDate int64
}
