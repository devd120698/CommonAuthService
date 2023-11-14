package repository

import (
	"commonauthsvc/models"
	"context"
)

type Repository interface {
	AddUser(ctx context.Context, info models.UserInfoDB) (int, error)
	GetUserByEmail(ctx context.Context, email string) (*models.UserInfoDB, error)
}
