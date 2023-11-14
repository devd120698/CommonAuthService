package impl

import (
	"commonauthsvc/models"
	repo "commonauthsvc/repository"
	"context"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v3"
)

type UserServiceImpl struct {
	UserRepo repo.Repository
}

func (svc *UserServiceImpl) CreateUser(ctx context.Context, userInfo models.UserInfo) (int, error) {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(userInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	hashedPassword := string(hashedPasswordBytes)

	userInfo.Password = hashedPassword
	userInfo.IsActive = null.BoolFrom(true)
	lastId, err := svc.UserRepo.AddUser(ctx, userInfo)
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

func (svc *UserServiceImpl) GetUser(ctx context.Context, emailID string) (*[]models.UserInfo, error) {
	res, err := svc.UserRepo.GetUserByEmail(ctx, emailID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
