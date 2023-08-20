package value_object

import (
	"errors"
	"restful-api/pkg"
)

type Gender struct {
	value string
}

const (
	GENDER_MALE   = "male"
	GENDER_FEMALE = "female"
)

func (g Gender) NewGender(gender string) (*Gender, error) {

	if gender != GENDER_MALE && gender != GENDER_FEMALE {
		return nil, errors.New(pkg.ERROR_GENDER)
	}
	g.value = gender

	return &g, nil
}

func (e Gender) GetGender() string {
	return e.value
}
