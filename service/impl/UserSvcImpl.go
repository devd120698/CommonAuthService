package impl

import (
	"commonauthsvc/constants"
	"commonauthsvc/models"
	repo "commonauthsvc/repository"
	"context"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v3"
	"log"
)

type UserServiceImpl struct {
	UserRepo repo.Repository
}

func (svc *UserServiceImpl) CreateUser(ctx context.Context, userInfo models.UserInfoDB) (int, error) {
	// check for emailID
	_, err := svc.UserRepo.GetUserByEmail(ctx, userInfo.Email)
	if err != nil {
		log.Println("Couldnt get user by emailID ", err)
		return 0, &models.BaseError{ErrType: constants.InvalidRequest, ErrDetails: constants.UserAlreadyExists}
	}

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(userInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Could not generate hashed password", err)
		return 0, err
	}
	hashedPassword := string(hashedPasswordBytes)

	userInfo.Password = hashedPassword
	userInfo.IsActive = null.BoolFrom(true)
	lastId, err := svc.UserRepo.AddUser(ctx, userInfo)
	if err != nil {
		log.Println("Could not add user", err)
		return 0, err
	}
	return lastId, nil
}

func (svc *UserServiceImpl) GetUser(ctx context.Context, emailID string) (*models.UserInfoDB, error) {
	res, err := svc.UserRepo.GetUserByEmail(ctx, emailID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (svc *UserServiceImpl) SignIn(ctx context.Context, request *models.UserSignInRequest) (*models.UserSignInResponse, error) {
	res, err := svc.UserRepo.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, &models.BaseError{ErrType: constants.InvalidRequest, ErrDetails: constants.UserUnauthenticated}
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(request.Password))
	if err != nil {
		return nil, err
	}
	return &models.UserSignInResponse{Status: "User Authenticated"}, nil
}
