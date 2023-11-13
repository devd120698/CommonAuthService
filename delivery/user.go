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
	userSvc svc.UserService
}

func ConfigureHTTPHandler(e *echo.Echo, userSvc svc.UserService) {
	userHTTPHandler := UserHTTPHandler{
		userSvc: userSvc,
	}
	userHTTPHandler.AddHandlers(e)
}

func (userHttp *UserHTTPHandler) AddHandlers(e *echo.Echo) {
	e.POST("/createUser", userHttp.createUser)
	e.GET("/getUser", userHttp.getUser)
}

func (userHttp *UserHTTPHandler) createUser(c echo.Context) error {
	userInfo := models.UserInfo{}

	err := IsRequestValid(c, userInfo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseError{ErrType: constants.BadRequestForm, ErrDetails: "Couldn't decode request body"})
	}
	lastId, err := userHttp.userSvc.CreateUser(c.Request().Context(), userInfo)
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
	//_ = c.QueryParam("id")
	//_, err := GetDbConnection()
	//if err != nil {
	//	return c.JSON(http.StatusBadGateway, "Couldn't get DB connection")
	//}
	return nil
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
