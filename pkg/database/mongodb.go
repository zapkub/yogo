package database

import "github.com/mongodb/mongo-go-driver/mongo"

// CreateMongoDBClient create and connect to mongodb
func CreateMongoDBClient(c context) (*mongo.Client, error) {
	client, err := mongo.NewClient(c.DatabaseConfig().MongoDBURL)
	return client, err
}
