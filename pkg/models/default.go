package models

import "github.com/mongodb/mongo-go-driver/mongo"

// YogoModel base Model to use
// when mutate data
type YogoModel interface {
	Save()
}

type context interface {
	DB() *mongo.Client
}
