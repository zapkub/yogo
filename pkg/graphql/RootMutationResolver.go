package graphql

import (
	"context"
	models "yogo/pkg/models"
)

type RootMutationResolver struct {
	container container
}

func (rootMutation *RootMutationResolver) Ping(ctx context.Context) (bool, error) {
	return true, nil
}

func (rootMutation *RootMutationResolver) DeleteUserByID(ctx context.Context) (*models.UserDocument, error) {
	return &models.UserDocument{}, nil
}
