package constants

// ErrType
const (
	InvalidRequest      = "INVALID_REQUEST"
	InternalServerError = "Internal Server Error"
)

// ErrDetails
const (
	BadRequestForm      = "Bad request, please check API documentation."
	UserUnauthenticated = "User couldn't be authenticated."
	UserAlreadyExists   = "User with provided emailID already registered."
)
