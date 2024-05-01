package errhandler

type CommonError struct {
	Message string            `json:"message"`
	Code    string            `json:"code"`
	Errors  []ValidationError `json:"errors"`
}

type ValidationError struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}

type BadRequestError CommonError

func (e *BadRequestError) Error() string {
	return e.Message
}

func NewBadRequestError(message string, errors []ValidationError) error {
	return &BadRequestError{
		Message: message,
		Code:    "invalid.request",
		Errors:  errors,
	}
}

type NotFoundError CommonError

func (e *NotFoundError) Error() string {
	return e.Message
}

func NewNotFoundError(message string, errors []ValidationError) error {
	return &BadRequestError{
		Message: message,
		Code:    "not_found",
		Errors:  errors,
	}
}

type AuthenticationError CommonError

func (e *AuthenticationError) Error() string {
	return e.Message
}

func NewAuthenticationError(message string, errors []ValidationError) error {
	return &BadRequestError{
		Message: message,
		Code:    "invalid credentials",
		Errors:  errors,
	}
}
