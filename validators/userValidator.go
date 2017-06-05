package validators

import (
	"gopkg.in/validator.v2"

	"errors"
	"reflect"
	"regexp"
)

func userEmail(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	if st.Kind() != reflect.String {
		return validator.ErrUnsupported
	}
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if matchesEmailRegex := Re.MatchString(st.String()); matchesEmailRegex {
		return nil
	}
	return errors.New("Enter Valid Email")
}
