package errs

import (
	"errors"
)

var (
	ErrorEmptyBodyContent       = errors.New("request body is empty")
	ErrorRequestBodyInvalid     = errors.New("invalid request body")
	ErrorAccountExist           = errors.New("account with the document number exists")
	ErrorAccountNotExist        = errors.New("account with the account id not exist")
	ErrorIncorrectOperationType = errors.New("operation type incorrect")
)

// ErrorResponse - Use to trow the errors to users
type ErrorResponse struct {
	Error string `json:"error"`
}

// FormatErrorResponse ...
func FormatErrorResponse(mErr error) (res ErrorResponse) {
	errRes := &ErrorResponse{
		Error: mErr.Error(),
	}
	return *errRes
}
