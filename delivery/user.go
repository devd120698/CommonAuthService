package delivery

import (
	"commonauthsvc/constants"
	"commonauthsvc/models"
	svc "commonauthsvc/service"
	"github.com/labstack/echo"
	"log"
	"net/http"
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
}

func (userHttp *UserHTTPHandler) verifyToken(c echo.Context) error {
	return nil
}

func (userHttp *UserHTTPHandler) signin(c echo.Context) error {
	request := models.UserSignInRequest{}
	if err := IsRequestValid(c, request); err != nil {
		return c.JSON(400, &models.BaseError{ErrType: constants.InvalidRequest, ErrDetails: constants.BadRequestForm})
	}

	return nil
}

func (userHttp *UserHTTPHandler) createUser(c echo.Context) error {
	userInfo := models.UserInfo{}

	err := IsRequestValid(c, userInfo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseError{ErrType: constants.InvalidRequest, ErrDetails: constants.BadRequestForm})
	}
	lastId, err := userHttp.UserSvc.CreateUser(c.Request().Context(), userInfo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseError{ErrType: constants.InternalServerError, ErrDetails: "Couldn't insert into data"})
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
