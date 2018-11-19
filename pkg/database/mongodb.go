package database

import (
	"context"
	"fmt"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// CreateMongoDBClient create and connect to mongodb
func CreateMongoDBClient(c container) (*mongo.Database, error) {
	client, err := mongo.NewClient(c.DatabaseConfig().MongoDBURL)

	err = client.Connect(context.TODO())
	fmt.Printf("\nConnect to Yogo database at %s\n", c.DatabaseConfig().MongoDBURL)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	return client.Database("yogo"), err
}
