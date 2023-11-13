package repository

import (
	"commonauthsvc/models"
	"context"
)

type Repository interface {
	AddUser(ctx context.Context, info models.UserInfo) (int, error)
}
