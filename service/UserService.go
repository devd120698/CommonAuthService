package service

import (
	"commonauthsvc/models"
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, info models.UserInfoDB) (int, error)
	GetUser(ctx context.Context, emailID string) (*models.UserInfoDB, error)
	SignIn(ctx context.Context, request *models.UserSignInRequest) (*models.UserSignInResponse, error)
}
