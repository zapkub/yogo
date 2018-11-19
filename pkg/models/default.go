package models

import "github.com/mongodb/mongo-go-driver/mongo"

// YogoDocument base Model to use
// when mutate data
type YogoDocument interface {
	Save()
}

type container interface {
	DB() *mongo.Database
}

// Models todo
type Models struct {
	UserModel UserModel
}

// CreateNewModels Todo
func CreateNewModels(c container) *Models {
	return &Models{
		UserModel: CreateUserModel(c),
	}
}
