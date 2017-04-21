package errors

import (
	"errors"
)

type error struct {
	Message string
}

// ErrUserExist is returned when creating a user that exists.
var ErrUserExists = errors.New("User already exists")

var ErrValidatingSignUp = &error{
	Message: "We couldn't validate that email address",
}
