package data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"sync"
	"time"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

var col *mongo.Collection

var DefaultTimeout = 100000 * time.Second

const (
	DB      = "writerxl"
	PROFILE = "profiles"
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		defer cancel()

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_CONNECTION")))
		if err != nil {
			clientInstanceError = err
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}
