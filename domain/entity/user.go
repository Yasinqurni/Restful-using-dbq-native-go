package entity

import "salt-academy_clean-archy/domain/value_object"

type User struct {
	name     string
	email    string
	phone    string
	password string
	role     *value_object.Role
	gender   *value_object.Gender
}
