package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func ConnectClient(query string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(query))
	if err != nil {
		return nil, err
	}

	return client, Ping(client)
}

func Ping(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return client.Ping(ctx, readpref.Primary())
}
