package value_object

import "errors"

type Gender struct {
	value string
}

const (
	GENDER_MALE   = "male"
	GENDER_FEMALE = "female"
)

func (g *Gender) GetGender(gender string) (*Gender, error) {

	if gender != GENDER_MALE && gender != GENDER_FEMALE {
		return nil, errors.New("please input valid gender")
	}
	g.value = gender

	return g, nil
}
