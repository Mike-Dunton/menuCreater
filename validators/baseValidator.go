package validators

import (
	"gopkg.in/validator.v2"
)

func InitValidator() {
	validator.SetValidationFunc("userEmail", userEmail)
}
