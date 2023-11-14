package impl

import (
	"commonauthsvc/constants"
	"commonauthsvc/models"
	repo "commonauthsvc/repository"
	"context"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v3"
	"log"
	"time"
)

type UserServiceImpl struct {
	UserRepo repo.Repository
}

func (svc *UserServiceImpl) CreateUser(ctx context.Context, userInfo models.UserInfoDB) (int, error) {
	// check for emailID
	res, err := svc.UserRepo.GetUserByEmail(ctx, userInfo.Email)
	if (err != nil || res != nil) && err.Error() != "No user with this email ID" {
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
	if (err != nil) && err.Error() != "No user with this email ID" {
		return nil, &models.BaseError{ErrType: constants.InvalidRequest, ErrDetails: constants.InvalidCreds}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(res.EncPassword), []byte(request.Password)); err != nil {
		log.Println("invalid Credentials")
		return nil, &models.BaseError{ErrType: constants.InvalidRequest, ErrDetails: constants.InvalidCreds}
	}
	// create a JWT token and return
	// Set custom claims
	claims := &models.JwtCustomClaims{
		Email: res.Email,
		Role:  res.Role.String,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 25)),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}
	return &models.UserSignInResponse{Status: "User Authenticated", Token: t}, nil
}
