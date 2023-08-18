package value_object

import "errors"

type Role struct {
	value string
}

const (
	ROLE_ADMIN = "admin"
	ROLE_USER  = "user"
)

func (r *Role) NewRole(role string) (*Role, error) {

	if role != ROLE_ADMIN && role != ROLE_USER {
		return nil, errors.New("please input valid role")
	}
	r.value = role

	return r, nil
}
