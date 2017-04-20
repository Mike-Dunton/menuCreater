package errors

import (
	"errors"
)

// ErrUserExist is returned when creating a user that exists.
var ErrUserExists = errors.New("User already exists")
