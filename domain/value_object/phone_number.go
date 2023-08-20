package value_object

import (
	"errors"
	"restful-api/pkg"
)

type Phone struct {
	value string
}

func (p Phone) NewPhone(phone string) (*Phone, error) {
	if len(phone) < 10 || len(phone) > 12 {
		return nil, errors.New(pkg.ERROR_PHONE)
	}
	p.value = phone

	return &p, nil
}

func (e Phone) GetPhone() string {
	return e.value
}
