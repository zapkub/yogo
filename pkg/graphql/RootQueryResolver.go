package graphql

import (
	context "context"
	models "yogo/pkg/models"
)

type RootQueryResolver struct {
	container container
}

func (queryResolver *RootQueryResolver) Version(c context.Context) (string, error) {
	return "1.0.0", nil
}

func (queryResolver *RootQueryResolver) Users(c context.Context) ([]models.UserDocument, error) {
	var result []models.UserDocument
	return result, nil
}
