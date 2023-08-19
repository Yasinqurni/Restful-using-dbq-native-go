package value_object

import (
	"fmt"
	"restful-api/pkg"
	"strings"
)

type Email struct {
	value string
}

func (e *Email) NewEmail(email string) (*Email, error) {
	if !strings.Contains(email, "@") {
		return nil, fmt.Errorf(pkg.ERROR_EMAIL)
	}
	e.value = email

	return e, nil
}
