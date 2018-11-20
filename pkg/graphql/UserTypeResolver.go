package graphql

import (
	"context"
	models "yogo/pkg/models"
)

type UserTypeResolver struct {
	container container
}

func (userResolver UserTypeResolver) ID(ctx context.Context, parent *models.UserDocument) (string, error) {
	return "", nil
}
