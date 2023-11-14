package models

import "gopkg.in/guregu/null.v3"

type UserInfo struct {
	Id           int         `json:"id" db:"id"`
	Name         string      `json:"name" validate:"required"`
	Email        string      `json:"email"`
	PhoneNo      string      `json:"phoneNo" db:"phoneNo"`
	Address      string      `json:"address"`
	Password     string      `json:"password"`
	Role         null.String `json:"role"`
	IsActive     null.Bool   `json:"isActive" db:"isActive"`
	AddedOn      string      `json:"addedOn" db:"addedOn"`
	UpdatedOn    string      `json:"updatedOn" db:"updatedOn"`
	EncPassword  string      `json:"-" db:"encPassword"`
	ProfileImage string      `json:"profileImage" db:"profileImage"`
}

type CreateUserResponse struct {
	Message string `json:"message"`
	UserId  int    `json:"userId"`
}

type UserSignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserSignInResponse struct {
	Token  string `json:"token"`
	Status string `json:"status"`
}

type BaseError struct {
	ErrType    string `json:"errType"`
	ErrDetails string `json:"errDetails"`
}

func (err *BaseError) Error() string {
	return err.ErrDetails
}
