package models

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// UserDocument basic user data
type UserDocument struct {
	ID         objectid.ObjectID
	Email      string
	collection *mongo.Collection
}

// Save do mutate User data to DB
func (u *UserDocument) Save() (interface{}, error) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	timeout, err := time.ParseDuration("10s")
	ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if u.ID == objectid.NilObjectID {
		fmt.Println("Create new User")
		if err != nil {
			return nil, err
		}

		result, err := u.collection.InsertOne(ctx, map[string]string{
			"email": u.Email,
		})

		if err != nil {
			return nil, err
		}

		fmt.Println("New user created")
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
		defer cancel()
		insertedResult := u.collection.FindOne(ctx, bson.D{{
			"_id", result.InsertedID,
		}})

		document := UserDocument{}
		err = insertedResult.Decode(&document)

		if err != nil {
			fmt.Println("Error")
			panic(err)
		}

		fmt.Println(document.Email)

		return result.InsertedID, err
	} else {

		fmt.Printf("Update current user %s\n", u.ID.Hex())

	}
	return nil, nil

}

// UserModel data of user
// to use in application
type UserModel struct {
	context container
}

// Create init new user object
func (u *UserModel) Create() UserDocument {
	collection := u.context.DB().Collection("Users")
	if collection == nil {
		panic("Collection not found")
	}
	return UserDocument{
		collection: collection,
	}
}

// CreateUserModel create new user and save
// into database
func CreateUserModel(ctx container) UserModel {
	return UserModel{
		context: ctx,
	}
}
