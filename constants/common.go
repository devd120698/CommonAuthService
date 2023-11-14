package constants

// ErrType
const (
	InvalidRequest      = "INVALID_REQUEST"
	InternalServerError = "Internal Server Error"
)

// ErrDetails
const (
	BadRequestForm    = "Bad request, please check API documentation."
	InvalidCreds      = "Invalid credentials. Please try again."
	UserAlreadyExists = "User with provided emailID already registered."
)
