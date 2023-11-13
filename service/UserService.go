package service

import (
	"commonauthsvc/models"
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, info models.UserInfo) (int, error)
	GetUser()
}
