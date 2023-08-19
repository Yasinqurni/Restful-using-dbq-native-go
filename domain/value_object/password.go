package value_object

import (
	"errors"
	"restful-api/pkg"
)

type Password struct {
	value string
}

func (p *Password) NewPassword(password string) (*Password, error) {
	if len(password) != 6 {
		return nil, errors.New(pkg.ERROR_PASSWORD)
	}
	p.value = password

	return p, nil
}
