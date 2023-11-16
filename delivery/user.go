package delivery

import (
	"commonauthsvc/constants"
	"commonauthsvc/models"
	svc "commonauthsvc/service"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strings"
)

type UserHTTPHandler struct {
	UserSvc svc.UserService
}

func ConfigureHTTPHandler(e *echo.Echo, userSvc svc.UserService) {
	userHTTPHandler := UserHTTPHandler{
		UserSvc: userSvc,
	}
	userHTTPHandler.AddHandlers(e)
}

func (userHttp *UserHTTPHandler) AddHandlers(e *echo.Echo) {
	e.POST("/createUser", userHttp.createUser)
	e.GET("/getUser", userHttp.getUser)
	e.POST("/auth", userHttp.signin)
	e.POST("/verify", userHttp.verifyToken)
	e.PATCH("/signout", userHttp.signOut)
}

func (userHttp *UserHTTPHandler) signOut(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Authorization header is missing")
	}
	//session management and cookies
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid token")
	}
}

func (userHttp *UserHTTPHandler) verifyToken(c echo.Context) error {
	// Get token from 'Authorization' header
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Authorization header is missing")
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid token")
	}
	tokenClaims, err := jwt.ParseWithClaims(token, &models.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		log.Printf("error parsing token: %v\n", err)
		return c.JSON(http.StatusBadRequest, &models.BaseError{ErrType: constants.InvalidRequest, ErrDetails: err.Error()})
	}

	emailID := tokenClaims.Claims.(*models.JwtCustomClaims).Email
	//get user info from emailID
	user, err := userHttp.UserSvc.GetUser(c.Request().Context(), emailID)
	if err != nil || user == nil {
		return c.JSON(http.StatusBadRequest, models.BaseError{ErrType: constants.InvalidRequest, ErrDetails: constants.BadRequestForm})
	}
	log.Println("Token is valid")
	return c.JSON(http.StatusOK, models.Response{Message: "Token in valid"})
}

func (userHttp *UserHTTPHandler) signin(c echo.Context) error {
	request := models.UserSignInRequest{}
	//if err := IsRequestValid(c, request); err != nil {
	//	return c.JSON(http.StatusBadRequest, &models.BaseError{ErrType: constants.InvalidRequest, ErrDetails: constants.BadRequestForm})
	//}
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Couldn't decode request body")
	}
	res, err := userHttp.UserSvc.SignIn(c.Request().Context(), &request)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusForbidden, &models.BaseError{ErrType: constants.InvalidRequest, ErrDetails: constants.InvalidCreds})
	}
	return c.JSON(http.StatusAccepted, res)
}

func (userHttp *UserHTTPHandler) createUser(c echo.Context) error {
	userInfo := &models.UserInfoDB{}

	//err := IsRequestValid(c, userInfo)
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, models.BaseError{ErrType: constants.InvalidRequest, ErrDetails: constants.BadRequestForm})
	//}
	err := json.NewDecoder(c.Request().Body).Decode(&userInfo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Couldn't decode request body")
	}
	lastId, err := userHttp.UserSvc.CreateUser(c.Request().Context(), *userInfo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseError{ErrType: err.Error(), ErrDetails: err.Error()})
	}
	msg := models.CreateUserResponse{
		Message: "User created",
		UserId:  lastId,
	}
	return c.JSON(http.StatusCreated, msg)
}

func (userHttp *UserHTTPHandler) getUser(c echo.Context) error {
	emailID := c.QueryParam("email")
	user, err := userHttp.UserSvc.GetUser(c.Request().Context(), emailID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseError{ErrType: constants.InvalidRequest, ErrDetails: constants.BadRequestForm})
	}
	return c.JSON(http.StatusOK, user)
}

func IsRequestValid(c echo.Context, m interface{}) error {
	err := c.Bind(m)
	if err != nil {
		log.Println(err)
		return &models.BaseError{ErrType: constants.BadRequestForm}
	}
	err = c.Validate(m)
	if err != nil {
		return err
	}
	return nil
}
