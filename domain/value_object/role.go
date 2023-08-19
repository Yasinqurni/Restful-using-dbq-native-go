package value_object

import (
	"errors"
	"restful-api/pkg"
)

type Role struct {
	value string
}

const (
	ROLE_ADMIN = "admin"
	ROLE_USER  = "user"
)

func (r *Role) NewRole(role string) (*Role, error) {
	if role != ROLE_ADMIN && role != ROLE_USER {
		return nil, errors.New(pkg.ERROR_ROLE)
	}
	r.value = role

	return r, nil
}
