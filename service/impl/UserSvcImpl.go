package impl

import (
	"commonauthsvc/models"
	repo "commonauthsvc/repository"
	"context"
	"golang.org/x/crypto/bcrypt"
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
	userInfo.IsActive = true
	lastId, err := svc.UserRepo.AddUser(ctx, userInfo)
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

func (svc *UserServiceImpl) GetUser() {

}
