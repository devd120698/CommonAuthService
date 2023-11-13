package models

type UserInfo struct {
	Name           string `json:"name" validate:"required"`
	Email          string `json:"email"`
	PhoneNo        string `json:"phoneNo"`
	Address        string `json:"address"`
	Password       string `json:"password"`
	Role           string `json:"role"`
	ProfilePicture string `json:"profilePicture"`
	IsActive       bool   `json:"isActive"`
}

type CreateUserResponse struct {
	Message string `json:"message"`
	UserId  int    `json:"userId"`
}

type BaseError struct {
	ErrType    string `json:"errType"`
	ErrDetails string `json:"errDetails"`
}

func (err *BaseError) Error() string {
	return err.ErrDetails
}
